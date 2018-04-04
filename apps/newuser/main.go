// newuser is a simple program that posts a new user
// to https://dasc.capstone.ischool.uw.edu/api/v1/users
// and reports the results then exits.

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Pallinder/go-randomdata"
)

// APIURL is the url of the users API that will be posted to.
const APIURL = "https://dasc.capstone.ischool.uw.edu/api/v1/users"

// NewUser defines the structure of a request for
// a new user sign up
type NewUser struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
}

func main() {
	for i := 0; i < 10; i++ {
		payload := &NewUser{
			FirstName:    randomdata.FirstName(randomdata.RandomGender),
			LastName:     randomdata.LastName(),
			Email:        randomdata.Email(),
			Bio:          "some bio",
			Password:     "password",
			PasswordConf: "password",
		}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("error marshalling JSON: %v", err)
		}
		payloadReader := bytes.NewReader(payloadBytes)

		res, err := http.Post(APIURL, "application/json", payloadReader)
		if err != nil {
			log.Fatalf("error posting user: %v", err)
		}
		if res.StatusCode >= 400 {
			log.Fatalf("error response from API: %d:%v", res.StatusCode, res.Status)
		}
		io.Copy(os.Stdout, res.Body)
	}
}
