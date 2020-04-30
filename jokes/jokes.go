package jokes

import (
	"fmt"
	"net/http"

	"github.com/cclulu/code_assign/icndb"
	"github.com/cclulu/code_assign/uinames"
)

func Build(w http.ResponseWriter, r *http.Request) {

	person, err := uinames.GetRandomPerson()
	if err != nil {
		err := fmt.Errorf("Error on Get Person call: %s", err)
		http.Error(w, err.Error(), 424)
		return
	}

	joke, err := icndb.GetNerdyJokeGivenName(person.Name, person.Surname)
	if err != nil {
		err := fmt.Errorf("Error on Get Joke call: %s", err)
		http.Error(w, err.Error(), 424)
		return
	}
	w.Write([]byte(joke))
}
