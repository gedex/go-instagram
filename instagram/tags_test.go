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

func TestTagsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/tag-name", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"name": "tag-name"}}`)
	})

	tag, err := client.Tags.Get("tag-name")
	if err != nil {
		t.Errorf("Tags.Get returned error: %v", err)
	}

	want := &Tag{Name: "tag-name"}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tag.Get returned %+v, want %+v", tag, want)
	}
}

func TestTagsService_RecentMedia(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/tag-name/media/recent", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"min_id": "1",
			"max_id": "1",
		})
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	opt := &Parameters{
		MinID: "1",
		MaxID: "1",
	}
	media, _, err := client.Tags.RecentMedia("tag-name", opt)
	if err != nil {
		t.Errorf("Tags.RecentMedia returned error: %v", err)
	}

	want := []Media{Media{ID: "1"}}
	if !reflect.DeepEqual(media, want) {
		t.Errorf("Tags.RecentMedia returned %+v, want %+v", media, want)
	}
}

func TestTagsService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"q": "tag-name",
		})
		fmt.Fprint(w, `{"data": [{"name":"tag-name"}]}`)
	})

	tags, _, err := client.Tags.Search("tag-name")
	if err != nil {
		t.Errorf("Tags.Search returned error: %v", err)
	}

	want := []Tag{Tag{Name: "tag-name"}}
	if !reflect.DeepEqual(tags, want) {
		t.Errorf("Tags.Search returned %+v, want %+v", tags, want)
	}
}
