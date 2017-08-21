package tests

import (
	"github.com/revel/revel/testing"
)

type ErrorsTest struct {
	testing.TestSuite
}

func (t *ErrorsTest) TestRandomURLsRespondNotFound() {
	t.Get("/")
	t.AssertNotFound()

	t.Get("/asdasd")
	t.AssertNotFound()

	t.Get("/random")
	t.AssertNotFound()
}
