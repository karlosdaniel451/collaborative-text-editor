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
	{UserId: 0, DocumentId: 1, CurrentPosition: 0, IsEditingActive: true},
	{UserId: 2, DocumentId: 1, CurrentPosition: 0, IsEditingActive: false},
	{UserId: 3, DocumentId: 1, CurrentPosition: 0, IsEditingActive: true},
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
		// If cursor is on the last position.
		document.Content += s
	} else {
		// If cursor is not on the last position.
		document.Content = document.Content[:editingSession.CurrentPosition] + s +
			document.Content[editingSession.CurrentPosition:]
	}

	// Make cursor position follow the insertion of content.
	editingSession.CurrentPosition += len(s)

	return nil
}

func (editingSession *EditingSession) DeleteFromDocument(n int) error {
	document, err := GetDocumentById(editingSession.DocumentId)
	if err != nil {
		return err
	}
	log.Printf("%s\n", document.Content)

	if n > len(document.Content) || n > editingSession.CurrentPosition {
		return fmt.Errorf("insufficient characters to be deleted")
	}

	if n == len(document.Content) {
		document.Content = ""
		// Make cursor position follow the insertion of content.
		editingSession.CurrentPosition = 0
		return nil
	}

	// document.Content = document.Content[:len(document.Content)-n]
	// document.Content = document.Content[:editingSession.CurrentPosition-n]
	document.Content = document.Content[:editingSession.CurrentPosition-n] +
		document.Content[editingSession.CurrentPosition:]

	// Make cursor position follow the insertion of content.
	editingSession.CurrentPosition -= n

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