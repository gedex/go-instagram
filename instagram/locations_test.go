// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

func TestLocationsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/locations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": "1"}}`)
	})

	loc, err := client.Locations.Get("1")
	if err != nil {
		t.Errorf("Locations.Get returned error: %v", err)
	}

	want := &Location{ID: "1"}
	if !reflect.DeepEqual(loc, want) {
		t.Errorf("Locations.Get returned %+v, want %+v", loc, want)
	}
}

func TestLocationsService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/locations/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"lat":      "37.7749295",
			"lng":      "-122.4194155",
			"distance": strconv.FormatFloat(5000, 'f', 7, 64),
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		Distance: 5000,
	}
	locations, err := client.Locations.Search(37.7749295, -122.4194155, opt)
	if err != nil {
		t.Errorf("Locations.Search returned error: %v", err)
	}

	want := []Location{Location{ID: "1"}}
	if !reflect.DeepEqual(locations, want) {
		t.Errorf("Locations.Search returned %+v, want %+v", locations, want)
	}
}

func TestLocationsService_RecentMedia(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/locations/1/media/recent", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"max_timestamp": "1",
			"min_timestamp": "1",
			"max_id":        "1",
			"min_id":        "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		MinTimestamp: 1,
		MaxTimestamp: 1,
		MinID:        "1",
		MaxID:        "1",
	}
	media, _, err := client.Locations.RecentMedia("1", opt)
	if err != nil {
		t.Errorf("Location.RecentMedia returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Location.RecentMedia returned %+v, want %+v", media, want)
	}
}
