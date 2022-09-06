package proto

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestProto(t *testing.T) {
	req := &HelloRequest{
		Name: "tom",
	}

	reqByte, err := proto.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	newReq := &HelloRequest{}
	if err = proto.Unmarshal(reqByte, newReq); err != nil {
		t.Error(err)
	}
	t.Log(newReq.String())
}
