package setup

import (
	"log"
	"os/exec"
)

func Tun0Upstatus() error {

	   status,uperr := exec.Command("ip","link","set","guard0","up").CombinedOutput()

       log.Println("Tun0 status",status)

	   if uperr != nil {
		    return uperr
	   }

	   return nil 
}