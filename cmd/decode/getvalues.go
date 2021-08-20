package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/rafaeltorres324/multicode/decode"
)

func GetValues(input []byte) string {

	if strings.TrimSpace(string(input)) == "" {
		log.Fatalln("empty input")
	}

	var opts []decode.Option

	opts = append(opts, decode.WithBase64())

	var (
		decoder = decode.New(opts...)
		result  = input
		enc     decode.Encoding
	)
	for result, enc = decoder.Decode(result); enc != decode.None; result, enc = decoder.Decode(result) {
		logVerbose("- applied decoding '%v':\n%s\n\n", enc, result)
	}

	if bytes.Equal(input, result) {
		log.Fatalln("failed to decode")
	}
	return string(result)
}

func logVerbose(format string, v ...interface{}) {
	if !verbose {
		return
	}
	fmt.Printf(format, v...)
}
