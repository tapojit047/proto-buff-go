package main

import (
	"fmt"
	"github.com/tapojit047/proto-buff-go/github.com/yourusername/proto-buff-go/personpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello World!")

	tapojit := &personpb.Person{
		Name: "Tapojit",
		Age:  26,
	}

	data, err := proto.Marshal(tapojit)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	newTapojit := &personpb.Person{}
	err = proto.Unmarshal(data, newTapojit)
	fmt.Println(newTapojit)
}
