package pingify

import (
	"github.com/gen2brain/beeep"
	probing "github.com/prometheus-community/pro-bing"
)

// returns true if connected, false if not
func DisplayNotification(connectStatus bool) {
	if connectStatus == false {
		e := beeep.Notify("ConnectivityAlert", "Disconnected","")
		if e != nil {
			panic(e)
		}
	} else {
		e := beeep.Notify("ConnectivityAlert", "Connected","")
		if e != nil {
			panic(e)
		}
	}
}

// returns true if connected, false if not
func TryPinging(pingAddress string) bool {
	pinger, err := probing.NewPinger(pingAddress)
	if err != nil {
		// fail silently for now, handle without panic
		// panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.

	if err != nil {
		return false
	} else {
		return true
	}
}
