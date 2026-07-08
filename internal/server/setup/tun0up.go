package serversetup

import (
	"log"
	"os/exec"
)

func Tun0Upstatus() error {

	status, uperr := exec.Command("ip", "link", "set", "guard-server", "up").CombinedOutput()


	if uperr != nil {
		log.Println("Interface Up Error",string(status))
		return uperr
	}

	return nil
}
