// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transport

import (
	"fmt"
	"net"
	"os"
)

func Start(network, address string, handle func(conn net.Conn)) {
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			os.Exit(1)
		}
		go handle(conn)
	}
}
