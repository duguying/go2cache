// Copyright 2018. All rights reserved.
// This file is part of go2cache project
// Created by duguying on 2018/5/17.

package go2cache

import (
	"testing"
	"time"
)

func TestBroadcast(t *testing.T) {

	go dial([]string{"127.0.0.1:9981", "192.168.2.227:60000"})

	go listen("127.0.0.1:9981")

	time.Sleep(time.Second * 3)
}
