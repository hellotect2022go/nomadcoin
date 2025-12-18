package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffers bytes.Buffer
	encoder := gob.NewEncoder(&buffers)
	HandleErr(encoder.Encode(i))
	return buffers.Bytes()
}

func FromBytes(data []byte, decodeValue interface{}) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(decodeValue))
}
