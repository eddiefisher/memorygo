// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"crypto/md5"
	"encoding/hex"
)

//Messages ...
var Messages = map[string]string{}

func keyExist(key string) bool {
	_, ok := Messages[uid(key)]
	return ok
}

func uid(key string) (result string) {
	b := md5.Sum([]byte(key))
	result = hex.EncodeToString(b[:])
	return
}
