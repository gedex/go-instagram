// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
)

func main() {

	client := instagram.NewClient(nil)
	client.ClientID = "c54ce323e05a45409217344f8ee7e542"
	client.ClientSecret = "1c94683d110a459999a1268582d2fbc7"
	client.AccessToken = "1651306413.c54ce32.a05d9ce60d4a40059d8369192c848ce7"

	// If you have access_token you can supply it
	// client.AccessToken = ""

	//corrct access header 67.79.8.126|d86ae2b3cb7db26841c58763c3c276e5581d6f334bfd860ce217719ed1188da4

	fmt.Println("external ip =>", instagram.ExternalIP())
	fmt.Println(instagram.ComputeHmac256("67.79.8.126", client.ClientSecret))
	fmt.Println(instagram.ComputeHmac256(instagram.ExternalIP(), client.ClientSecret))

	_, err := client.Relationships.Unfollow("1420811656")

	if err != nil {
		panic(err)
	}

	// fmt.Println(r)

	// media, next, err := client.Media.Popular()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	// }
	// for _, m := range media {
	// 	fmt.Printf("ID: %v, Type: %v\n", m.ID, m.Type)
	// }
	// if next.NextURL != "" {
	// 	fmt.Println("Next URL", next.NextURL)
	// }

}
