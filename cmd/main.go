package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/defkoi/passgen"
)

var (
	length int
	script string
)

func init() {
	flag.IntVar(&length, "l", 8, "password length")
	flag.StringVar(&script, "s", "", "custom script")
	flag.Parse()
}

func init() {
	if script != "" {
		source, err := os.ReadFile(script)
		if err != nil {
			log.Fatal(err)
		}
		passgen.SetScript(string(source))
	}
}

func main() {
	if v, err := passgen.GeneratePassword(length); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
