package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chrisbradleydev/go-io/pkg/reader"
	"github.com/chrisbradleydev/go-io/pkg/writer"
)

func main() {
	s := "the quick brown fox jumped over the fence"
	sr := strings.NewReader(s)

	lettersCount, err := reader.CountLetters(sr)
	if err != nil {
		log.Fatal(err)
	}

	writeCount, err := writer.WriteString(s, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(writeCount)
	fmt.Println(lettersCount)
}
