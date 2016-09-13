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

type FirebaseClient struct {
	FirebaseRef *firego.Firebase
}

func NewFirebaseClient(firebaseApp string, keyFile string) *FirebaseClient {
	fbClient := new(FirebaseClient)
	client := NewOAuthHttpClient(keyFile)
	fbClient.FirebaseRef = firego.New(fmt.Sprintf("https://%s.firebaseio.com", firebaseApp), client)
	return fbClient
}

func NewOAuthHttpClient(keyFile string) *http.Client {
	jsonKey, err := ioutil.ReadFile(keyFile) // or path to whatever name you downloaded the JWT to
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)

	return client
}

func (fbClient *FirebaseClient) GetData(path string) []byte {
	var v map[string]interface{}
	if err := fbClient.FirebaseRef.Child(path).Value(&v); err != nil {
		log.Fatal(err)
	}
	raw, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return raw
}

func (fbClient *FirebaseClient) GetDataAsString(path string) string {
	return fmt.Sprintf("%s", fbClient.GetData(path))
}
