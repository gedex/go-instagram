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

func TestGeographiesService_RecentMedia(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/geographies/1/media/recent", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"min_id": "1",
			"count":  "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		MinID: "1",
		Count: 1,
	}
	media, _, err := client.Geographies.RecentMedia("1", opt)
	if err != nil {
		t.Errorf("Geographies.RecentMedia returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Geographies.RecentMedia returned %+v, want %+v", media, want)
	}
}
