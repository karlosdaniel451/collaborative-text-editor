package models

import (
	"fmt"
	"log"
	// "regexp"
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

func (editingSession *EditingSession) SetCurrentPosition(newPosition int) error {
	document, err := GetDocumentById(editingSession.DocumentId)
	if err != nil {
		return err
	}
	if newPosition >= 0 && newPosition <= len(document.Content) {
		editingSession.CurrentPosition = newPosition
		return nil
	}

	return fmt.Errorf("invalid cursor position")
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
		// Make cursor position follow the insertion of content.
		editingSession.CurrentPosition += len(s)
		return nil
	}

	// If cursor is not on the last position.

	// Check if editing operation will not affect other editing sessions.
	for _, otherEditingSession := range MockedEditingSessionsTable {
		if editingSession.UserId != otherEditingSession.UserId &&
			editingSession.DocumentId == otherEditingSession.DocumentId &&
			editingSession.CurrentPosition == otherEditingSession.CurrentPosition {

			return fmt.Errorf("editing operation blocked: another User is on the same" +
				" position")
		}
	}

	document.Content = document.Content[:editingSession.CurrentPosition] + s +
		document.Content[editingSession.CurrentPosition:]

	// Make cursor position follow the insertion of content.
	editingSession.CurrentPosition += len(s)

	// Make cursor position of other EditingSessions follow the insertion of content
	for _, otherEditingSession := range MockedEditingSessionsTable {
		if editingSession.UserId != otherEditingSession.UserId &&
			editingSession.DocumentId == otherEditingSession.DocumentId &&
			editingSession.CurrentPosition < otherEditingSession.CurrentPosition {

			otherEditingSession.CurrentPosition += len(s)
		}
	}

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

	deletingRangeStart := editingSession.CurrentPosition - n
	deletingRangeEnd := editingSession.CurrentPosition

	// Check if editing operation will not affect other editing sessions.
	for _, otherEditingSession := range MockedEditingSessionsTable {
		if editingSession.UserId != otherEditingSession.UserId &&
			editingSession.DocumentId == otherEditingSession.DocumentId {

			if otherEditingSession.CurrentPosition >= deletingRangeStart &&
				otherEditingSession.CurrentPosition <= deletingRangeEnd {

				return fmt.Errorf("editing operation blocked: deleting range would" +
					" affect another User")
			}

		}
	}

	// document.Content = document.Content[:len(document.Content)-n]
	// document.Content = document.Content[:editingSession.CurrentPosition-n]
	document.Content = document.Content[:editingSession.CurrentPosition-n] +
		document.Content[editingSession.CurrentPosition:]

	// Make cursor position follow the deleting of content.
	editingSession.CurrentPosition -= n

	// Make cursor position of other EditingSessions follow the deleting of content
	for _, otherEditingSession := range MockedEditingSessionsTable {
		if editingSession.UserId != otherEditingSession.UserId &&
			editingSession.DocumentId == otherEditingSession.DocumentId &&
			editingSession.CurrentPosition < otherEditingSession.CurrentPosition {

			otherEditingSession.CurrentPosition -= n
		}
	}

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

// func GetIndicesOfEditingOpeartion(
// 	userId int,
// 	document *Document,
// 	currentPosition int,
// ) (leftIndex, rightIndex int) {
// 	// Get the current word
// 	re := regexp.MustCompile(`(\w+)`)
// 	// currentWord := re.FindString(document.Content[currentPosition:])

// 	indices := re.FindStringIndex(document.Content)
// 	return indices[0], indices[1]
// }

// func IsEditingOperationBlocked(leftIndex, rightIndex, userId, documentId int) bool {
// 	for _, editingSession := range MockedEditingSessionsTable {
// 		if userId == editingSession.UserId && documentId == editingSession.DocumentId {
// 			if leftIndex
// 		}
// 	}
// }
