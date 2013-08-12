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

func TestCommentsService_MediaComments(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/comments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":[{"id": "123"}]}`)
	})

	comments, err := client.Comments.MediaComments("1")
	if err != nil {
		t.Errorf("Comments.MediaComments returned error: %v", err)
	}

	want := []Comment{Comment{ID: "123"}}
	if !reflect.DeepEqual(comments, want) {
		t.Errorf("Comments.MediaComments returned %+v, want %+v", comments, want)
	}
}

func TestCommentsService_Add(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/comments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"text": "comment text",
		})
		fmt.Fprint(w, `{"meta":{"code":200},"data":null}`)
	})

	err := client.Comments.Add("1", []string{"comment text"})
	if err != nil {
		t.Errorf("Comments.Add returned error: %v", err)
	}
}

func TestCommentsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/media/1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{"meta":{"code":200},"data":null}`)
	})

	err := client.Comments.Delete("1", "1")
	if err != nil {
		t.Errorf("Comments.Delete returned error: %v", err)
	}
}
