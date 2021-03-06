package cloudpan

import (
	"github.com/tickstep/cloudpan189-go/library/requester"
	"net/http"
	"net/url"
)

const (
	// PathSeparator 路径分隔符
	PathSeparator = "/"
)

var (
	cloudpanDomainUrl = &url.URL{
		Scheme: "http",
		Host:   ".cloud.189.cn",
	}
)

type (
	PanClient struct {
		client     *requester.HTTPClient // http 客户端
		webToken WebLoginToken
		appToken AppLoginToken
	}
)


func NewPanClient(webToken WebLoginToken, appToken AppLoginToken) *PanClient {
	client := requester.NewHTTPClient()
	client.ResetCookiejar()
	client.Jar.SetCookies(cloudpanDomainUrl, []*http.Cookie{
		&http.Cookie{
			Name:   "COOKIE_LOGIN_USER",
			Value:  webToken.CookieLoginUser,
			Domain: "cloud.189.cn",
			Path: "/",
		},
	})

	return &PanClient{
		client: client,
		webToken: webToken,
		appToken: appToken,
	}
}

//func (p *PanClient) HttpClient() *requester.HTTPClient {
//	return p.client
//}