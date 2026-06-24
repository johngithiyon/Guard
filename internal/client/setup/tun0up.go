package setup

import (
	"log"
	"os/exec"
)

func Tun0Upstatus() error {

	status, uperr := exec.Command("ip", "link", "set", "guard0", "up").CombinedOutput()

	log.Println("Tun0 status", status)

	if uperr != nil {
		log.Println("status error", uperr)
		return uperr
	}

	return nil
}