package firebase_helper

import (
	"encoding/json"
	"fmt"
	"github.com/zabawaba99/firego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
)

func NewFirebaseClient(keyFile string) *http.Client {
	jsonKey, err := ioutil.ReadFile(keyFile) // or path to whatever name you downloaded the JWT to
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)

	return client
}

func GetData(url string, client *http.Client) []byte {
	fb := firego.New(url, client)
	var v map[string]interface{}
	if err := fb.Value(&v); err != nil {
		log.Fatal(err)
	}
	raw, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return raw
}

func GetDataAsString(url string, client *http.Client) string {
	return fmt.Sprintf("%s", GetData(url, client))
}
