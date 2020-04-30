package icndb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cclulu/code_assign/utils"
	"github.com/pkg/errors"
)

type Joke struct {
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}
type Icndb struct {
	Value Joke `json:"value"`
}

func GetNerdyJokeGivenName(firstname, lastname string) (string, error) {
	icndb := Icndb{}
	baseUrl := "http://api.icndb.com/jokes/random?"
	pathParams := url.PathEscape(fmt.Sprintf("firstName=%s&lastName=%s&limitTo=[nerdy]", firstname, lastname))

	resp, err := http.Get(baseUrl + pathParams)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve random joke")
	}

	if resp.StatusCode > 399 {
		return "", errors.Wrap(utils.StatusError(resp), "Random joke status code is greater than 399")
	}

	err = json.NewDecoder(resp.Body).Decode(&icndb)

	return icndb.Value.Joke, err
}
