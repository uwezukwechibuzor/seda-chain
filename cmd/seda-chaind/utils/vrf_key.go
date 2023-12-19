package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	cmtos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/types"

	"github.com/cosmos/cosmos-sdk/client"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	vrf "github.com/sedaprotocol/vrf-go"
)

const VRFKeyFileName = "vrf_key.json"

type VRFSigner interface {
	VRFProve(alpha []byte) (pi, beta []byte, err error)
	VRFVerify(publicKey, alpha, pi []byte) (beta []byte, err error)
	SignTransaction(ctx sdk.Context, txBuilder client.TxBuilder, txConfig client.TxConfig,
		signMode signing.SignMode, account sdk.AccountI) (signing.SignatureV2, error)
}

var _ VRFSigner = &VRFKey{}

type VRFKey struct {
	Address types.Address    `json:"address"`
	PubKey  sdkcrypto.PubKey `json:"pub_key"`
	PrivKey crypto.PrivKey   `json:"priv_key"` // TO-DO can we not export it?

	filePath string
	vrf      *vrf.VRFStruct
}

// Save persists the VRFKey to its filePath.
func (key VRFKey) Save() error {
	outFile := key.filePath
	if outFile == "" {
		return fmt.Errorf("key's file path is empty")
	}

	vrfKeyFile := struct {
		PrivKey crypto.PrivKey `json:"priv_key"` // TO-DO can we not export it?
	}{
		PrivKey: key.PrivKey,
	}

	jsonBytes, err := cmtjson.MarshalIndent(vrfKeyFile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal key: %v", err)
	}

	err = os.WriteFile(outFile, jsonBytes, 0o600)
	if err != nil {
		return fmt.Errorf("failed to write key file: %v", err)
	}
	return nil
}

// VRFProve uses the VRF key to compute the VRF hash output (beta)
// and the proof that it was computed correctly (pi).
func (v *VRFKey) VRFProve(alpha []byte) (pi, beta []byte, err error) {
	pi, err = v.vrf.Prove(v.PrivKey.Bytes(), alpha)
	if err != nil {
		return nil, nil, err
	}
	beta, err = v.vrf.ProofToHash(pi)
	if err != nil {
		return nil, nil, err
	}
	return pi, beta, nil
}

// VRFVerify verifies that beta is the correct VRF hash of the alpha
// under private key associated with the given public key. It also
// outputs the hash output beta.
func (v *VRFKey) VRFVerify(publicKey, alpha, pi []byte) (beta []byte, err error) {
	beta, err = v.vrf.Verify(publicKey, alpha, pi)
	if err != nil {
		return nil, err
	}
	return beta, nil
}

// SignTransaction signs a given transaction with the VRF key and
// returns the resulting signature. The given account must belong
// to the VRF key.
func (v *VRFKey) SignTransaction(
	ctx sdk.Context, txBuilder client.TxBuilder, txConfig client.TxConfig,
	signMode signing.SignMode, account sdk.AccountI,
) (signing.SignatureV2, error) {
	var sigV2 signing.SignatureV2

	if !bytes.Equal(account.GetPubKey().Bytes(), v.PubKey.Bytes()) {
		return sigV2, fmt.Errorf("the account does not belong to the vrf key")
	}

	signerData := authsigning.SignerData{
		ChainID:       ctx.ChainID(),
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
		PubKey:        v.PubKey,
		Address:       account.GetAddress().String(),
	}

	bytesToSign, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		txConfig.SignModeHandler(),
		signMode,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return sigV2, err
	}

	sigBytes, err := v.PrivKey.Sign(bytesToSign)
	if err != nil {
		return sigV2, err
	}

	sigV2 = signing.SignatureV2{
		PubKey: v.PubKey,
		Data: &txsigning.SingleSignatureData{
			SignMode:  signMode,
			Signature: sigBytes,
		},
		Sequence: account.GetSequence(),
	}
	return sigV2, nil
}

// NewVRFKey generates a new VRFKey from the given key and key file path.
func NewVRFKey(privKey crypto.PrivKey, keyFilePath string) (*VRFKey, error) {
	vrfStruct := vrf.NewK256VRF()
	pubKey, err := cryptocodec.FromCmtPubKeyInterface(privKey.PubKey())
	if err != nil {
		return nil, err
	}
	return &VRFKey{
		Address:  privKey.PubKey().Address(),
		PubKey:   pubKey,
		PrivKey:  privKey,
		filePath: keyFilePath,
		vrf:      &vrfStruct,
	}, nil
}

// LoadOrGenVRFKey loads a VRFKey from the given file path
// or else generates a new one and saves it to the file path.
func LoadOrGenVRFKey(keyFilePath string) (*VRFKey, error) {
	var vrfKey *VRFKey
	var err error
	if cmtos.FileExists(keyFilePath) {
		vrfKey, err = LoadVRFKey(keyFilePath)
		if err != nil {
			return nil, err
		}
	} else {
		vrfKey, err = NewVRFKey(secp256k1.GenPrivKey(), keyFilePath)
		if err != nil {
			return nil, err
		}
		err = vrfKey.Save()
		if err != nil {
			return nil, err
		}
	}
	return vrfKey, nil
}

func LoadVRFKey(keyFilePath string) (*VRFKey, error) {
	keyJSONBytes, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading VRF key from %v: %v", keyFilePath, err)
	}

	vrfKeyFile := struct {
		PrivKey crypto.PrivKey `json:"priv_key"` // TO-DO can we not export it?
	}{}
	err = cmtjson.Unmarshal(keyJSONBytes, &vrfKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling VRF key from %v: %v", keyFilePath, err)
	}

	vrfKey, err := NewVRFKey(vrfKeyFile.PrivKey, keyFilePath)
	if err != nil {
		return nil, err
	}

	return vrfKey, nil
}

func InitializeVRFKey(config *cfg.Config) (vrfPubKey sdkcrypto.PubKey, err error) {
	pvKeyFile := config.PrivValidatorKeyFile()
	if err := os.MkdirAll(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return nil, fmt.Errorf("could not create directory %q: %w", filepath.Dir(pvKeyFile), err)
	}

	vrfKeyFile := PrivValidatorKeyFileToVRFKeyFile(config.PrivValidatorKeyFile())
	vrfKey, err := LoadOrGenVRFKey(vrfKeyFile)
	if err != nil {
		return nil, err
	}
	return vrfKey.PubKey, nil
}

// PrivValidatorKeyFileToVRFKeyFile returns the path to the VRF key file
// given a path to the private validator key file. The two files should
// be placed in the same directory.
func PrivValidatorKeyFileToVRFKeyFile(pvFile string) string {
	return filepath.Join(filepath.Dir(pvFile), VRFKeyFileName)
}