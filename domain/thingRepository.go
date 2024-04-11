package thing

import (
	"errors"
)

type MockThingRepository struct {
	things []Thing
}

func (r *MockThingRepository) GetByID(id int) (*Thing, error) {
	if id >= len(r.things) {
		return nil, errors.New("invalid Id")
	}

	return &r.things[id], nil
}

func (r *MockThingRepository) Save(new Thing) Thing {
	new.Id = len(r.things)
	r.things = append(r.things, new)

	return new
}

func NewMockThingRepository() *MockThingRepository {
	things := []Thing{
		{
			Id:         0,
			ThingOne:   "Zero",
			OtherStuff: "Other Zero",
		},
		{
			Id:         1,
			ThingOne:   "One",
			OtherStuff: "Other One",
		},
	}
	return &MockThingRepository{
		things: things,
	}
}
