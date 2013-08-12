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

func TestLikesService_MediaLikes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/likes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, err := client.Likes.MediaLikes("1")
	if err != nil {
		t.Errorf("Likes.MediaLikes returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Likes.MediaLikes returned %+v, want %+v", users, want)
	}
}

func TestLikesService_Like(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/likes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"meta": {"code": 200}, "data": null}`)
	})

	err := client.Likes.Like("1")
	if err != nil {
		t.Errorf("Likes.Like returned error: %v", err)
	}
}

func TestLikesService_Unlike(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/likes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{"meta": {"code": 200}, "data": null}`)
	})

	err := client.Likes.Unlike("1")
	if err != nil {
		t.Errorf("Likes.Unlike returned error: %v", err)
	}
}
