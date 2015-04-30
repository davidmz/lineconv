package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var encList = map[string]encoding.Encoding{
	"utf-8":  nil,
	"cp1251": charmap.Windows1251,
	"koi8-r": charmap.KOI8R,
}

func main() {
	var (
		fromEncName string
		toEncName   string
		printHelp   bool
		printList   bool
	)

	flag.StringVar(&fromEncName, "f", "utf-8", "input encoding name")
	flag.StringVar(&toEncName, "t", "utf-8", "output encoding name")
	flag.BoolVar(&printHelp, "h", false, "print help")
	flag.BoolVar(&printList, "l", false, "print encoding list")
	flag.Parse()

	if printHelp {
		flag.Usage()
		return
	}

	if printList {
		for k := range encList {
			fmt.Println(k)
		}
		return
	}

	fromEnc, ok := encList[strings.ToLower(fromEncName)]
	if !ok {
		fmt.Printf("Unknown encoding %q\n", fromEncName)
	}

	toEnc, ok := encList[strings.ToLower(toEncName)]
	if !ok {
		fmt.Printf("Unknown encoding %q\n", toEncName)
	}

	fromTr, toTr := transform.Nop, transform.Nop

	if fromEnc != toEnc {
		if fromEnc != nil {
			fromTr = fromEnc.NewDecoder()
		}
		if toEnc != nil {
			toTr = toEnc.NewEncoder()
		}
	}

	io.Copy(
		transform.NewWriter(os.Stdout, toTr),
		transform.NewReader(os.Stdin, fromTr),
	)
}
