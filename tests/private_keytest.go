package tests

import (
	"encoding/json"
	"github.com/revel/revel/testing"
	"io"
	"strings"
)

type PrivateKeyTest struct {
	testing.TestSuite
}

type mnemonicKeyTest struct {
	Key      string `json:"key"`
	Mnemonic string `json:"mnemonic"`
}

type mnemonicKeyJsonParam struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func testParamsAsJsonReader(id, password string) io.Reader {
	a, _ := json.Marshal(&mnemonicKeyJsonParam{id, password})
	return strings.NewReader(string(a))
}

func (t *PrivateKeyTest) TestCreateActionValidationWorksProperly() {
	endPoint := t.BaseUrl() + "/private_key"
	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("", "")).Send()
	t.AssertStatus(422)
	t.AssertContains("Id is required.")
	t.AssertContains("Password is required.")

	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("", "abcd")).Send()
	t.AssertStatus(422)
	t.AssertContains("Id is required.")
	t.AssertContains("Password must be at least 8 characters.")

	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("test", "abcd")).Send()
	t.AssertStatus(422)
	t.AssertNotContains("Id is required.")
	t.AssertContains("Password must be at least 8 characters.")

	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("test", "abcdefgh")).Send()
	t.AssertStatus(200)
	t.AssertNotContains("Id is required.")
	t.AssertNotContains("Password must be at least 8 characters.")

	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("test", "abcdefghasodijfasdlifjadsoifjaodsifjdsaifjadosifjasdoifjsdfdsfsdf")).Send()
	t.AssertStatus(422)
	t.AssertNotContains("Id is required.")
	t.AssertContains("Password must be at most 64 characters.")
}

func (t *PrivateKeyTest) TestCreateActionGereratesKeyAndMneonics() {
	endPoint := t.BaseUrl() + "/private_key"
	
	t.PostCustom(endPoint, "application/json", testParamsAsJsonReader("12345", "abcdefgh")).Send()
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	// Parsing the response from server.
	var response mnemonicKeyTest
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
