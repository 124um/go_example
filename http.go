package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

// logging and next responce or request  - middleware
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func mainHttp() {

	// jar, err := cookiejar.New(nil)
	// jar.SetCookies()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("StatusCode ---->", req.Response.Status)
			fmt.Println("REDIRECT--------------------------------------")
			return nil
		},
		Transport: &loggingRoundTripper{  // this is middleware
			logger: os.Stdout,
			next: http.DefaultTransport,
		},
		Timeout: time.Second * 10,
		Jar: jar,
	}
	resp, err := client.Get("https://example.com/") // REDIRECT - and run CheckRedirect
	// resp, err := http.DefaultClient.Get("https://www.example.com/")
	// resp, err := http.DefaultClient.Post("https://www.example.com/")
	// resp, err := http.DefaultClient.PostForm("https://www.example.com/")
	// resp, err := http.DefaultClient.Do("https://www.example.com/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("StatusCode ---->", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if err != nil {
		log.Fatal(err)
	}
}
