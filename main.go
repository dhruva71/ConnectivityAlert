package main

import (
	"connectivity-alert/pingify"
	"time"
)

func main() {
	last := false
	for {
		current := pingify.TryPinging("www.google.com")
		if last != current {
			pingify.DisplayNotification(current)
			last = current
		}
		time.Sleep(1 * time.Minute)
	}
}
