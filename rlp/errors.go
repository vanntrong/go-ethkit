package rlp

import "errors"

var ErrNegativeBigInt = errors.New("rlp: cannot encode negative big.Int")
var ErrLimitSizeString = errors.New("rlp: reach limit size of string")
var ErrUnSupportInputType = errors.New("rlp: unsupported input type")
