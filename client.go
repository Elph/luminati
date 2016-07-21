package luminati

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

// Client represents a luminati.io client that setups internally the proxy
type Client struct {
	username  string
	password  string
	host      string
	port      int
	SessionID string
}

// NewClient Creates a Client with the luminati.io credentials
func NewClient(username string, password string, host string, port int) Client {
	c := Client{
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
	c.NewSession()
	return c
}

// NewSession generates a new SessionId to be used with luminati network
func (c *Client) NewSession() {
	s := randomString(10)
	c.SessionID = s
}

// Do executes a reqeuquest using a proxy from luminaati.io network
func (c *Client) Do(request *http.Request) (resp *http.Response, err error) {
	client := &http.Client{}
	c.setProxy(client, request)
	response, err := client.Do(request)
	return response, err
}

func (c *Client) setProxy(client *http.Client, request *http.Request) {
	proxyURL, _ := url.Parse("http://" + c.host + ":" + strconv.Itoa(c.port) + " " + request.URL.String())
	client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	login := c.username + "-session-" + c.SessionID
	auth := login + ":" + c.password
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Proxy-Authorization", basic)
}

func randomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
