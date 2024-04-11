package thing

import (
	"strconv"
)

type Thing struct {
	Id         int
	ThingOne   string
	OtherStuff string
}

func (t Thing) ToString() string {
	return "Id: " + strconv.Itoa(t.Id) + "\nThingOne: " + t.ThingOne + "\nOtherStuff: " + t.OtherStuff
}
