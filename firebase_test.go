package firebase_helper_test

import (
	app "github.com/ndphu/espresso.helper.firebase"
	"github.com/ndphu/espresso.model"
	"github.com/stretchr/testify/assert"
	//"log"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	TEST_KEY_FILE          = "key.json"
	TEST_FIREBASE_APP_NAME = "test-7a4ff"
)

func GetTestKeyFile() string {
	keyFile, exists := os.LookupEnv("KEY_FILE")
	if !exists {
		keyFile = TEST_KEY_FILE
	}
	return keyFile
}

func GetTestFirebaseAppName() string {
	firebaseAppName, exists := os.LookupEnv("FIREBASE_APP")
	if !exists {
		firebaseAppName = TEST_FIREBASE_APP_NAME
	}
	return firebaseAppName
}

func NewTestClient() *app.FirebaseClient {
	return app.NewFirebaseClient(GetTestFirebaseAppName(), GetTestKeyFile())
}

func TestUpdateStruct(t *testing.T) {
	fbClient := NewTestClient()

	newObjectPath := fmt.Sprintf("testing/%s", strconv.Itoa(time.Now().Nanosecond()))

	event := model.IREvent{}
	event.DeviceId = "00001111"
	event.Event.Button = "SOME_BUTTON"
	event.Event.Code = 11
	event.Event.Remote = "TestRemote"
	event.Event.Repeat = 22

	fbClient.InsertOrUpdateStruct(newObjectPath, event)
	defer fbClient.Delete(newObjectPath)

	eventActual := model.IREvent{}
	err := fbClient.GetValueAsStruct(newObjectPath, &eventActual)
	if err != nil {
		t.Error("Failed to get value as struct")
		t.FailNow()
	}
	assert.Equal(t, event, eventActual, "Object pushed to Firebase is mismatch")
}

func TestUpdateString(t *testing.T) {
	f := NewTestClient()
	randString := strconv.Itoa(time.Now().Nanosecond())
	newObjectPath := fmt.Sprintf("testing/%s", randString)
	expected := "Some data: " + randString
	f.SetValue(newObjectPath, expected)
	defer f.Delete(newObjectPath)

	actual := f.GetValueAsString(newObjectPath)
	assert.Equal(t, expected, actual, "String pushed mismatch")

	nonExists := f.GetValueAsString(randString)
	assert.Empty(t, nonExists, "Fail to get non-exists path")

}
