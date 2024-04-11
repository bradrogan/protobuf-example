package thingsvc

import (
	fakeproto "github.com/bradrogan/protobuf-example/clients"
	thing "github.com/bradrogan/protobuf-example/domain"
)

type thingRepository interface {
	GetByID(id int) (*thing.Thing, error)
	Save(new thing.Thing) thing.Thing
}

type thingClient interface {
	Fetch() thing.Thing
}

type ThingService struct {
	repo   thingRepository
	client thingClient
}

func (s ThingService) CreateThing() thing.Thing {

	//our fake client gets some data from another service via protobuf
	t := s.client.Fetch()

	//lets say our service adds the OtherStuff, our "business logic"
	t.OtherStuff = "blah"

	return s.repo.Save(t)
}

func (s ThingService) GetThing(id int) (*thing.Thing, error) {

	t, error := s.repo.GetByID(id)
	if error != nil {
		return nil, error
	}
	return t, nil
}

func New(r thingRepository, c fakeproto.Fakeproto) *ThingService {
	return &ThingService{
		repo:   r,
		client: c,
	}
}
