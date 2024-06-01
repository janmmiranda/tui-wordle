package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const wordleUrl = "https://wordle-api.vercel.app/api/wordle"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) CheckWord(word string) (Wordle, error) {
	if len(word) > 5 || len(word) < 5 {
		return Wordle{}, errors.New("guess must be 5 characters")
	}

	requestBody := GuessRequest{Guess: word}
	wordJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("error occured marshalling request: %v", err)
		return Wordle{}, err
	}

	req, err := http.NewRequest("POST", wordleUrl, bytes.NewBuffer(wordJSON))
	if err != nil {
		log.Fatalf("error occured making request: %v", err)
		return Wordle{}, nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("error occured executing request: %v", err)
		return Wordle{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error occured reading response: %v", err)
		return Wordle{}, err
	}

	checkedWordle := Wordle{}
	err = json.Unmarshal(data, &checkedWordle)
	if err != nil {
		log.Fatalf("error occured unmarshalling response: %v \nresponse: %v\nrequest: %v", err, resp.Body, req)
		return Wordle{}, err
	}

	return checkedWordle, nil
}
