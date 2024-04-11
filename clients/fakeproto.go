package fakeproto

import (
	thing "github.com/bradrogan/protobuf-example/domain"
	"github.com/bradrogan/protobuf-example/dummy"
	"google.golang.org/protobuf/proto"
)

type Fakeproto struct {
	buf []byte
}

func (f Fakeproto) Fetch() thing.Thing {
	dummy := dummy.Data{}
	proto.Unmarshal(f.buf, &dummy)

	// we only care about ThingOne, and ThingTwo or any new ThingThree won't cause any issue
	return thing.Thing{
		ThingOne:   dummy.ThingOne,
		OtherStuff: "",
	}
}

func New(buf []byte) Fakeproto {
	return Fakeproto{
		buf: buf,
	}
}
