package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	xx, err := base64.StdEncoding.DecodeString("MCUCAQECICNBHMErbu4gDyG7OVKu6+WuT7dhmgrPHkt4nz1SSAqu")
	fmt.Println("ERR = ", err)
	fmt.Printf("Result = %+v", xx)

	//x2, err := base64.StdEncoding.DecodeString("ZUVvUvCScFUAWRk3E3gwIWnsOaFNDUXT2C4x6Wl3GTMsnWTk1gaFny4G582NTpzqWGR755zCQ5kDA8YyMl3R2w==")	// stateHash
	//x2, err := base64.StdEncoding.DecodeString("CAI=")                     //consensusMetadata
	//x2, err := base64.StdEncoding.DecodeString("Eg9nYW1lcGFpY29yZV92MDE=") //ChaincodeID
	//x2, err := base64.StdEncoding.DecodeString("+SuQi13GO8bGYik2SQOIeRcRCiH21NUVH7uSRtMTtq/td0EaKTw/jTQdPvICWq2Gn+klrJ2Hs0DqRYukPVUyow==") //'previousBlockHash'
	//fmt.Println("X2 = ", string(x2))

	//  //p23, err := base64.StdEncoding.DecodeString("CiI1WmFQWGZKYUxOckduWHV5WHVuRkU0eEt4YWtFemdUSVpREhxGcmkgQXVnIDE3IDExOjA5OjI1IENTVCAyMDE4IkgIARJECiDii7hQ6FHfEscWWJxkV3gcKZphiXnt3+2vbJ3hUqkL5hIg7t8wD3SjhMoFOiIb54UUP0aLpNwrEXiflLDgxdWjS9s=")
	//  p23, err := base64.StdEncoding.DecodeString("8CiYIARIiNVpGVlZQNDdSZjVqLWs3TG9pUmNOb3psYzhkeW5iUFluZw==")
	//  //p23, err := base64.StdEncoding.DecodeString("CkQKIGnwE1MpqraCAvg2mxnxsnX5T+iNbC0THejYY1QXB/eXEiAD4FZQ25G3OjxRCPPJDDdcLQXA/180ykUHW4TB0Fs+Vg==")

	//  c321 := fmt.Sprint(p23)

	// Transaction Signature
	//p23, err := base64.StdEncoding.DecodeString("MEUCIQDYAShn5zmsl6DR7SQkNUMsgzzfjgCDAJaG2SH0m1/CswIgMbWJ0x0MzlS6GzYhNRohAT05udstGqS0KMsCvJVrwb0=")
	//fmt.Println("TEST = ", c321)
	//fmt.Printf("TEST = %s\n", p23)

	//v1, err := base64.StdEncoding.DecodeString("CiI1WmFQWGZKYUxOckduWHV5WHVuRkU0eEt4YWtFemdUSVpREhxGcmkgQXVnIDE3IDExOjA5OjI1IENTVCAyMDE4IkgIARJECiDii7hQ6FHfEscWWJxkV3gcKZphiXnt3+2vbJ3hUqkL5hIg7t8wD3SjhMoFOiIb54UUP0aLpNwrEXiflLDgxdWjS9s=")
}
