// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"errors"
	"strings"
)

//Set means "store this data"
func Set(key, value string) {
	Messages[uid(key)] = value
}

//Add means "store this data, but only if the server *doesn't* already hold data for this key".
func Add(key, value string) error {
	if !keyExist(key) {
		Messages[uid(key)] = value
		return nil
	}
	return errors.New("key exist")
}

//Replace means "store this data, but only if the server *does* already hold data for this key".
func Replace(key, value string) error {
	if keyExist(key) {
		Messages[uid(key)] = value
		return nil
	}
	return errors.New("key not exist")
}

//Append means "add this data to an existing key after existing data".
// TODO The append command do not accept flags or exptime.
// TODO They update existing data portions, and ignore new flag and exptime settings.
func Append(key, value string) error {
	if keyExist(key) {
		Messages[uid(key)] = strings.Join([]string{Messages[uid(key)], value}, "")
		return nil
	}
	return errors.New("key not exist")
}

//Prepend means "add this data to an existing key before existing data".
// TODO The prepend command do not accept flags or exptime.
// TODO They update existing data portions, and ignore new flag and exptime settings.
func Prepend(key, value string) error {
	if keyExist(key) {
		Messages[uid(key)] = strings.Join([]string{value, Messages[uid(key)]}, "")
		return nil
	}
	return errors.New("key not exist")
}
