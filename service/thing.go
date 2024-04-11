package thingsvc

import "github.com/bradrogan/protobuf-example/thing"

type thingRepository interface {
	GetByID(id int) (*Thing, error)
}
