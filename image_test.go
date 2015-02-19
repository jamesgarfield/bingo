package bingo

import (
	"testing"
)

func Test_ImageRequest(t *testing.T) {

	req := ImageRequest{"Eye Bleach", "Off", ACCOUNT_KEY}

	results, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	if len(results) == 0 {
		t.Error("No Results")
	}
}
