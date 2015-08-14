package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/taik/go-adx-parser/proto"
	"encoding/base64"
	"bytes"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw := scanner.Bytes()

		if len(raw) < 2 {
			log.Fatalln("Invalid bytes")
		}

		// The starting bytes looks like the message is base64, so decode it first
		if bytes.Equal(raw[0:2], []byte("Eh")) {
			var err error
			raw, err = base64.StdEncoding.DecodeString(string(raw))
			if err != nil {
				log.Fatalln("Base64 decode error: %s", err.Error())
			}
		}

		request := &realtime_bidding_proto.BidRequest{}
		err := proto.Unmarshal(raw, request)
		if err != nil {
			log.Fatalf("Error unmarshaling: %s", err.Error())
		}

		json, err := json.Marshal(request)
		if err != nil {
			log.Fatalf("Error serializing to json: %s", err.Error())
		} else {
			fmt.Println(string(json))
		}
	}
}
