package repos

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/guitarbum722/repos/mocks"
)

var userCases = []struct {
	input      string
	result     int
	shouldFail bool
	err        error
}{
	{
		"testuser",
		16,
		false,
		nil,
	},
	{
		"nonexist",
		-1,
		true,
		errors.New("user does not exist"),
	},
}

func TestUser(t *testing.T) {
	for _, tt := range userCases {
		cl := &mocks.Clienter{}

		cl.On("RepoCount", mock.AnythingOfType("string")).Return(tt.result, tt.err)

		if got, err := cl.RepoCount(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("RepoCount(%v) = %v; expected %v", tt.input, got, tt.shouldFail)
		}
	}
}
