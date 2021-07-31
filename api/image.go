package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/davidPatters0n/imgurcli/pkg/authorization"
	"github.com/davidPatters0n/imgurcli/pkg/request"
	"github.com/davidPatters0n/imgurcli/types"
)

const API_BASE = "https://api.imgur.com/3/image"

func Fetch(id string) types.Image {
	body, err := request.Generate(
		http.MethodGet,
		(fmt.Sprintf("%s/%s", API_BASE, id)),
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	imageData := types.ImageData{}
	if err := json.Unmarshal(body, &imageData); err != nil {
		log.Fatalln(err)
	}
	return imageData.Image
}

func Upload(filepath string) (int, error) {
	auth, err := authorization.ReadAuthorization()

	if err != nil {
		return 401, err
	}

	imageBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 406, err
	}

	var base64Encoding string

	mimeType := http.DetectContentType(imageBytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(imageBytes)
	request, err := http.NewRequest(
		http.MethodPost,
		API_BASE,
		strings.NewReader(base64Encoding),
	)

	if err != nil {
		return 400, err
	}

	var bearer = "Bearer " + auth.AccessToken
	request.Header.Add("User-Agent", "Imgur-CLI")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", bearer)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	response, err := client.Do(request)

	fmt.Println(response)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func Delete(id string) (int, error) {
	resp, err := request.Generate(
		http.MethodDelete,
		fmt.Sprintf("%s/%s/favorite", API_BASE, id),
		nil,
	)

	if err != nil {
		fmt.Println(resp)
		return 400, err
	}

	return 200, nil
}

func Favourite(id string) (int, error) {
	resp, err := request.Generate(
		http.MethodPost,
		fmt.Sprintf("%s/%s/favorite", API_BASE, id),
		nil,
	)

	if err != nil {
		fmt.Println(resp)
		return 500, err
	}

	return 200, nil
}
