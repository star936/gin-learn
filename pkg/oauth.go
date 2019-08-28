package pkg

import (
	"fmt"
	"gin-learn/config"
	"gin-learn/pkg/utils"
	"log"
)

type IOauth interface {
	GetAuthURL()     				string
	GetAccessToken(code string) 	string
	GetUserInfo(accessToken string) string
}


type Oauth struct {
	AuthorizeUrl   string
	ClientID       string
	ClientSecrect  string
	AccessTokenUrl string
	RedirectUrl    string
	UserInfoUrl    string
}

type GitHubOauth struct {
	Oauth
}

type BaiduOauth struct {
	Oauth
}

func New(name string) Oauth {
	var o Oauth
	if err := config.Config.UnmarshalKey("oauth." + name, &o); err != nil {
		log.Fatalf("Unmarshal oauth failed, err:%s", err)
	}
	o.RedirectUrl = fmt.Sprintf("%s/%s/oauth/redirect", config.Config.GetString("baseUrl"), name)
	return o
}

func NewGitHub() GitHubOauth  {
	gho := GitHubOauth{New("github")}
	return gho
}

func NewBaidu() BaiduOauth  {
	bdo := BaiduOauth{New("baidu")}
	return bdo
}


func (o *Oauth)GetAuthURL() string {
	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", o.AuthorizeUrl, o.ClientID, o.RedirectUrl);
}

func (o *Oauth)GetAccessToken(code string) string {
	url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", o.AccessTokenUrl, o.ClientID, o.ClientSecrect, code)
	result, err := utils.Post(url, nil, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get oauth access token failed, err:%v", err)
		return ""
	}
	if accessToken, ok := result["access_token"]; ok {
		return accessToken.(string)
	}
	return ""
}

func (o *Oauth)GetUserInfo(accessToken string) string {
	return ""
}

func (gho *GitHubOauth)GetAccessToken(code string) string  {
	url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", gho.AccessTokenUrl, gho.ClientID, gho.ClientSecrect, code)
	headers := make(map[string]string)
	headers["accept"] = "application/json"
	result, err := utils.Post(url, nil, headers)
	if err != nil {
		log.Fatalf("Get access token from github failed, err:%v", err)
		return ""
	}
	if accessToken, ok := result["access_token"]; ok {
		return accessToken.(string)
	}
	return ""
}

func (gho *GitHubOauth)GetUserInfo(accessToken string) string {
	var headers = make(map[string]string)
	headers["accept"] = "application/json"
	headers["Authorization"] = "token " + accessToken
	result, err := utils.Get(gho.UserInfoUrl, headers)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get userinfo from github failed, err:%v", err)
		return ""
	}
	if name, ok := result["login"]; ok {
		return name.(string)
	}
	return ""
}

func (bdo *BaiduOauth)GetAuthURL() string  {
	return fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&scope=basic&display=popup", bdo.AuthorizeUrl, bdo.ClientID, bdo.RedirectUrl)
}

func (bdo *BaiduOauth)GetAccessToken(code string) string  {
	url := fmt.Sprintf("%s?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s",
		bdo.AccessTokenUrl, code, bdo.ClientID, bdo.ClientSecrect, bdo.RedirectUrl)
	result, err := utils.Post(url, nil, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get access token from baidu failed, err:%v", err)
		return ""
	}
	if accessToken, ok := result["access_token"]; ok {
		return accessToken.(string)
	}
	return ""
}

func (bdo *BaiduOauth)GetUserInfo(accessToken string) string  {
	url := bdo.UserInfoUrl + "?access_token=" + accessToken
	result , err :=utils.Post(url, nil, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get userinfo from baidu failed, err:%v", err)
	}
	if name, ok := result["uname"]; ok {
		return name.(string)
	}
	return ""
}


var GHO IOauth
var BDO IOauth

func SetUp()  {
	gho := NewGitHub()
	GHO = &gho

	bdo := NewBaidu()
	BDO = &bdo
}

