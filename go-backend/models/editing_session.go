package models

import (
	"fmt"
	"log"
)

type EditingSession struct {
	UserId          int  `json:"user_id"`
	DocumentId      int  `json:"document_id"`
	CurrentPosition int  `json:"current_position"`
	IsEditingActive bool `json:"is_editing_active"`
}

var MockedEditingSessionsTable = []*EditingSession{
	{UserId: 0, DocumentId: 1, CurrentPosition: 14, IsEditingActive: true},
	{UserId: 2, DocumentId: 1, CurrentPosition: 3, IsEditingActive: false},
	{UserId: 3, DocumentId: 1, CurrentPosition: 3, IsEditingActive: true},
}

func (editingSession *EditingSession) SetCurrentPosition(newPosition int) {
	document, err := GetDocumentById(editingSession.DocumentId)
	if err != nil {
		return
	}
	if newPosition >= 0 && newPosition <= len(document.Content) {
		editingSession.CurrentPosition = newPosition
	}
}

func (editingSession *EditingSession) WriteToDocument(s string) error {
	document, err := GetDocumentById(editingSession.DocumentId)
	if err != nil {
		return err
	}
	log.Printf("%+s\n", document.Content)
	fmt.Printf("string to be inserted: %s\n", s)
	if editingSession.CurrentPosition == len(document.Content) {
		document.Content += s
	}
	document.Content = document.Content[:editingSession.CurrentPosition] + s +
		document.Content[editingSession.CurrentPosition:]

	log.Printf("%+s\n", document.Content)
	return nil
}

func GetEditingSessionByUserIdAndDocumentId(
	userId,
	documentId int,
) (*EditingSession, error) {

	for _, editingSession := range MockedEditingSessionsTable {
		if editingSession.UserId == userId && editingSession.DocumentId == documentId {
			return editingSession, nil
		}
	}
	return &EditingSession{}, fmt.Errorf("editing session not found")
}
