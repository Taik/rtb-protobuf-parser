package parser

import (
	"encoding/json"
	"errors"

	"github.com/taik/rtb-protobuf-parser/proto_adx"
	"github.com/taik/rtb-protobuf-parser/proto_openx"
)

type ProtoMessager interface {
	Unmarshal([]byte) error
	Marshal() ([]byte, error)
}

var messageTypes = []ProtoMessager{
	&adx_rtb.BidRequest{},
	&adx_rtb.BidResponse{},
	&openx_rtb.BidRequest{},
	&openx_rtb.BidResponse{},
}

func Decode(b []byte) ([]byte, error) {
	for _, msg := range messageTypes {
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
