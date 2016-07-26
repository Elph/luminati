package luminati

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLuminatiClient_ShoulReturnContent_When200(t *testing.T) {
	const DATA string = "Hello World"
	c := NewClient("username", "password", "host.proxy.com", 123456)

	server := testTools(200, DATA)
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)
	response, err := c.Do(request)

	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "200 OK" {
		t.Fatal("Response is not OK " + response.Status)
	}

	content, _ := ioutil.ReadAll(response.Body)
	if string(content) != DATA {
		t.Fatal("Response data is not correct")
	}

}

func TestNewSession_ShouldGenerateNewSessionID(t *testing.T) {
	c := NewClient("username", "password", "host", 123456)
	s := c.SessionID

	c.NewSession()
	if s == c.SessionID {
		t.Fatal("SessionID has not changed")
	}

}

func TestSetProxy_ShouldAddTransport(t *testing.T) {
	server := testTools(200, "body")
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)

	client := &http.Client{}
	c := NewClient("username", "password", "host", 123456)
	c.setProxy(client, request)

	if client.Transport == nil {
		t.Fatal("Transport has not bee setted")
	}

}

func TestSetProxy_ShouldAddProxyAuthorizationHeader(t *testing.T) {
	server := testTools(200, "body")
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)

	client := &http.Client{}
	c := NewClient("username", "password", "host", 123456)
	c.setProxy(client, request)

	currentHeader := request.Header.Get("Proxy-Authorization")

	login := c.username + "-session-" + c.SessionID
	auth := login + ":" + c.password
	expectedHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	if currentHeader != expectedHeader {
		t.Fatal("The Proxy-Authorization header is incorrect")
	}

}

func testTools(code int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		//w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
	}))
	return server
}
