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

func TestRelationshipsService_Follows_self(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/follows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, _, err := client.Relationships.Follows("")
	if err != nil {
		t.Errorf("Relationships.Follows returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Relationships.Follows returned %+v, want %+v", users, want)
	}
}

func TestRelationshipsService_Follows_userId(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/follows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, _, err := client.Relationships.Follows("1")
	if err != nil {
		t.Errorf("Relationships.Follows returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Relationships.Follows returned %+v, want %+v", users, want)
	}
}

func TestRelationshipsService_FollowedBy_self(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/followed-by", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, _, err := client.Relationships.FollowedBy("")
	if err != nil {
		t.Errorf("Relationships.FollowedBy returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Relationships.FollowedBy returned %+v, want %+v", users, want)
	}
}

func TestRelationshipsService_FollowedBy_userId(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/followed-by", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, _, err := client.Relationships.FollowedBy("1")
	if err != nil {
		t.Errorf("Relationships.FollowedBy returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Relationships.FollowedBy returned %+v, want %+v", users, want)
	}
}

func TestRelationshipsService_RequestedBy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/self/requested-by", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{"id":"1"}]}`)
	})

	users, _, err := client.Relationships.RequestedBy()
	if err != nil {
		t.Errorf("Relationships.RequestedBy returned error: %v", err)
	}

	want := []User{User{ID: "1"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Relationships.RequestedBy returned %+v, want %+v", users, want)
	}
}

func TestRelationshipsService_Relationship(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/relationship", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": {"outgoing_status":"none"}}`)
	})

	rel, err := client.Relationships.Relationship("1")
	if err != nil {
		t.Errorf("Relationships.Relationship returned error: %v", err)
	}

	want := &Relationship{OutgoingStatus: "none"}
	if !reflect.DeepEqual(rel, want) {
		t.Errorf("Relationships.Relationship returned %+v, want %+v", rel, want)
	}
}

func TestRelationshipsService_actions(t *testing.T) {
	setup()
	defer teardown()

	actions := make(map[string]func(userId string) (*Relationship, error))
	actions["Follow"] = client.Relationships.Follow
	actions["Unfollow"] = client.Relationships.Unfollow
	actions["Block"] = client.Relationships.Block
	actions["Unblock"] = client.Relationships.Unblock
	actions["Approve"] = client.Relationships.Approve
	actions["Deny"] = client.Relationships.Deny

	id := 1
	for a, f := range actions {
		urlStr := fmt.Sprintf("/users/%d/relationship", id)
		mux.HandleFunc(urlStr, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{"data": {"outgoing_status":"requested"}}`)
		})

		rel, err := f("1")
		if err != nil {
			t.Errorf("Relationships.%s returned error: %v", a, err)
		}

		want := &Relationship{OutgoingStatus: "requested"}
		if !reflect.DeepEqual(rel, want) {
			t.Errorf("Relationships.%s returned %+v, want %+v", a, rel, want)
		}
		id++
	}
}
