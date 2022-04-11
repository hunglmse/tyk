package main

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/TykTechnologies/tyk/analytics"
	"github.com/buger/jsonparser"
)

func ExampleMyAnalyticsPluginMaskJSONLoginBody() {
	record := analytics.Record{
		ContentLength: 72,
		RawRequest:    base64.StdEncoding.EncodeToString([]byte("POST / HTTP/1.1\r\nHost: server.com\r\nContent_Length: 72\r\n\r\n{\"email\": \"m\", \"password\": \"p\", \"data\": {\"email\": \"m\", \"password\": \"p\"}}")),
	}
	MyAnalyticsPluginMaskJSONLoginBody(&record)
	data, _ := base64.StdEncoding.DecodeString(record.RawRequest)
	const endOfHeaders = "\r\n\r\n"
	paths := [][]string{
		{"email"},
		{"password"},
		{"data", "email"},
		{"data", "password"},
	}
	if i := bytes.Index(data, []byte(endOfHeaders)); i > 0 || (i+4) < len(data) {
		jsonparser.EachKey(data[i+4:], func(_ int, v []byte, _ jsonparser.ValueType, _ error) {
			fmt.Println(string(v))
		}, paths...)
	}
	// Output: ****
	//****
	//****
	//****
}
