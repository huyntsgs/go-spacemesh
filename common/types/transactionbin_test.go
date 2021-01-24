package types

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var txValueOldCoinTx1 = OldCoinTx{
	AccountNonce: 5,
	Recipient:    Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	GasLimit:     10,
	Fee:          1,
	Amount:       100,
}
var txCaseOldCoinTx1Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x06,
	/*sign:*/ 0xfd, 0xbd, 0xe5, 0xb1, 0x76, 0xc8, 0x08, 0x01, 0x6a, 0x70, 0xc8, 0xd3, 0x80, 0xd9, 0x3d, 0xeb, 0xf7, 0xa0, 0x4b, 0xef, 0x64, 0x01, 0xeb, 0xc5, 0xe7, 0xb3, 0x04, 0xe8, 0x6d, 0x2e, 0xee, 0x60, 0xbe, 0xe2, 0x84, 0x57, 0x18, 0x0a, 0x71, 0xb7, 0xc6, 0x19, 0xf0, 0xb9, 0x34, 0xd9, 0x50, 0x5a, 0xba, 0xd3, 0xa7, 0xf5, 0x60, 0x40, 0x51, 0xd5, 0xd8, 0x45, 0x29, 0x1a, 0xb8, 0x31, 0x9c, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseOldCoinTx1EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x07,
	/*sign:*/ 0xba, 0x83, 0xd0, 0xe6, 0x72, 0x3f, 0xde, 0xac, 0x06, 0x00, 0x76, 0xf9, 0xa8, 0x83, 0x2c, 0x81, 0x45, 0x62, 0x07, 0x6d, 0x40, 0x49, 0xe0, 0x3c, 0x09, 0xcd, 0x98, 0xd8, 0xa4, 0x63, 0x4b, 0x33, 0xff, 0x85, 0x69, 0x6b, 0x60, 0x04, 0xef, 0x22, 0xf5, 0x4f, 0x1f, 0x00, 0x32, 0x5d, 0x8c, 0xc7, 0xfa, 0x2e, 0x80, 0xc8, 0x0f, 0x84, 0xde, 0x16, 0x92, 0xa7, 0xfb, 0x4f, 0xb2, 0x3a, 0xd3, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	/*pkey:*/
}
var txValueOldCoinTx2 = OldCoinTx{
	AccountNonce: 6,
	Recipient:    Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	GasLimit:     10,
	Fee:          2,
	Amount:       101,
}
var txCaseOldCoinTx2Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x06,
	/*sign:*/ 0xc6, 0x83, 0x8d, 0xdd, 0x4d, 0x79, 0x98, 0xde, 0x68, 0x42, 0x85, 0x2e, 0x6e, 0x39, 0xb2, 0xe2, 0x43, 0xfc, 0x09, 0xb0, 0x8f, 0x03, 0x1b, 0xe9, 0x0b, 0xe9, 0x20, 0x23, 0xa3, 0xcf, 0x8c, 0x12, 0xdf, 0x75, 0xfa, 0xa8, 0xf5, 0x96, 0x37, 0xd9, 0x3d, 0x18, 0xc9, 0x5a, 0x1c, 0xfc, 0x8b, 0xca, 0x61, 0x09, 0x2c, 0x07, 0xd6, 0x20, 0x3c, 0xc5, 0xab, 0x94, 0xa7, 0x9e, 0x9b, 0x39, 0xd9, 0x08,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseOldCoinTx2EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x07,
	/*sign:*/ 0x22, 0x2f, 0xeb, 0xf9, 0xb8, 0x10, 0xec, 0x05, 0x7c, 0xa8, 0xa2, 0xea, 0x01, 0xf7, 0x69, 0x22, 0xd4, 0x1e, 0x1a, 0xaa, 0x70, 0x58, 0x5c, 0x7b, 0x10, 0xa8, 0x86, 0xe4, 0x3e, 0x92, 0x5d, 0xe1, 0x6c, 0x29, 0x53, 0x2e, 0xdb, 0x72, 0xb1, 0x8f, 0xe1, 0xea, 0x95, 0x88, 0x12, 0xcd, 0x7c, 0x36, 0x24, 0x62, 0x83, 0x40, 0x66, 0x9a, 0xa5, 0xda, 0x5a, 0xbd, 0x58, 0xc2, 0xcc, 0x6f, 0x6f, 0x01,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65,
	/*pkey:*/
}
var txValueSimpleCoinTx1 = SimpleCoinTx{
	TTL:       0,
	Nonce:     5,
	Recipient: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:    102,
	GasLimit:  1,
	GasPrice:  10,
}
var txCaseSimpleCoinTx1Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x00,
	/*sign:*/ 0x1e, 0x31, 0xe8, 0xf4, 0xc3, 0xca, 0x13, 0x91, 0xf0, 0x28, 0x63, 0xac, 0x70, 0x2d, 0xea, 0x48, 0xef, 0x9d, 0x92, 0xfb, 0xad, 0x6f, 0x78, 0x48, 0x6f, 0xb6, 0xa3, 0xf9, 0x89, 0x70, 0xee, 0x80, 0x76, 0x8f, 0xd0, 0xa9, 0x20, 0x2c, 0x7c, 0x3a, 0x1b, 0x76, 0x6e, 0xda, 0x0c, 0xab, 0x9a, 0xbe, 0x8f, 0x8a, 0x37, 0xc1, 0x0d, 0xfa, 0x06, 0xbd, 0x5a, 0x18, 0xd7, 0xf1, 0x3a, 0xb3, 0x74, 0x0e,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSimpleCoinTx1EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x01,
	/*sign:*/ 0xb4, 0x9c, 0xd9, 0x3a, 0x91, 0xff, 0x24, 0xb6, 0xfa, 0xe2, 0x74, 0xa3, 0x9e, 0x29, 0xe2, 0xcc, 0x3c, 0x99, 0xbb, 0x9f, 0x03, 0x53, 0x74, 0x9b, 0x0d, 0xe5, 0xe8, 0x95, 0xb1, 0xd2, 0x61, 0x1b, 0xb9, 0xdb, 0xc7, 0x2a, 0x02, 0xfa, 0xed, 0x97, 0x04, 0x6d, 0x2d, 0xa5, 0x96, 0x37, 0x09, 0x7f, 0xdf, 0x34, 0xd5, 0x6d, 0xea, 0x80, 0x8d, 0x89, 0xb1, 0x5b, 0xe6, 0x72, 0x91, 0xd7, 0x12, 0x02,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	/*pkey:*/
}
var txValueSimpleCoinTx2 = SimpleCoinTx{
	TTL:       0,
	Nonce:     8,
	Recipient: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:    103,
	GasLimit:  10,
	GasPrice:  1,
}
var txCaseSimpleCoinTx2Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x00,
	/*sign:*/ 0xc2, 0x97, 0x4f, 0x82, 0x41, 0xd6, 0xfb, 0x49, 0x1e, 0x39, 0x2e, 0xea, 0x32, 0xe5, 0x9a, 0x41, 0xb2, 0x42, 0x5e, 0x83, 0xc0, 0xb0, 0x9e, 0x60, 0x6d, 0xb6, 0x1a, 0xc9, 0xe3, 0xf4, 0x55, 0x09, 0x50, 0x97, 0xb2, 0xc4, 0x47, 0xe3, 0x1d, 0x7e, 0xad, 0x11, 0xef, 0xd7, 0xb1, 0x38, 0xcc, 0x2b, 0x49, 0x1b, 0x51, 0xa6, 0xfd, 0x69, 0xfa, 0x7b, 0x3b, 0x16, 0xde, 0xf6, 0x18, 0x7c, 0x37, 0x05,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSimpleCoinTx2EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x01,
	/*sign:*/ 0xdd, 0xd1, 0xe9, 0x56, 0xc0, 0xf2, 0x8a, 0x12, 0xb3, 0x79, 0x5e, 0xa9, 0xa6, 0x08, 0x59, 0xb3, 0x10, 0x95, 0xed, 0xe0, 0xf1, 0x4f, 0x8e, 0x09, 0xb2, 0x7f, 0x3c, 0xa6, 0x28, 0x02, 0xd3, 0xae, 0x22, 0xe1, 0xeb, 0xff, 0x15, 0xf2, 0x7a, 0xc1, 0xd6, 0xce, 0xff, 0x7b, 0x5d, 0xa1, 0x18, 0x36, 0x11, 0xae, 0x79, 0x05, 0xc3, 0x99, 0x4d, 0x5e, 0x67, 0xfe, 0xf1, 0x39, 0xd3, 0xba, 0xd3, 0x0b,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	/*pkey:*/
}
var txValueCallAppTx1 = CallAppTx{
	TTL:        0,
	Nonce:      5,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     104,
	GasLimit:   1,
	GasPrice:   10,
	CallData:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
}
var txCaseCallAppTx1Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x02,
	/*sign:*/ 0xed, 0xe8, 0x51, 0x5f, 0x05, 0xa5, 0x4d, 0xf6, 0x60, 0xb8, 0x3c, 0x6c, 0xc3, 0xd3, 0xa9, 0x12, 0x91, 0x19, 0x0f, 0xe2, 0xaa, 0x4b, 0x66, 0x9d, 0x45, 0x29, 0x3b, 0xe4, 0xe3, 0xdf, 0x96, 0x00, 0xb6, 0xa9, 0x2d, 0xf9, 0xae, 0x6a, 0x84, 0xd4, 0x5f, 0xeb, 0x92, 0x97, 0x4d, 0x74, 0x8c, 0x91, 0xb9, 0xe5, 0x5c, 0x05, 0xe9, 0x0c, 0xb1, 0x3d, 0xb7, 0xee, 0xb3, 0x99, 0x44, 0x82, 0x5c, 0x0e,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx1EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x03,
	/*sign:*/ 0xdb, 0xe5, 0x56, 0x2b, 0xc1, 0x9d, 0x1a, 0x54, 0x9d, 0x5a, 0x14, 0xd1, 0xba, 0x15, 0x0c, 0x05, 0x65, 0x2b, 0x88, 0x74, 0xc3, 0x25, 0xc2, 0xc9, 0x68, 0x3a, 0xd6, 0x97, 0x7e, 0x73, 0x05, 0x19, 0xf8, 0xac, 0xf0, 0x4a, 0x2f, 0xb5, 0x1b, 0xe9, 0x97, 0x5c, 0x06, 0x66, 0xc5, 0x5b, 0xdb, 0x44, 0x12, 0xed, 0x26, 0x90, 0x9d, 0x9f, 0x73, 0x3b, 0x34, 0x34, 0x75, 0x9c, 0xa8, 0xfb, 0xe8, 0x05,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/
}
var txValueCallAppTx2 = CallAppTx{
	TTL:        0,
	Nonce:      8,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     105,
	GasLimit:   1,
	GasPrice:   5,
	CallData:   nil,
}
var txCaseCallAppTx2Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x02,
	/*sign:*/ 0x1a, 0x37, 0x21, 0xe0, 0x20, 0x73, 0x5b, 0xdf, 0x80, 0xc2, 0x8f, 0x7d, 0x48, 0x8c, 0x78, 0x65, 0xcc, 0xc9, 0x96, 0x96, 0x75, 0xfd, 0xbe, 0x14, 0xdd, 0x24, 0x62, 0x15, 0x57, 0xb2, 0xec, 0x0b, 0xde, 0xe2, 0x0e, 0xa5, 0x2a, 0xfb, 0xba, 0x60, 0x4a, 0x04, 0x75, 0x4e, 0x8b, 0xcc, 0x16, 0xd6, 0x7a, 0xc0, 0x89, 0xb1, 0xbd, 0x2b, 0xcb, 0x45, 0xdf, 0xdd, 0xeb, 0x9d, 0x7e, 0xc0, 0xca, 0x00,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx2EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x03,
	/*sign:*/ 0x7b, 0x0c, 0x08, 0x3d, 0x05, 0xd2, 0x8f, 0x2a, 0x5d, 0xd7, 0x4e, 0xb7, 0x15, 0x56, 0x35, 0x08, 0x5c, 0x6c, 0x41, 0xfc, 0x28, 0x35, 0x52, 0x30, 0x08, 0x15, 0xf8, 0xa8, 0x45, 0x49, 0xf6, 0xcc, 0x6b, 0x4f, 0x30, 0x48, 0xbc, 0x28, 0xb2, 0x83, 0x7e, 0xce, 0xcc, 0x53, 0x21, 0x9a, 0x31, 0x98, 0xcf, 0x54, 0x16, 0x10, 0xe4, 0x07, 0x16, 0x0a, 0x0e, 0x5c, 0x05, 0x65, 0x26, 0xc7, 0x12, 0x0f,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/
}
var txValueCallAppTx3 = CallAppTx{
	TTL:        0,
	Nonce:      12,
	AppAddress: Address{0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb},
	Amount:     106,
	GasLimit:   1,
	GasPrice:   33,
	CallData:   nil,
}
var txCaseCallAppTx3Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x02,
	/*sign:*/ 0xa5, 0x29, 0xac, 0x8d, 0xfd, 0xf5, 0xe3, 0x52, 0xb6, 0x18, 0xa1, 0xff, 0x4f, 0xd3, 0x84, 0x38, 0x90, 0xe2, 0xe6, 0x70, 0x41, 0x12, 0xaf, 0x89, 0x4a, 0x0f, 0x3f, 0x41, 0x87, 0x32, 0x99, 0x16, 0x1e, 0x7d, 0x01, 0xc9, 0xb8, 0x95, 0x05, 0x44, 0x02, 0xaf, 0xcf, 0x71, 0x0b, 0xdc, 0x20, 0x37, 0xa0, 0x5a, 0x7c, 0x35, 0x79, 0x77, 0x9f, 0xb6, 0x5d, 0x20, 0xf9, 0xed, 0xc9, 0xec, 0x03, 0x0c,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx3EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x03,
	/*sign:*/ 0xe0, 0x2a, 0x36, 0xf0, 0xb7, 0x71, 0x03, 0xcf, 0xd1, 0x38, 0x06, 0x0c, 0xcd, 0xa5, 0x0f, 0x29, 0x90, 0x4d, 0x80, 0x31, 0x1d, 0xe0, 0xe3, 0xe6, 0x20, 0xf0, 0x92, 0xb2, 0x06, 0x88, 0xb6, 0xf3, 0x14, 0x1d, 0x90, 0x70, 0x42, 0x64, 0xa0, 0x60, 0x70, 0x1f, 0xc6, 0xc9, 0x89, 0x92, 0xe7, 0xdf, 0xeb, 0xa3, 0x21, 0x6e, 0x45, 0x89, 0xe6, 0x69, 0xa1, 0xf4, 0x02, 0xd1, 0x26, 0x3a, 0x98, 0x03,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/
}
var txValueSpawnAppTx1 = SpawnAppTx{
	TTL:        0,
	Nonce:      5,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     107,
	GasLimit:   1,
	GasPrice:   10,
	CallData:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
}
var txCaseSpawnAppTx1Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x04,
	/*sign:*/ 0xbf, 0x1d, 0xfa, 0x61, 0x0a, 0x15, 0x41, 0x56, 0xc7, 0x36, 0x37, 0xe6, 0x44, 0x20, 0xde, 0x8c, 0x27, 0xc1, 0xce, 0x28, 0xc5, 0xdb, 0x34, 0x30, 0x5e, 0xda, 0xe7, 0x50, 0x25, 0xb2, 0x7d, 0x71, 0xcc, 0x74, 0x85, 0x91, 0x45, 0x00, 0x30, 0xf7, 0xfb, 0x60, 0xcb, 0x9c, 0x75, 0x08, 0xa2, 0x81, 0x40, 0xf7, 0xc8, 0x95, 0xf7, 0xea, 0xce, 0x8b, 0xa0, 0xb6, 0x68, 0x1f, 0xf5, 0x6d, 0xa3, 0x05,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx1EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x05,
	/*sign:*/ 0x83, 0x44, 0xc7, 0xa0, 0x20, 0x0d, 0x25, 0x45, 0x5f, 0xea, 0x14, 0x3b, 0x94, 0xc6, 0xe2, 0xc8, 0x8f, 0x14, 0x97, 0x9d, 0xad, 0xa9, 0x3b, 0x5d, 0xa9, 0x51, 0xf3, 0x22, 0x07, 0xc9, 0x71, 0xc2, 0xda, 0xc0, 0x74, 0x86, 0x3b, 0x0d, 0x13, 0x11, 0xb5, 0xfc, 0xdc, 0x56, 0x76, 0x8e, 0x5e, 0xcf, 0x3e, 0xa8, 0xff, 0x23, 0x7e, 0xe3, 0xf0, 0x28, 0xee, 0x4a, 0x4d, 0xf2, 0x5d, 0x72, 0x51, 0x0b,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/
}
var txValueSpawnAppTx2 = SpawnAppTx{
	TTL:        0,
	Nonce:      8,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     108,
	GasLimit:   1,
	GasPrice:   5,
	CallData:   nil,
}
var txCaseSpawnAppTx2Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x04,
	/*sign:*/ 0x35, 0x82, 0x79, 0x53, 0xd2, 0x63, 0x3b, 0x46, 0x9c, 0x82, 0x5d, 0x74, 0x44, 0x56, 0xdc, 0xdd, 0x8f, 0xa6, 0x96, 0xc1, 0xfa, 0x85, 0x17, 0x07, 0x44, 0x8a, 0x41, 0xa9, 0xd7, 0x37, 0x25, 0x3f, 0x33, 0x33, 0xaf, 0x72, 0x9a, 0x08, 0x96, 0x19, 0xf7, 0x79, 0xa8, 0x86, 0x02, 0xd1, 0xa8, 0x08, 0x57, 0x77, 0x12, 0x12, 0xfe, 0xb7, 0x2d, 0x3b, 0xc1, 0x80, 0x78, 0x74, 0x06, 0x5f, 0x65, 0x05,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx2EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x05,
	/*sign:*/ 0x4d, 0xb0, 0x14, 0xd8, 0xc5, 0x69, 0xcc, 0xad, 0x9c, 0x4a, 0x40, 0x6e, 0x35, 0xa0, 0x6d, 0xe3, 0xec, 0x77, 0xdb, 0x07, 0x58, 0xf7, 0x90, 0x44, 0x24, 0x87, 0x94, 0x4c, 0xd5, 0x3f, 0xd2, 0x13, 0x95, 0x33, 0xea, 0x24, 0xbf, 0x0c, 0xe2, 0x74, 0x71, 0xeb, 0x86, 0x2d, 0x74, 0x70, 0xd2, 0xf2, 0x56, 0x89, 0x99, 0x3b, 0x84, 0x0d, 0x2a, 0x48, 0xb7, 0x2a, 0x5e, 0xe8, 0x95, 0x81, 0x30, 0x01,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/
}
var txValueSpawnAppTx3 = SpawnAppTx{
	TTL:        0,
	Nonce:      12,
	AppAddress: Address{0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb},
	Amount:     109,
	GasLimit:   1,
	GasPrice:   33,
	CallData:   nil,
}
var txCaseSpawnAppTx3Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x04,
	/*sign:*/ 0x6e, 0x79, 0xa9, 0xa0, 0xf5, 0x0e, 0xa3, 0x48, 0x68, 0x5e, 0x98, 0xbe, 0x40, 0x34, 0xc8, 0x7a, 0x5a, 0xe7, 0x77, 0xef, 0x74, 0xf3, 0xda, 0x6f, 0x7c, 0x93, 0x34, 0xad, 0x8b, 0x11, 0x82, 0xba, 0xb1, 0xff, 0xda, 0x6f, 0x06, 0xc6, 0xd9, 0xad, 0x42, 0x7e, 0x24, 0x25, 0x0c, 0x46, 0x35, 0x4d, 0x76, 0x8d, 0x9a, 0xda, 0xfe, 0xe0, 0xba, 0xb1, 0x40, 0x05, 0x86, 0xf4, 0xb4, 0x6e, 0x5b, 0x08,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx3EdPlus = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x05,
	/*sign:*/ 0x8b, 0xc0, 0x38, 0x02, 0x4c, 0x9d, 0x38, 0x1a, 0xfd, 0x9b, 0x38, 0x03, 0xc3, 0xcb, 0x9c, 0xe7, 0x56, 0xd2, 0xb0, 0xb8, 0x06, 0xbb, 0x38, 0xc2, 0x55, 0x55, 0x43, 0x64, 0xe7, 0xaf, 0x91, 0xa5, 0x28, 0x62, 0x4a, 0x19, 0x00, 0xce, 0x19, 0x5a, 0x18, 0x5a, 0xba, 0x3e, 0x78, 0x54, 0xeb, 0x1c, 0x99, 0x14, 0x25, 0x76, 0xbd, 0x0c, 0xae, 0xee, 0x20, 0x99, 0x16, 0x67, 0x2d, 0xef, 0x1b, 0x03,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00,
	/*pkey:*/
}
var txCases = []struct {
	Tx     interface{}
	Signed SignedTransaction
}{
	{txValueOldCoinTx1, txCaseOldCoinTx1Ed},
	{txValueOldCoinTx1, txCaseOldCoinTx1EdPlus},
	{txValueOldCoinTx2, txCaseOldCoinTx2Ed},
	{txValueOldCoinTx2, txCaseOldCoinTx2EdPlus},
	{txValueSimpleCoinTx1, txCaseSimpleCoinTx1Ed},
	{txValueSimpleCoinTx1, txCaseSimpleCoinTx1EdPlus},
	{txValueSimpleCoinTx2, txCaseSimpleCoinTx2Ed},
	{txValueSimpleCoinTx2, txCaseSimpleCoinTx2EdPlus},
	{txValueCallAppTx1, txCaseCallAppTx1Ed},
	{txValueCallAppTx1, txCaseCallAppTx1EdPlus},
	{txValueCallAppTx2, txCaseCallAppTx2Ed},
	{txValueCallAppTx2, txCaseCallAppTx2EdPlus},
	{txValueCallAppTx3, txCaseCallAppTx3Ed},
	{txValueCallAppTx3, txCaseCallAppTx3EdPlus},
	{txValueSpawnAppTx1, txCaseSpawnAppTx1Ed},
	{txValueSpawnAppTx1, txCaseSpawnAppTx1EdPlus},
	{txValueSpawnAppTx2, txCaseSpawnAppTx2Ed},
	{txValueSpawnAppTx2, txCaseSpawnAppTx2EdPlus},
	{txValueSpawnAppTx3, txCaseSpawnAppTx3Ed},
	{txValueSpawnAppTx3, txCaseSpawnAppTx3EdPlus},
}

func TestBinaryTransactions(t *testing.T) {
	for _, v := range txCases {
		tx, err := v.Signed.Decode()
		require.NoError(t, err)
		val := reflect.ValueOf(v.Tx)
		b := reflect.New(val.Type())
		ok := tx.Extract(b.Interface())
		require.True(t, ok)
		require.Equal(t, v.Tx, b.Elem().Interface())
	}
}