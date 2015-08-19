package parser

import (
	"encoding/json"
	"errors"

	"github.com/taik/go-adx-parser/proto_adx"
	"github.com/taik/go-adx-parser/proto_openx"
)

type ProtoMessager interface {
	Unmarshal([]byte) error
	Marshal() ([]byte, error)
}

func Decode(b []byte) ([]byte, error) {
	types := []ProtoMessager{
		&adx_rtb.BidRequest{},
		&adx_rtb.BidResponse{},
		&openx_rtb.BidRequest{},
		&openx_rtb.BidResponse{},
	}

	for _, msg := range types {
		if err := msg.Unmarshal(b); err != nil {
			continue
		}

		if json, err := json.Marshal(msg); err != nil {
			continue
		} else {
			return json, nil
		}
	}
	return nil, errors.New("Unable to decode to a supported Protobuf Message")
}
