package models

import (
	"fmt"
	"sync"
)

type Document struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	mu sync.Mutex
}

var MockedDocumentsTable = []*Document{}

func init() {
	document1 := Document{
		Id: 1, Name: "First document", Content: "Document 1 initial content"}
	document2 := Document{Id: 2, Name: "Second document"}
	document3 := Document{Id: 3, Name: "Third document"}

	MockedDocumentsTable = append(MockedDocumentsTable, &document1)
	MockedDocumentsTable = append(MockedDocumentsTable, &document2)
	MockedDocumentsTable = append(MockedDocumentsTable, &document3)
}

func GetDocumentById(id int) (*Document, error) {
	for _, document := range MockedDocumentsTable {
		if document.Id == id {
			return document, nil
		}
	}
	return &Document{}, fmt.Errorf("there is no document with such id")
}

func (document *Document) SetContent(newContent string) {
	document.mu.Lock()
	defer document.mu.Unlock()

	document.Content = newContent
}

func (document *Document) GetContent() string {
	document.mu.Lock()
	defer document.mu.Unlock()

	return document.Content
}