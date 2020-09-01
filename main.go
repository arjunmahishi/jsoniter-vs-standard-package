package main

import (
	"encoding/json"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var (
	small = "small.json"
	med   = "med.json"
	large = "large.json"
)

func main() {
	bytes := readFile(small)

	// encoding/json
	j := parseJSON(bytes)
	marshalJSON(j)

	// jsoniter
	k := parseJSONIter(bytes)
	marshalJSONIter(k)
}

func readFile(fileName string) []byte {
	reader, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return reader
}

func parseJSON(raw []byte) json.RawMessage {
	var j json.RawMessage
	if err := json.Unmarshal(raw, &j); err != nil {
		panic(err)
	}
	return j
}

func marshalJSON(j json.RawMessage) {
	if _, err := json.Marshal(j); err != nil {
		panic(err)
	}
}

func parseJSONIter(raw []byte) json.RawMessage {
	var j json.RawMessage
	if err := jsoniter.Unmarshal(raw, &j); err != nil {
		panic(err)
	}
	return j
}

func marshalJSONIter(j json.RawMessage) {
	if _, err := jsoniter.Marshal(j); err != nil {
		panic(err)
	}
}
