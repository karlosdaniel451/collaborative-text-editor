package main

import (
	"fmt"
	"time"

	"github.com/mb0/lab/ot"
)

func threadUnsatefyManipulations() {
	initialDocContent := []byte("abc")
	sharedDoc := ot.Doc(initialDocContent)

	server := ot.Server{
		Doc:     &sharedDoc,
		History: make([]ot.Ops, 0),
	}

	client := ot.Client{
		Doc:  &sharedDoc,
		Rev:  server.Rev(),
		Wait: make(ot.Ops, 0),
		Buf:  make(ot.Ops, 0),
	}

	fmt.Printf("%+v\n", *server.Doc)
	fmt.Printf("%+v\n", client)
	printTextDocContent(server.Doc)

	// err := appendStringToDoc("hiii", server.Doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = appendStringToDoc("x", server.Doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	go appendStringToDoc("x", server.Doc)
	go appendStringToDoc("123", server.Doc)

	time.Sleep(time.Second * 10)
	printTextDocContent(server.Doc)
}
