// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUsersService_Get_self(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": "1"}}`)
	})

	user, err := client.Users.Get("")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &User{ID: "1"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Get_userId(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": "1"}}`)
	})

	user, err := client.Users.Get("1")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &User{ID: "1"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_MediaFeed(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/feed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"count":  "1",
			"min_id": "1",
			"max_id": "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		Count: 1,
		MinID: "1",
		MaxID: "1",
	}
	media, _, err := client.Users.MediaFeed(opt)
	if err != nil {
		t.Errorf("Users.MediaFeed returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Users.MediaFeed returned %+v, want %+v", media, want)
	}
}

func TestUsersService_RecentMedia(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/media/recent", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"count":         "1",
			"max_timestamp": "1",
			"min_timestamp": "1",
			"min_id":        "1",
			"max_id":        "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		Count:        1,
		MinTimestamp: 1,
		MaxTimestamp: 1,
		MinID:        "1",
		MaxID:        "1",
	}
	media, _, err := client.Users.RecentMedia("", opt)
	if err != nil {
		t.Errorf("Users.RecentMedia returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Users.RecentMedia returned %+v, want %+v", media, want)
	}
}

func TestUsersService_LikedMedia(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/media/liked", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"count":       "1",
			"max_like_id": "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		Count: 1,
		MaxID: "1",
	}
	media, _, err := client.Users.LikedMedia(opt)
	if err != nil {
		t.Errorf("Users.LikedMedia returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Users.LikedMedia returned %+v, want %+v", media, want)
	}
}

func TestUsersService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"count": "1",
			"q":     "u",
		})
		fmt.Fprint(w, `{"data": [{"username":"u"}]}`)
	})

	opt := &Parameters{
		Count: 1,
	}
	users, _, err := client.Users.Search("u", opt)
	if err != nil {
		t.Errorf("Users.Search returned error: %v", err)
	}

	want := []User{User{Username: "u"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.Search returned %+v, want %+v", users, want)
	}
}
