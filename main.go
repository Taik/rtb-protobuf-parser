package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/taik/rtb-protobuf-parser/parser"
)

func main() {
	log.SetLevel(log.WarnLevel)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw := scanner.Bytes()

		// The starting bytes looks like the message is base64, so decode it first
		if len(raw) >= 2 && bytes.Equal(raw[0:2], []byte("Eh")) {
			log.WithFields(log.Fields{
				"startingBytes": raw[0:2],
			}).Debugln("Detected base64-encoded string")

			var err error
			raw, err = base64.StdEncoding.DecodeString(string(raw))
			if err != nil {
				log.WithFields(log.Fields{
					"msg": err.Error(),
				}).Warnln("Base64 decode error")
			}
		}

		result, err := parser.Decode(raw)
		if err != nil {
			log.Warnln(err.Error())
		} else {
			println(string(result))
		}
	}
}
