package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserParams struct {
	Id string `json:"sub"`
}

func Authenticate(AccToken string) *string {
	var response, err = http.Get(fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?access_token=%s", AccToken))

	if err != nil {
		return nil
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var user UserParams
	err = json.Unmarshal(body, &user)

	if err != nil {
		return nil
	}

	return &user.Id
}
