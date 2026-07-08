package setup

import (
	"log"
	"os/exec"
)

func Tun0Upstatus() error {

	status, uperr := exec.Command("ip", "link", "set", "guard0", "up").CombinedOutput()

	if uperr != nil {
		log.Println("Interface Up Err",string(status))
		return uperr
	}

	return nil
}