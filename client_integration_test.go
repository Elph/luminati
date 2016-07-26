package luminati

import (
	"flag"
	"fmt"
	"net/http"
	"testing"
)

var user string
var pwd string
var host string
var port int

func init() {
	flag.StringVar(&user, "user", "", "luminati.io username")
	flag.StringVar(&pwd, "pwd", "", "luminati.io password")
	flag.StringVar(&host, "host", "", "luminati.io proxy")
	flag.IntVar(&port, "port", 22225, "luminati.io port")
	flag.Parse()
}

func TestClient_ShouldWorkAsExpected(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	if user == "" || pwd == "" || host == "" || port == 0 {
		t.Skip("You need to provide a valid luminati.io credentials to run this test (Skipping)")
	}

	urls := []string{
		"http://www.google.com",
		"http://twitter.com",
		"http://github.com",
	}

	for _, url := range urls {

		c := NewClient(user, pwd, host, port)
		request, _ := http.NewRequest("GET", url, nil)
		resp, err := c.Do(request)
		if err != nil {
			t.Fatal("The response for "+url+" returned an error ", err)
		}
		if resp.StatusCode != 200 {
			t.Fatal("Url: " + url + " returned a non valid status code " + string(resp.StatusCode))
		}
		fmt.Println("Url:" + url + " OK")

	}
}
