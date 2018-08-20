package main

import (
	//"crypto/sha256"
	"encoding/base64"
	"fmt"
	//"crypto/x509"
	// "encoding/pem"
	//"hash"

	//	tr "gamecenter.mobi/paicode/transactions"
	pb "gamecenter.mobi/paicode/protos"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos"
	//"github.com/golang/protobuf/proto/proto3_proto"
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/protos/common"
	//"github.com/hyperledger/fabric/protos/peer"
	//"github.com/davecgh/go-spew/spew"
)

//var _ hash.Hash
//var _ proto3_proto.Message
//var _ proto.Unmarshaler

//var _ peer.Endorsement
//var _ common.Payload

func main() {
	/*
		"payload":"CpwDCAESERIPZ2FtZXBhaWNvcmVfdjAxGuYCCglVU0VSX0ZVTkQKvAFDaUkxV21GUVdHWktZVXhPY2tkdVdIVjVXSFZ1UmtVMGVFdDRZV3RGZW1kVVNWcFJFaHhHY21rZ1FYVm5JREUzSURFeE9qQTVPakkxSUVOVFZDQXlNREU0SWtnSUFSSkVDaURpaTdoUTZGSGZFc2NXV0p4a1YzZ2NLWnBoaVhudDMrMnZiSjNoVXFrTDVoSWc3dDh3RDNTamhNb0ZPaUliNTRVVVAwYUxwTndyRVhpZmxMRGd4ZFdqUzlzPQo4Q2lZSUFSSWlOVnBHVmxaUU5EZFNaalZxTFdzM1RHOXBVbU5PYjNwc1l6aGtlVzVpVUZsdVp3PT0KYENrUUtJR253RTFNcHFyYUNBdmcybXhueHNuWDVUK2lOYkMwVEhlallZMVFYQi9lWEVpQUQ0RlpRMjVHM09qeFJDUFBKRERkY0xRWEEvMTgweWtVSFc0VEIwRnMrVmc9PUIMUGFpQWRtaW5Sb2xlQg5QYWlBZG1pblJlZ2lvbg=="
		"txid":"26fd4378-d55a-4ec9-a93d-dd60ac54e6b9"
	*/

	/*

		chaincodeSpec:<	type:GOLANG
						chaincodeID:<name:"gamepaicore_v01" >
						ctorMsg:<args:"USER_FUND"
								 args:"CiI1WmFQWGZKYUxOckduWHV5WHVuRkU0eEt4YWtFemdUSVpREhxGcmkgQXVnIDE3IDExOjA5OjI1IENTVCAyMDE4IkgIARJECiDii7hQ6FHfEscWWJxkV3gcKZphiXnt3+2vbJ3hUqkL5hIg7t8wD3SjhMoFOiIb54UUP0aLpNwrEXiflLDgxdWjS9s="		// more addr else
								 args:"CiYIARIiNVpGVlZQNDdSZjVqLWs3TG9pUmNOb3psYzhkeW5iUFluZw=="	// addr:xcluo
								 args:"CkQKIGnwE1MpqraCAvg2mxnxsnX5T+iNbC0THejYY1QXB/eXEiAD4FZQ25G3OjxRCPPJDDdcLQXA/180ykUHW4TB0Fs+Vg==" > // Unkown
						attributes:"PaiAdminRole"
						attributes:"PaiAdminRegion" >

		ctorMsg_args := []slice{
			Fund_Type
			&pb.UserTxHeader{FundId: userid, Nounce: m.Nounce}
			&pb.Fund
			&pb.Signature{P: &pb.ECPoint{rx.Bytes(), ry.Bytes()}}

		}
	*/

	var payload string
	payload = "CpwDCAESERIPZ2FtZXBhaWNvcmVfdjAxGuYCCglVU0VSX0ZVTkQKvAFDaUkxV21GUVdHWktZVXhPY2tkdVdIVjVXSFZ1UmtVMGVFdDRZV3RGZW1kVVNWcFJFaHhHY21rZ1FYVm5JREUzSURFeE9qQTVPakkxSUVOVFZDQXlNREU0SWtnSUFSSkVDaURpaTdoUTZGSGZFc2NXV0p4a1YzZ2NLWnBoaVhudDMrMnZiSjNoVXFrTDVoSWc3dDh3RDNTamhNb0ZPaUliNTRVVVAwYUxwTndyRVhpZmxMRGd4ZFdqUzlzPQo4Q2lZSUFSSWlOVnBHVmxaUU5EZFNaalZxTFdzM1RHOXBVbU5PYjNwc1l6aGtlVzVpVUZsdVp3PT0KYENrUUtJR253RTFNcHFyYUNBdmcybXhueHNuWDVUK2lOYkMwVEhlallZMVFYQi9lWEVpQUQ0RlpRMjVHM09qeFJDUFBKRERkY0xRWEEvMTgweWtVSFc0VEIwRnMrVmc9PUIMUGFpQWRtaW5Sb2xlQg5QYWlBZG1pblJlZ2lvbg=="

	// Base64 解码
	payloadProtobuf, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		// cc := fmt.Sprintf("%q", payloadProtobuf)
		//fmt.Println("A = ", fmt.Sprintf("%s", payloadProtobuf))
		fmt.Println(err)
	} else {
		fmt.Printf("A %T = %s\n", payloadProtobuf, payloadProtobuf)
	}

	//cpp := &peer.ChaincodeProposalPayload{}
	//cpp := &peer.ChaincodeActionPayload{}
	//cpp := &peer.ChaincodeEndorsedAction{}
	//cpp := &peer.Endorsement{}
	//cpp := &common.Payload{}
	//cpp := &protos.Transaction{}
	//cpp := &protos.TransactionResult{}
	//cpp := &protos.ChaincodeSpec{}
	//cpp := &protos.ChaincodeDeploymentSpec{}

	cpp := &protos.ChaincodeInvocationSpec{}
	proto.Unmarshal(payloadProtobuf, cpp)
	fmt.Printf("EMC[%T] = %+v\n", cpp, cpp.ChaincodeSpec) //fmt.Println("HAHAH2 =", cpp.ChaincodeSpec.String())
	//for _, v := range cpp.ChaincodeSpec.CtorMsg.Args {
	//	fmt.Printf("EMC-1 = %s\n", v)
	//	//for _, e := range v {
	//	//	fmt.Printf("\tEMC-2 = %+v\n", e)
	//	//}
	//}
	//spew.Dump(cpp)
	ctorMsg_args := cpp.ChaincodeSpec.CtorMsg.Args
	fmt.Println("Jie-[FundType]=", string(ctorMsg_args[0]))
	v1, err := base64.StdEncoding.DecodeString(string(ctorMsg_args[1]))
	//fmt.Printf("Jie-1 = %s \n", v1)
	cdd31 := &pb.UserTxHeader{}
	proto.Unmarshal(v1, cdd31)
	fmt.Println("Jie-1[pb.UserTxHeader] = ", cdd31)
	fmt.Println("Jie-1[pb.UserTxHeader] = ", cdd31.FundId)
	fmt.Println("Jie-1[pb.UserTxHeader] = ", string(cdd31.Nounce))
	fmt.Println("Jie-1[pb.UserTxHeader] = ", cdd31.Ts)
	//fmt.Println("\n")
	/* Fund */
	v1, err = base64.StdEncoding.DecodeString(string(ctorMsg_args[2]))
	//fmt.Printf("Jie-2 = %s \n", v1)
	cdd32 := &pb.Fund{}
	proto.Unmarshal(v1, cdd32)
	fmt.Println("Jie-2[pb.Fund] = ", cdd32)
	fmt.Println("Jie-2[pb.Fund.InvokeChaincode] = ", cdd32.InvokeChaincode)
	fmt.Println("Jie-2[pb.Fund.D] = ", cdd32.D)
	userfund := cdd32.GetUserfund()
	fmt.Println("Jie-2[pb.Fund.D.Pai] = ", userfund.Pai)
	fmt.Println("Jie-2[pb.Fund.D.ToUserId] = ", userfund.ToUserId)
	switch cdd32.D.(type) {
	case *pb.Fund_Userfund:
		fmt.Println("Type1")
	case *pb.Fund_Null:
		fmt.Println("Type2")
	case *pb.Fund_Contract:
		fmt.Println("Type3")
	}
	//fmt.Println("\n")
	/* Args[3] - Signature */
	v1, err = base64.StdEncoding.DecodeString(string(ctorMsg_args[3]))
	cdd33 := &pb.Signature{}
	proto.Unmarshal(v1, cdd33)
	fmt.Println("Jie-3[pb.Signature] = ", cdd33)
	fmt.Println("Jie-3[pb.Signature.P] = ", cdd33.P)
	//fmt.Println("Jie-3[pb.Signature] Unmar = ", string(cdd33.P.X))
	//fmt.Println("Jie-3[pb.Signature] Unmar = ", string(cdd33.P.Y))

	/*
		// X.509 Test
		ba := []byte{}
		for _, b := range payloadProtobuf {
			ba = append(ba, byte(b))
		}
		payloadProtobuf, err = x509.ParsePKIXPublicKey(ba)
		payloadProtobufPEMData = payloadProtobuf
		payloadProtobuf, err = x509.ParseCertificates(payloadProtobufPEMData)
		fmt.Printf("Got a %T, with remaining data: %q\n", payloadProtobuf, "bb")
	*/
}
