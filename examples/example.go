// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

func main() {

	client := instagram.NewClient(nil)
	client.ClientID = "8f2c0ad697ea4094beb2b1753b7cde9c"
	// If you have access_token you can supply it
	// client.AccessToken = ""

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
}
