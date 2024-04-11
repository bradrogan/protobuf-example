package main

import (
	"fmt"
	"os"

	"github.com/kfelter/protobuf-example/dummy"
	"google.golang.org/protobuf/proto"
)

func main() {
	e := dummy.Thing{
		ThingOne: "Hello!",
		ThingTwo: "World!",
	}

	buf, _ := proto.Marshal(&e)
	os.WriteFile("events.protobuf", buf, os.ModePerm)
	fmt.Println("marshal events.protobuf")

	events := dummy.Thing{}
	buf, _ = os.ReadFile("events.protobuf")
	proto.Unmarshal(buf, &events)
	fmt.Println("unmarshal protobuf", events.String())

}
