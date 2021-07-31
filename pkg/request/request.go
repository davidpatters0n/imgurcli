package request

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/davidPatters0n/imgurcli/pkg/authorization"
)

func Generate(httpMethod string, url string, body io.Reader) ([]byte, error) {
	auth, err := authorization.ReadAuthorization()

	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest(
		httpMethod,
		url,
		body,
	)
	if err != nil {
		return []byte{}, err
	}
	var bearer = "Bearer " + auth.AccessToken
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Imgur CLI")
	request.Header.Add("Authorization", bearer)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	response, err := client.Do(request)

	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return responseBody, nil
}
