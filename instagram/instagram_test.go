// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Instagram client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Instagram client configured to use test server
	client = NewClient(nil)
	client.BaseURL, _ = url.Parse(server.URL)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	for key, want := range values {
		if v := r.FormValue(key); v != want {
			t.Errorf("Request parameter %v = %v, want %v", key, v, want)
		}
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	want := "https://api.instagram.com/v1/"
	if c.BaseURL.String() != want {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL.String(), want)
	}
	want = "github.com/gedex/go-instagram v0.5"
	if c.UserAgent != want {
		t.Errorf("NewClient UserAgent = %v, want %v", c.UserAgent, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	// set access_token
	c.AccessToken = "token"

	inURL, outURL := "foo", c.BaseURL.String()+"foo?access_token="+c.AccessToken
	req, _ := c.NewRequest("GET", inURL, "")

	// test that relative URL was expanded and access token appears in query string
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that default user-agent is attached to the requet
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, want %v", userAgent, c.UserAgent)
	}
}
