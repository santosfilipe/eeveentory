package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/santosfilipe/eeveentory/aws"
)

var secret aws.Credential

func Test_GetRdsSecret(t *testing.T) {
	jsonData, err := os.Open("testdata/secretsmanager.json")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer jsonData.Close()

	byteValue, _ := ioutil.ReadAll(jsonData)

	err = json.Unmarshal(byteValue, &secret)
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}

	dbuser, dbpassword := aws.ParseSecretsManagerResponse(secret)

	if dbpassword != "rdstestpassword" {
		t.Error("ERROR: Incorrect result: expected rdstestpassword, got ", dbpassword)
	}

	if dbuser != "dba" {
		t.Error("ERROR: Incorrect result: expected dba, got ", dbuser)
	}
}
