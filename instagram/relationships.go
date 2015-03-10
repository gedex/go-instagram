// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
)

// RelationshipsService handles communication with the user's relationships related
// methods of the Instagram API.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/
type RelationshipsService struct {
	client *Client
}

// Relationship represents relationship authenticated user with another user.
type Relationship struct {
	// Current user's relationship to another user. Can be "follows", "requested", or "none".
	OutgoingStatus string `json:"outgoing_status,omitempty"`

	// A user's relationship to current user. Can be "followed_by", "requested_by",
	// "blocked_by_you", or "none".
	IncomingStatus string `json:"incoming_status,omitempty"`
}

// Follows gets the list of users this user follows. If empty string is
// passed then it refers to `self` or curret authenticated user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#get_users_follows
func (s *RelationshipsService) Follows(userId string) ([]User, *ResponsePagination, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v/follows", userId)
	} else {
		u = "users/self/follows"
	}

	users := new([]User)

RETRY:

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	tmp := new([]User)

	_, err = s.client.Do(req, tmp)
	if err != nil {
		return nil, nil, err
	}

	*users = append(*users, *tmp...)

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil && s.client.Response.Pagination.NextURL != "" {
		page = s.client.Response.Pagination
		u = page.NextURL
		goto RETRY
	}

	return *users, page, err
}

// FollowedBy gets the list of users this user is followed by. If empty string is
// passed then it refers to `self` or curret authenticated user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#get_users_followed_by
func (s *RelationshipsService) FollowedBy(userId string) ([]User, *ResponsePagination, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v/followed-by", userId)
	} else {
		u = "users/self/followed-by"
	}

	users := new([]User)

RETRY:

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	tmp := new([]User)

	_, err = s.client.Do(req, tmp)
	if err != nil {
		return nil, nil, err
	}

	*users = append(*users, *tmp...)

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil && s.client.Response.Pagination.NextURL != "" {
		page = s.client.Response.Pagination
		u = page.NextURL
		goto RETRY
	}

	return *users, page, err
}

// RequestedBy lists the users who have requested this user's permission to follow.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#get_incoming_requests
func (s *RelationshipsService) RequestedBy() ([]User, *ResponsePagination, error) {
	u := "users/self/requested-by"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	users := new([]User)

	_, err = s.client.Do(req, users)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *users, page, err
}

// Relationship gets information about a relationship to another user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#get_relationship
func (s *RelationshipsService) Relationship(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "", "GET")
}

// Follow a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Follow(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "follow", "POST")
}

// Unfollow a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Unfollow(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "unfollow", "POST")
}

// Block a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Block(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "block", "POST")
}

// Unblock a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Unblock(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "unblock", "POST")
}

// Approve a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Approve(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "approve", "POST")
}

// Deny a user.
//
// Instagram API docs: http://instagram.com/developer/endpoints/relationships/#post_relationship
func (s *RelationshipsService) Deny(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "deny", "POST")
}

func relationshipAction(s *RelationshipsService, userId, action, method string) (*Relationship, error) {
	u := fmt.Sprintf("users/%v/relationship", userId)
	if action != "" {
		action = "action=" + action
	}
	req, err := s.client.NewRequest(method, u, action)

	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Insta-Forwarded-For", s.client.XInstaForwardedFor())

	rel := new(Relationship)
	_, err = s.client.Do(req, rel)
	return rel, err
}
