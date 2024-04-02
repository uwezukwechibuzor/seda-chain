package types

import errorsmod "cosmossdk.io/errors"

// x/wasm-storage module sentinel errors
var (
	ErrDataRequestWasmExists   = errorsmod.Register(ModuleName, 2, "data Request Wasm with given hash already exists")
	ErrOverlayWasmExists       = errorsmod.Register(ModuleName, 3, "overlay Wasm with given hash already exists")
	ErrInvalidAuthorityAddress = errorsmod.Register(ModuleName, 4, "invalid authority address")
	ErrInvalidAuthority        = errorsmod.Register(ModuleName, 5, "invalid authority")
	ErrInvalidSenderAddress    = errorsmod.Register(ModuleName, 6, "invalid sender address")
	ErrInvalidAdminAddress     = errorsmod.Register(ModuleName, 7, "invalid admin address")
	ErrWasmNotGzipCompressed   = errorsmod.Register(ModuleName, 8, "wasm is not gzip compressed")
	ErrInstantiateContract     = errorsmod.Register(ModuleName, 9, "error during contract instantiation")
)
