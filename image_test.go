package bingo

import (
	"fmt"
	"testing"
)

func Test_ImageRequest(t *testing.T) {

	req := ImageRequest{"Eye Bleach", "Off", ACCOUNT_KEY}

	results, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for _, r := range results {
		fmt.Println(r.Json())
	}

}
