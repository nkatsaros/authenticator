package main

import (
	"github.com/nkatsaros/authenticator/google"

	"flag"
	"log"
)

var secretFlag = flag.String("secret", "", "Your authenticator secret")

func main() {
	flag.Parse()
	log.SetFlags(0)

	if *secretFlag == "" {
		flag.Usage()
		log.Fatalln("The secret flag is required")
	}

	code, err := google.Code(*secretFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(code)
}
