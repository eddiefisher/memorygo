// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"errors"
	"strings"
)

//Get After this command, the client expects zero or more items, each of
// which is received as a text line followed by a data block. After all
// the items have been transmitted, the server sends the string
func Get(key string) (string, error) {
	if keyExist(key) {
		return Messages[uid(key)], nil
	}
	return "", errors.New("key not exist")
}

//Gets After this command, the client expects zero or more items, each of
// which is received as a text line followed by a data block. After all
// the items have been transmitted, the server sends the string
func Gets(keys ...string) ([]string, error) {
	var result, errs []string
	for _, key := range keys {
		if !keyExist(key) {
			errs = append(errs, key)
		}
	}
	if len(errs) > 0 {
		return []string{}, errors.New("keys is not exist: " + strings.Join(errs, " "))
	}
	for _, key := range keys {
		result = append(result, Messages[uid(key)])
	}
	return result, nil
}
