package firebase_helper

import (
	"fmt"
	"github.com/fatih/structs"
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

func (f *FirebaseClient) GetValueAsString(path string) string {
	var result string
	f.FirebaseRef.Child(path).Value(&result)
	return result
}

func (f *FirebaseClient) GetValueAsStruct(path string, v interface{}) error {
	return f.FirebaseRef.Child(path).Value(v)
}

func (f *FirebaseClient) SetValue(path string, value string) error {
	return f.FirebaseRef.Child(path).Set(value)
}

func (f *FirebaseClient) InsertOrUpdateStruct(path string, v interface{}) error {
	return f.FirebaseRef.Child(path).Update(structs.Map(v))
}

func (f *FirebaseClient) InsertOrUpdateString(path string, dataAsString string) error {
	return f.FirebaseRef.Child(path).Set(dataAsString)
}

func (f *FirebaseClient) Delete(path string) error {
	return f.FirebaseRef.Child(path).Remove()
}
