package main

import (
	"fmt"
	"log"

	// pb "./example.pb"
	pb "example"
	"github.com/golang/protobuf/proto"
)

func main() {
	test := &pb.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		// // Optionalgroup: &pb.Test_OptionalGroup{
		// // 	RequiredField: proto.String("good bye"),
		// // },
		//Union: &pb.Test_Name{"fred"},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	ss := fmt.Sprint(data)
	fmt.Println("data =", ss)
	fmt.Printf("data = %s\n", data)

	newTest := &pb.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	mm := fmt.Sprint(newTest)
	fmt.Println("data2-1 =", *newTest.Label)
	fmt.Println("data2-1 =", *newTest.Type)
	fmt.Println("data2-1 =", newTest.Reps)
	fmt.Println("data2-2 =", mm)
	fmt.Printf("data2-3 = %T\n", newTest)

	// 	// Now test and newTest contain the same data.
	// 	if test.GetLabel() != newTest.GetLabel() {
	// 		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	// 	}
	// 	// Use a type switch to determine which oneof was set.
	// 	switch u := test.Union.(type) {
	// 	case *pb.Test_Number: // u.Number contains the number.
	// 		print(u)
	// 	case *pb.Test_Name: // u.Name contains the string.
	// 	}
	// 	// etc.
}
