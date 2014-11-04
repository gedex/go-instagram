go-instagram
============

go-instagram is Go library for accessing Instagram REST and Search APIs.

**Documentation:** <http://godoc.org/github.com/gedex/go-instagram/instagram>

**Build Status:** [![Build Status](https://travis-ci.org/gedex/go-instagram.png?branch=master)](https://travis-ci.org/gedex/go-instagram)
[![Build Status](https://drone.io/github.com/gedex/go-instagram/status.png)](https://drone.io/github.com/gedex/go-instagram/latest)
[![Coverage Status](https://coveralls.io/repos/gedex/go-instagram/badge.png?branch=master)](https://coveralls.io/r/gedex/go-instagram?branch=master)

## Basic Usage

Access different parts of the Instagram API using the various services on a Instagram
Client:

~~~go
// You can optionally pass your own HTTP's client, otherwise pass it with nil.
client := instagram.NewClient(nil)
~~~

You can then optionally set ClientID, ClientSecret and AccessToken:

~~~go
client.ClientID = "8f2c0ad697ea4094beb2b1753b7cde9c"
~~~

With client object set, you can communicate with Instagram endpoints:

~~~go
// Gets the most recent media published by a user with id "3"
media, next, err := client.Users.RecentMedia("3", nil)
~~~

Set optional parameters for an API method by passing an Parameters object.

~~~go
// Gets user's feed.
opt := &instagram.Parameters{Count: 3}
media, next, err := client.Users.RecentMedia("3", opt)
~~~

Please see [examples/example.go](./examples/example.go) for a complete example.

## Data Retrieval

The methods which return slice in first return value will return three values (data, pagination, and error).
Here's an example of retrieving popular media:

~~~
media, next, err := client.Media.Popular()
if err != nil {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}
for _, m := range media {
	fmt.Printf("ID: %v, Type: %v\n", m.ID, m.Type)
}
if next.NextURL != "" {
	fmt.Println("Next URL", next.NextURL)
}
~~~

If a single type is returned in first return value, then only two values returned. Here's an example
of retrieving user's information:

~~~
user, err := client.Users.Get("3")
if err != nil {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}
fmt.Println("Username", user.Username)
~~~

## Credits

* [go-github](https://github.com/google/go-github) in which this library mimics the structure.
  LICENSE for go-github is included in [go-github-LICENSE.md](./go-github-LICENSE.md)
* [python-instagram](https://github.com/Instagram/python-instagram)
* [Instagram endpoints docs](http://instagram.com/developer/endpoints/)

## License

This library is distributed under the BSD-style license found in the LICENSE file.
