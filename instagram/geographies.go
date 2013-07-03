// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
)

// GeographiesService handles communication with the geographies related
// methods of the Instagram API.
//
// Instagram API docs: http://instagram.com/developer/endpoints/geographies/
type GeographiesService struct {
	client *Client
}

// RecentMedia gets recent media from a geography subscription that created by
// real-time subscriptions.
//
// Instagram API docs: http://instagram.com/developer/endpoints/geographies/#get_geographies_media_recent
func (s *GeographiesService) RecentMedia(geoId string) ([]Media, error) {
	u := fmt.Sprintf("geographies/%v/media/recent", geoId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	media := new([]Media)
	_, err = s.client.Do(req, media)
	return *media, err
}
