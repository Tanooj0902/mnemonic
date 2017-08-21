package tests

import (
	"github.com/revel/revel/testing"
	"encoding/json"
	"strings"
)

type PrivateKeyTest struct {
	testing.TestSuite
}

type MnemonicKeyTest struct {
	Key      string `json:"key"`
	Mnemonic string `json:"mnemonic"`
}

func (t *PrivateKeyTest) TestCreateActionValidationWorksProperly() {
	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader("{}")).Send()
	t.AssertStatus(422)
	t.AssertContains("Id is required.")
	t.AssertContains("Password is required.")

	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id": "", "password": ""}`,
	)).Send()
	t.AssertStatus(422)
	t.AssertContains("Id is required.")
	t.AssertContains("Password is required.")

	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id": "", "password": "abcd"}`,
	)).Send()
	t.AssertStatus(422)
	t.AssertContains("Id is required.")
	t.AssertContains("Password must be at least 8 characters.")


	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id":"test", "password":"abcd"}`,
	)).Send()
	t.AssertStatus(422)
	t.AssertNotContains("Id is required.")
	t.AssertContains("Password must be at least 8 characters.")

	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id": "test", "password": "abcdefgh"}`,
	)).Send()
	t.AssertStatus(200)
	t.AssertNotContains("Id is required.")
	t.AssertNotContains("Password must be at least 8 characters.")

	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id": "test", "password": "abcdefghasodijfasdlifjadsoifjaodsifjdsaifjadosifjasdoifjsdfaidsfjasiodjfasiodfjasdoifj"}`,
	)).Send()
	t.AssertStatus(422)
	t.AssertNotContains("Id is required.")
	t.AssertContains("Password must be at most 64 characters.")
}

func (t *PrivateKeyTest) TestCreateActionGereratesKeyAndMneonics() {
	t.PostCustom((t.BaseUrl() + "/private_key"), "application/json", strings.NewReader(
		`{"id": "12345", "password": "abcdefgh"}`,
	)).Send()
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	// Parsing the response from server.
	var response MnemonicKeyTest
	err := json.Unmarshal(t.ResponseBody, &response)
	t.Assert(err == nil)

	key := response.Key
	t.Assert(key != "")
	t.Assert(len(key) == 111)

	mnemonic := response.Mnemonic
	t.Assert(mnemonic != "")
	words := strings.Fields(mnemonic)
	t.Assert(len(words) == 24)
}
