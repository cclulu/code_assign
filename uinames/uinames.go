package uinames

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cclulu/code_assign/utils"
	"github.com/pkg/errors"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func GetRandomPerson() (Person, error) {
	person := Person{}
	url := fmt.Sprintf("http://uinames.com/api/")
	// url := fmt.Sprintf("http://localhost:8000/")

	resp, err := http.Get(url)
	if err != nil {
		return person, errors.Wrap(err, "failed to get random person")
	}

	if resp.StatusCode > 399 {
		return person, errors.Wrap(utils.StatusError(resp), "Random person status code is greater than 399")
	}

	err = json.NewDecoder(resp.Body).Decode(&person)

	return person, err
}
