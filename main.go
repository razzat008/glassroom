package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "client_secret.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Read the token from a local file in root dir
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}               // struct  to store credentials
	err = json.NewDecoder(f).Decode(tok) // decoding the contents of the json into tok
	return tok, err
}

func main() {
	fmt.Println("Hello, World!")
}

// Incase token from local file doesn't work
func getTokenFromWeb(token *oauth2.Config) *oauth2.Token{

}

// Save received token from the web into local file
func saveToken(path string, token *oauth2.Token)  {
	fmt.Printf("Saving credentials file to:%s\n",path)

}
