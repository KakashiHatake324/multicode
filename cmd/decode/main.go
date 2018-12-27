package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sj14/multicode/decode"
)

var (
	verbose bool
	hex     bool
	base64  bool
	proto   bool
	none    bool
)

func main() {
	// init flags
	flag.BoolVar(&hex, "hex", false, "use hex decoding")
	flag.BoolVar(&base64, "base64", false, "use base64 decoding")
	flag.BoolVar(&proto, "proto", false, "use proto decoding")
	flag.BoolVar(&none, "none", false, "disable all decodings")
	flag.BoolVar(&verbose, "v", false, "verbose ouput mode")
	flag.Parse()

	// read program input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalln("failed to read input")
	}

	// Default decoder enables all decodings.
	// Disable all and only enable specified ones.
	// Flags are set to true by default.
	var opts []decode.Option
	opts = append(opts, decode.WithoutAll())

	// Enable specifified decodings.
	if hex {
		opts = append(opts, decode.WithHex())
	}
	if base64 {
		opts = append(opts, decode.WithBase64())
	}
	if proto {
		opts = append(opts, decode.WithProto())
	}

	// decoding
	var (
		decoder = decode.New(opts...)
		result  = input
		enc     decode.Encryption
	)
	for result, enc = decoder.Decode(result); enc != decode.None; result, enc = decoder.Decode(result) {
		logVerbose("- applied decoding '%v':\n%s\n\n", enc, result)
	}

	// check if any kind of decryption was applied
	if bytes.Compare(input, result) == 0 {
		log.Fatalln("failed to decode")
	}

	// output result
	logVerbose("- result:\n")
	fmt.Printf("%v\n", string(result))
}

func logVerbose(format string, v ...interface{}) {
	if !verbose {
		return
	}
	fmt.Printf(format, v...)
}