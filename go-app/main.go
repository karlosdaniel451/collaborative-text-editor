package main

import (
	"fmt"
	"log"

	"github.com/mb0/lab/ot"
)

func main() {
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

	err := appendStringToDoc("hiii", server.Doc)
	if err != nil {
		log.Fatal(err)
	}

	err = appendStringToDoc("x", server.Doc)
	if err != nil {
		log.Fatal(err)
	}

	printTextDocContent(server.Doc)

	err = removeLastCharacterFromDoc(server.Doc)
	if err != nil {
		log.Fatal(err)
	}

	printTextDocContent(server.Doc)

	client.Apply(ot.Ops{{N: len([]byte(*client.Doc))}, {S: "A"}})
	printTextDocContent(server.Doc)

	client2 := ot.Client{
		Doc:  &sharedDoc,
		Rev:  server.Rev(),
		Wait: make(ot.Ops, 0),
		Buf:  make(ot.Ops, 0),
	}

	client2.Apply(ot.Ops{{N: len([]byte(*client.Doc))}, {S: "B"}})
	printTextDocContent(server.Doc)
}

func printTextDocContent(doc *ot.Doc) {
	fmt.Printf("%s\n", []byte(*doc))
}

func appendStringToDoc(s string, doc *ot.Doc) error {
	operations := ot.Ops{{N: len([]byte(*doc))}}
	for _, byteContent := range []byte(s) {
		operations = append(operations, ot.Op{S: string(byteContent)})
	}
	return doc.Apply(operations)
}

func removeLastCharacterFromDoc(doc *ot.Doc) error {
	operations := ot.Ops{{N: len([]byte(*doc)) - 1}}
	operations = append(operations, ot.Op{N: -1})
	return doc.Apply(operations)
}
