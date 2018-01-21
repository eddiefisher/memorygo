// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import "errors"

//Delete allows for explicit deletion of items
func Delete(key string) error {
	if keyExist(key) {
		delete(Messages, uid(key))
		return nil
	}
	return errors.New("key not exist")
}
