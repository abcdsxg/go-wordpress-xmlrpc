package xmlrpc

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abcdsxg/go-wordpress-xmlrpc/wordpress"
	"github.com/kolo/xmlrpc"
)

type HttpRT struct {
	http.RoundTripper
}

// NewHttpRT return a http.RoundTripper can print the request body
func NewHttpRT(t http.RoundTripper) http.RoundTripper {
	return &HttpRT{t}
}

// RoundTrip implement a http.RoundTripper can print the request body
func (t HttpRT) RoundTrip(req *http.Request) (*http.Response, error) {
	// you can customize  to get more control over connection options
	// example: print the request body

	b, err := req.GetBody()
	if err != nil {
		log.Println(err)
	}
	r, err := ioutil.ReadAll(b)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(r))
	return t.RoundTripper.RoundTrip(req)
}

// Client Packaging the xmlrpc client
type Client struct {
	*xmlrpc.Client
	UserInfo
}

// UserInfo wordpress's username and password
type UserInfo struct {
	Username string
	Password string
}

// NewDefaultClient default implement a http.RoundTripper can print the request body
func NewDefaultClient(url string, info UserInfo) (*Client, error) {
	t := NewHttpRT(http.DefaultTransport)
	c, err := xmlrpc.NewClient(url, t)
	return &Client{Client: c, UserInfo: info}, err
}

// NewClient without  http.RoundTripper
func NewClient(url string, info UserInfo) (*Client, error) {
	c, err := xmlrpc.NewClient(url, nil)
	return &Client{Client: c, UserInfo: info}, err
}

// NewCustomizeClient you can Customize your  http.RoundTripper
func NewCustomizeClient(url string, t http.RoundTripper, info UserInfo) (*Client, error) {
	c, err := xmlrpc.NewClient(url, t)
	return &Client{Client: c, UserInfo: info}, err
}

// Call abstract to proxy xmlrpc call
func (c *Client) Call(baseCall wordpress.BaseCall) (result interface{}, err error) {
	err = c.Client.Call(baseCall.GetMethord(), baseCall.GetArgs(c.Username, c.Password), &result)
	return result, err
}
