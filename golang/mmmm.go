package main

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"math/big"
	//"strconv"
	"encoding/binary"
)

func main() {
	// String of Privatekey
	privateKey := "MCUCAQECICNBHMErbu4gDyG7OVKu6+WuT7dhmgrPHkt4nz1SSAqu"
	xx, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		fmt.Println("ERR0 = ", err)
		return
	}
	fmt.Printf("Base64[%T] = %+v\n", xx, xx)

	// // String to Int64
	// pk_Int64, err := strconv.ParseInt(privateKey, 10, 64)
	// if err != nil {
	// 	fmt.Println("ERR1 = ", err)
	// } else {
	// }
	pk_Int64 := binary.BigEndian.Uint64(xx)
	//pk_Int64, _ := binary.Varint(xx)
	fmt.Printf("Int64[%T] = %+v\n", pk_Int64, pk_Int64)

	// Int64 to Big.Int64
	pk_Big := big.NewInt(int64(pk_Int64))
	fmt.Printf("Big = %T, %+v\n", pk_Big, pk_Big)

	// Big.Int64 to ECDSA struct
	ec_privateKey := ecdsa.PrivateKey{D: pk_Big}
	fmt.Printf("ECDSA = %T, %+v\n", ec_privateKey, ec_privateKey)

	// Generate PublicKey
	fmt.Println("FFF", ec_privateKey.Public())
	//fmt.Println("FFF", ec_privateKey.String())
}
