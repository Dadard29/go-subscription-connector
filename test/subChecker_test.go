package test

import (
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"testing"
)

func TestCheckToken(t *testing.T) {
	s := subChecker.NewSubChecker("localhost:8080")
	token := "262403d5e1330d63693919f10584e353b5af9e5762acc7761d0f16d0e240b350"
	apiName := "youtube-download"
	_, err := s.CheckToken(token, apiName)
	if err != nil {
		t.Error(err)
	}
}
