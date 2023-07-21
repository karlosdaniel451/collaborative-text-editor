package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const baseUrl string = "http://localhost:8080"

func Test_multiple_users_editing_same_document() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		setPositionOfUserInDocument(0, 1, 4)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setPositionOfUserInDocument(2, 1, 9)
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToDocument(0, 1, "abc")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToDocument(2, 1, "xyz")
	}()

	wg.Wait()
}

func setPositionOfUserInDocument(userId, documentId, newPosition int) {
	reqBody, _ := json.Marshal(map[string]any{
		"current_position": newPosition,
	})
	reqBodyReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(
		http.MethodPut,
		baseUrl+fmt.Sprintf("/editing-sessions/%d/%d", userId, documentId),
		reqBodyReader,
	)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status: %s\n", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("response body:\n%s\n", respBody)
}

func writeToDocument(userId, documentId int, s string) {
	reqBody := []byte(s)
	reqBodyReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(
		http.MethodPost,
		baseUrl+fmt.Sprintf("/editing-sessions/%d/%d", userId, documentId),
		reqBodyReader,
	)
	req.Header.Add("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status: %s\n", resp.Status)
	if resp.StatusCode != http.StatusNoContent {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("response body:\n%s\n", respBody)
	}
}
