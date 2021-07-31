package authorization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const API_ENDPOINT = "https://api.imgur.com"

type Authorization struct {
	AccessToken     string      `json:"access_token"`
	ExpiresIn       int         `json:"expires_in"`
	TokenType       string      `json:"token_type"`
	Scope           interface{} `json:"scope"`
	RefreshToken    string      `json:"refresh_token"`
	AccountID       int         `json:"account_id"`
	AccountUsername string      `json:"account_username"`
}

var IMGUR_CLIENT = os.Getenv("IMGUR_CLIENT")
var IMGUR_SECRET = os.Getenv("IMGUR_SECRET")

func Authorize() Authorization {
	// 1. Prompt for pin
	// 1.1 we'll want to check if the client & secret is set if not then we need to throw an error
	authorizeLink := fmt.Sprintf("https://api.imgur.com/oauth2/authorize?client_id=%v&response_type=pin", IMGUR_CLIENT)
	fmt.Println("Generate your pin from: " + authorizeLink)
	fmt.Println("Please enter your pin:")
	var pin string

	// Disable printing the time, source file & line number
	log.SetFlags(0)

	_, err := fmt.Scanln(&pin)
	if err != nil {
		log.Fatalln("No pin entered, please enter your pin")
	}

	// 1. Encode the JSON data
	postBody, _ := json.Marshal(map[string]string{
		"client_id":     IMGUR_CLIENT,
		"client_secret": IMGUR_SECRET,
		"grant_type":    "pin",
		"pin":           pin,
	})

	// 2. Take the encoded JSON data and convert it to byteData
	responseBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%v/oauth2/token", API_ENDPOINT)
	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An error occured %v", err)
	}

	// 3. Response body should be closed in order to prevent memory leaks
	defer resp.Body.Close()

	// 4. Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	authorization := Authorization{}

	if err := json.Unmarshal([]byte(body), &authorization); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}
	file, _ := json.Marshal(authorization)
	_ = ioutil.WriteFile("/tmp/imgur_authorization.json", file, 0644)

	return authorization
}

func ReadAuthorization() (Authorization, error) {
	contents, fileErr := ioutil.ReadFile("/tmp/imgur_authorization.json")
	authorization := Authorization{}

	if err := json.Unmarshal([]byte(contents), &authorization); err != nil {
		return authorization, err
	}

	if fileErr != nil {
		return authorization, fileErr
	}

	return authorization, nil
}
