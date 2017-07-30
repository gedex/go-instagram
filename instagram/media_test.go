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

func TestMediaService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": "1"}}`)
	})

	media, err := client.Media.Get("1")
	if err != nil {
		t.Errorf("Media.Get returned error: %v", err)
	}

	want := &Media{ID: "1"}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Media.Get returned %+v, want %+v", media, want)
	}
}

func TestMediaService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"lat":           "37.7749295",
			"lng":           "-122.4194155",
			"max_timestamp": "1",
			"min_timestamp": "1",
			"distance":      "1000.0000000",
			"count":         "100",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		Lat:          37.7749295,
		Lng:          -122.4194155,
		MinTimestamp: 1,
		MaxTimestamp: 1,
		Distance:     1000,
		Count:        100,
	}
	media, _, err := client.Media.Search(opt)
	if err != nil {
		t.Errorf("Media.Search returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Media.Search returned %+v, want %+v", media, want)
	}
}

func TestMediaService_Popular(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/popular", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	media, _, err := client.Media.Popular()
	if err != nil {
		t.Errorf("Media.Popular returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Media.Popular returned %+v, want %+v", media, want)
	}
}
