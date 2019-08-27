package models

import (
	"gin-learn/config"
	"log"
)

type Oauth struct {
	AuthorizeUrl   string
	ClientID       string
	ClientSecrect  string
	AccessTokenUrl string
	UserInfoUrl    string
}

func New(name string) *Oauth {
	var o Oauth
	if err := config.Config.UnmarshalKey("oauth." + name, &o); err != nil {
		log.Fatalf("Unmarshal oauth failed, err:%s", err)
	}
	return &o
}
