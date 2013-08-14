// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
	"net/url"
)

// TagsService handles communication with the tag related
// methods of the Instagram API.
//
// Instagram API docs: http://instagram.com/developer/endpoints/tags/
type TagsService struct {
	client *Client
}

// Tag represents information about a tag object.
type Tag struct {
	MediaCount int    `json:"media_count,omitempty"`
	Name       string `json:"name,omitempty"`
}

// Get information aout a tag object.
//
// Instagram API docs: http://instagram.com/developer/endpoints/tags/#get_tags
func (s *TagsService) Get(tagName string) (*Tag, error) {
	u := fmt.Sprintf("tags/%v", tagName)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	tag := new(Tag)
	_, err = s.client.Do(req, tag)
	return tag, err
}

// RecentMedia Get a list of recently tagged media.
//
// Instagram API docs: http://instagram.com/developer/endpoints/tags/#get_tags_media_recent
func (s *TagsService) RecentMedia(tagName string, opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := fmt.Sprintf("tags/%v/media/recent", tagName)
	if opt != nil {
		params := url.Values{}
		if opt.MinID != "" {
			params.Add("min_id", opt.MinID)
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		u += "?" + params.Encode()
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	media := new([]Media)

	_, err = s.client.Do(req, media)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *media, page, err
}

// Search for tags by name.
//
// Instagram API docs: http://instagram.com/developer/endpoints/tags/#get_tags_search
func (s *TagsService) Search(q string) ([]Tag, *ResponsePagination, error) {
	u := "tags/search?q=" + q
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	tags := new([]Tag)

	_, err = s.client.Do(req, tags)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *tags, page, err
}
