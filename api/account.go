package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davidPatters0n/imgurcli/pkg/request"
	"github.com/davidPatters0n/imgurcli/types"
	"github.com/spf13/cobra"
)

const ACCOUNT_BASE = "https://api.imgur.com/3/account/me"

func CurrentAccount() types.Account {
	response, err := request.Generate(
		http.MethodGet,
		ACCOUNT_BASE,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	rawAccount := types.RawAccount{}
	if err := json.Unmarshal(response, &rawAccount); err != nil {
		log.Fatalln(err)
	}
	return rawAccount.Account
}

func Settings() types.AccountSetting {
	url := fmt.Sprintf("%s/%s", ACCOUNT_BASE, "settings")
	response, err := request.Generate(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	rawAccountSetting := types.RawAccountSetting{}

	fmt.Println(string(response))

	if err := json.Unmarshal(response, &rawAccountSetting); err != nil {
		log.Fatalln(err)
	}
	return rawAccountSetting.AccountSetting
}

func UpdateSettings(cmd *cobra.Command) bool {
	bio, _ := cmd.Flags().GetString("bio")
	publicImages, _ := cmd.Flags().GetString("public-images")
	albumPrivacy, _ := cmd.Flags().GetString("album-privacy")
	username, _ := cmd.Flags().GetString("username")
	messaging, _ := cmd.Flags().GetBool("messaging")
	showMature, _ := cmd.Flags().GetBool("show-mature")
	newsLetterSubscribed, _ := cmd.Flags().GetBool("newsletter-subscribed")

	reqBody := make(map[string]interface{})

	if bio != "" {
		reqBody["bio"] = bio
	}

	if publicImages != "" {
		reqBody["public_images"] = publicImages
	}

	if albumPrivacy != "" {
		reqBody["album_privacy"] = albumPrivacy
	}

	if username != "" {
		reqBody["username"] = username
	}

	reqBody["messaging"] = messaging
	reqBody["show_mature"] = showMature
	reqBody["newsletter_subscribed"] = newsLetterSubscribed

	body, _ := json.Marshal(reqBody)
	url := fmt.Sprintf("%s/%s", ACCOUNT_BASE, "settings")
	response, err := request.Generate(
		http.MethodPost,
		url,
		strings.NewReader(string(body)),
	)

	if err != nil {
		log.Fatalln(err)
	}
	jsonResponse := types.JsonResponse{}
	if err := json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		return jsonResponse.Success
	}
	return jsonResponse.Success
}

func SendVerificationEmail() bool {
	url := fmt.Sprintf("%s/%s", ACCOUNT_BASE, "verifyemail")
	response, err := request.Generate(
		http.MethodPost,
		url,
		nil,
	)

	if err != nil {
		fmt.Println(response)
		log.Fatalln(err)
	}

	jsonResponse := types.JsonResponse{}

	if err := json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		return false
	}
	return jsonResponse.Success
}

func AccountVerified() bool {
	url := fmt.Sprintf("%s/%s", ACCOUNT_BASE, "verifyemail")
	response, err := request.Generate(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		fmt.Println(response)
		log.Fatalln(err)
	}
	jsonResponse := types.JsonResponse{}
	if err := json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		return jsonResponse.Success
	}
	return jsonResponse.Success
}
