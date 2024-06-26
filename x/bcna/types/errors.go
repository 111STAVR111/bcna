package types

import (
	sdkioerrors "cosmossdk.io/errors"
)

// x/bcna module sentinel errors
var (
	ErrDuplicateBitcannaid = sdkioerrors.Register(ModuleName, 1101, "BitCannaID already exists")
	ErrKeyNotFound         = sdkioerrors.Register(ModuleName, 1102, "Key doesn't exists")
	ErrUnauthorized        = sdkioerrors.Register(ModuleName, 1103, "Incorrect owner")
	ErrUnrecognized        = sdkioerrors.Register(ModuleName, 1104, "Unrecognized message")
	ErrInvalidAddress      = sdkioerrors.Register(ModuleName, 1105, "invalid address")
	ErrMaxCharacters       = sdkioerrors.Register(ModuleName, 1106, "input exceeds the permitted length limit")
)
