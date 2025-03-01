package utils

import (
	"encoding/hex"
	"strings"
)

// Contract name constants
const (
	ContractBrdgmng = "brdgmng.xsat"
	ContractCbridge = "cbridge.xsat"
	ContractRes     = "res.xsat"
)

// Index position constants
const (
	IndexPrimary   = "primary"
	IndexSecondary = "secondary"
	IndexTertiary  = "tertiary"
	IndexFourth    = "fourth"
	IndexFifth     = "fifth"
	IndexSixth     = "sixth"
	IndexSeventh   = "seventh"
	IndexEighth    = "eighth"
	IndexNinth     = "ninth"
	IndexTenth     = "tenth"
)

// Key type constants
const (
	KeyTypeI64       = "i64"
	KeyTypeI128      = "i128"
	KeyTypeI256      = "i256"
	KeyTypeFloat64   = "float64"
	KeyTypeFloat128  = "float128"
	KeyTypeRipemd160 = "ripemd160"
	KeyTypeSha256    = "sha256"
	KeyTypeName      = "name"
)

// ComputeId calculates SHA-256 hash
// Converts an EVM address to a specific format ID
func ComputeId(evmAddress string) string {
	evmAddress = strings.TrimPrefix(evmAddress, "0x")
	result := make([]byte, 32)
	decodedAddr, _ := hex.DecodeString(evmAddress)
	copy(result[12:], decodedAddr)
	return hex.EncodeToString(result)
}
