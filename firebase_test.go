package firebase_helper_test

import (
	app "github.com/ndphu/espresso.helper.firebase"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	FirebaseTestApp = "test-7a4ff"
)

func TestNewFirebaseClient(t *testing.T) {
	keyFile, exists := os.LookupEnv("KEY_FILE")
	if !exists {
		t.Error("KEY_FILE is undefined")
		return
	}

	fbClient := app.NewFirebaseClient(FirebaseTestApp, keyFile)

	expected := `{"device":{"id":"0","name":"test"},"schema":"1.0","server":{"firebase":{"appName":"test-7a4ff"},"mqtt":{"host":"19november.ddns.net","port":5384,"protocol":"tcp"}}}`
	actual := fbClient.GetDataAsString("config/0")
	assert.Equal(t, expected, actual, "Data retreived is incorrect")

	actual = fbClient.GetDataAsString("config/not_exists")
	assert.Equal(t, "null", actual, "Data retreived is incorrect")

}
