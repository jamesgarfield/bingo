package bingo

import (
	"fmt"
	"testing"
)

func Test_RequestUrl(t *testing.T) {

	req := ImageRequest{"Gurren Lagann", "Off", ACCOUNT_KEY}

	url, err := RequestUrl(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(url)

}
