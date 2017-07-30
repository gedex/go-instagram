// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

// MediaService handles communication with the media related
// methods of the Instagram API.
//
// Instagram API docs: http://instagram.com/developer/endpoints/media/
type MediaService struct {
	client *Client
}

// Media represents a single media (image or video) on Instagram.
type Media struct {
	Type         string         `json:"type,omitempty"`
	UsersInPhoto []*UserInPhoto `json:"users_in_photo,omitempty"`
	Filter       string         `json:"filter,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Comments     *MediaComments `json:"comments,omitempty"`
	Caption      *MediaCaption  `json:"caption,omitempty"`
	Likes        *MediaLikes    `json:"likes,omitempty"`
	Link         string         `json:"link,omitempty"`
	User         *User          `json:"user,omitempty"`
	UserHasLiked bool           `json:"user_has_liked,omitempty"`
	CreatedTime  int64          `json:"created_time,string,omitempty"`
	Images       *MediaImages   `json:"images,omitempty"`
	Videos       *MediaVideos   `json:"videos,omitempty"`
	ID           string         `json:"id,omitempty"`
	Location     *MediaLocation `json:"location,omitempty"`
}

// MediaComments represents comments on Instagram's media.
type MediaComments struct {
	Count int        `json:"count,omitempty"`
	Data  []*Comment `json:"data,omitempty"`
}

// MediaLikes represents likes on Instagram's media.
type MediaLikes struct {
	Count int `json:"count,omitempty"`
	Data  []*User
}

// MediaCaption represents caption on Instagram's media.
type MediaCaption struct {
	CreatedTime int64  `json:"created_time,string,omitempty"`
	Text        string `json:"text,omitempty"`
	From        *User  `json:"from,omitempty"`
	ID          string `json:"id,omitempty"`
}

// UserInPhoto represents a single user, with its position, on Instagram photo.
type UserInPhoto struct {
	User     *User                `json:"user,omitempty"`
	Position *UserInPhotoPosition `json:"position,omitempty"`
}

// UserInPhotoPosition represents position of the user on Instagram photo.
type UserInPhotoPosition struct {
	x float64 `json:"x,omitempty"`
	y float64 `json:"y,omitempty"`
}

// MediaImages represents MediaImage with various resolutions.
type MediaImages struct {
	LowResolution      *MediaImage `json:"low_resolution,omitempty"`
	Thumbnail          *MediaImage `json:"thumbnail,omitempty"`
	StandardResolution *MediaImage `json:"standard_resolution,omitempty"`
}

// MediaImage represents Instagram media with type image.
type MediaImage struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// MediaVideos represents MediaVideo with various resolutions.
type MediaVideos struct {
	LowResolution      *MediaVideo `json:"low_resolution,omitempty"`
	StandardResolution *MediaVideo `json:"standard_resolution,omitempty"`
}

// MediaVideo represents Instagram media with type video.
type MediaVideo struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// MediaLocation represents information about a location.
//
// There's Location type under LocationsService, the different is
// the ID type. I've reported this inconsistency to Instagram
// https://groups.google.com/forum/#!topic/instagram-api-developers/Fty5lOsOGEg
type MediaLocation struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// Get information about a media object.
//
// Instagram API docs: http://instagram.com/developer/endpoints/media/#get_media
func (s *MediaService) Get(mediaId string) (*Media, error) {
	u := fmt.Sprintf("media/%v", mediaId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	media := new(Media)
	_, err = s.client.Do(req, media)
	return media, err
}

// Search return search results for media in a given area.
//
// http://instagram.com/developer/endpoints/media/#get_media_search
func (s *MediaService) Search(opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := "media/search"
	if opt != nil {
		params := url.Values{}
		if opt.Lat != 0 {
			params.Add("lat", strconv.FormatFloat(opt.Lat, 'f', 7, 64))
		}
		if opt.Lng != 0 {
			params.Add("lng", strconv.FormatFloat(opt.Lng, 'f', 7, 64))
		}
		if opt.MinTimestamp != 0 {
			params.Add("min_timestamp", strconv.FormatInt(opt.MinTimestamp, 10))
		}
		if opt.MaxTimestamp != 0 {
			params.Add("max_timestamp", strconv.FormatInt(opt.MaxTimestamp, 10))
		}
		if opt.Distance != 0 {
			params.Add("distance", strconv.FormatFloat(opt.Distance, 'f', 7, 64))
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
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

// Popular gets a list of what media is most popular at the moment.
//
// Instagram API docs: http://instagram.com/developer/endpoints/media/#get_media_popular
func (s *MediaService) Popular() ([]Media, *ResponsePagination, error) {
	u := "media/popular"
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
