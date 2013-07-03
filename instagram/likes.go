// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
)

// LikesService handles communication with the likes related
// methods of the Instagram API.
//
// Instagram API docs: http://instagram.com/developer/endpoints/likes/
type LikesService struct {
	client *Client
}

// MediaLikes gets a list of users who have liked mediaId.
//
// Instagram API docs: http://instagram.com/developer/endpoints/likes/#get_media_likes
func (s *LikesService) MediaLikes(mediaId string) ([]User, error) {
	u := fmt.Sprintf("media/%v/likes", mediaId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	users := new([]User)
	_, err = s.client.Do(req, users)
	return *users, err
}

// Like a media.
//
// Instagram API docs: http://instagram.com/developer/endpoints/likes/#post_likes
func (s *LikesService) Like(mediaId string) error {
	return mediaLikesAction(s, mediaId, "POST")
}

// Unlike a media.
//
// Instagram API docs: http://instagram.com/developer/endpoints/likes/#delete_likes
func (s *LikesService) Unlike(mediaId string) error {
	return mediaLikesAction(s, mediaId, "DELETE")
}

func mediaLikesAction(s *LikesService, mediaId, method string) error {
	u := fmt.Sprintf("media/%v/likes", mediaId)
	req, err := s.client.NewRequest(method, u, "")
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
