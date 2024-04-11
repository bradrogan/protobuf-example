package main

import (
	"fmt"

	fakeproto "github.com/bradrogan/protobuf-example/clients"
	thing "github.com/bradrogan/protobuf-example/domain"
	"github.com/bradrogan/protobuf-example/dummy"
	thingsvc "github.com/bradrogan/protobuf-example/service"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Make some dummy data that we will pretend to pull from a proto
	// We don't care about ThingTwo, and this shows we can decouple domain objects from external protos
	d := dummy.Data{
		ThingOne: "Hello!",
		ThingTwo: "World!",
	}

	// create the raw proto to pretend we are ingesting it later
	buf, _ := proto.Marshal(&d)

	// create our service and inject it's dependencies, including our fake proto client of another service
	thingsvc := thingsvc.New(thing.NewMockThingRepository(), fakeproto.New(buf))

	// fetch one of our things, note they don't have knowledge of ThingTwo
	fetchedThing, err := thingsvc.GetThing(0)
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	v := *fetchedThing

	line := "\n---------------\n"

	fmt.Print("Fetched a thing:", line, v.ToString(), line)

	// now we do a typical service operation and create a new thing, which involves fetching from the fake proto client
	// hop to the service code to see how it doesn't care about ThingTwo, and adding ThingThree wouldn't be a breaking change
	createdThing := thingsvc.CreateThing()

	fmt.Print("\n\nCreated a thing:", line, createdThing.ToString(), line)

}
