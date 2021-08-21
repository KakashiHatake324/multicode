package main

import (
	"strconv"
	"time"
)

func GetTimestamp() string {
	now := time.Now()
	nanos := now.UnixNano()

	millis := nanos / 1000000000
	getStamp := strconv.FormatInt(millis, 10)
	return getStamp
}
