// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Sun,  7 Jan 2018]

package main

import (
	"fmt"
	"net"

	"github.com/eddiefisher/memorygo/src/store"
	"github.com/eddiefisher/memorygo/src/transport"
)

var (
	network = "tcp"
	address = "127.0.0.1:33600"
)

func main() {
	go transport.Start(network, address, Handle)
	store.Set("1", "2")
	store.Set("2", "3")
	fmt.Println(store.Messages)
	store.Delete("1")
	fmt.Println(store.Messages)
}

func Handle(conn net.Conn) {}
