package pkg

import (
	"encoding/json"
	"fmt"
	"gin-learn/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type IOauth interface {
	GetAuthURL()     			string
	GetAccessToken(code string) string
	GetUserInfo()    			string
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
	result, err := post(url, nil, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get oauth access token failed, err:%v", err)
		return ""
	}
	return result["access_token"].(string)
}

func (o *Oauth)GetUserInfo() string {
	return ""
}

func (gho *GitHubOauth)GetAccessToken(code string) string  {
	url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", gho.AccessTokenUrl, gho.ClientID, gho.ClientSecrect, code)
	headers := make(map[string]string)
	headers["accept"] = "application/json"
	result, err := post(url, nil, headers)
	if err != nil {
		log.Fatalf("Get oauth access token failed, err:%v", err)
		return ""
	}
	return result["access_token"].(string)
}

func (bdo *BaiduOauth)GetAuthURL() string  {
	return fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&scope=basic&display=popup", bdo.AuthorizeUrl, bdo.ClientID, bdo.RedirectUrl)
}

func (bdo *BaiduOauth)GetAccessToken(code string) string  {
	url := fmt.Sprintf("%s?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s",
		bdo.AccessTokenUrl, code, bdo.ClientID, bdo.ClientSecrect, bdo.RedirectUrl)
	result, err := post(url, nil, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Get oauth access token failed, err:%v", err)
		return ""
	}
	return result["access_token"].(string)
}

func post(url string, body io.Reader, headers map[string]string) (map[string]interface{}, error) {
	client := http.Client{}
	request, _ := http.NewRequest("POST", url, body)

	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	var result map[string]interface{}
	var data []byte
	data, _ = ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(data, &result)
	return result, nil
}

var GHO IOauth
var BDO IOauth

func SetUp()  {
	gho := NewGitHub()
	GHO = &gho

	bdo := NewBaidu()
	BDO = &bdo
}

