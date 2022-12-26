package main

import (
	"connectivity-alert/pingify"
	"time"
)

func main() {
	last := false
	for {
		println("Pinging at ", time.Now().Format(time.RFC822))
		current := pingify.TryPinging("www.google.com")
		if last != current {
			pingify.DisplayNotification(current)
			last = current
		}
		time.Sleep(5 * time.Minute)
	}
}
