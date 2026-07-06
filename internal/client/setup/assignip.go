package setup

import (
	"log"
	"os/exec"
)

//Function to assign the ip to the tun0 interface

func Assignip(ip string) error {

	 ipassign,ipassignerr := exec.Command("ip","addr","add",ip,"dev","guard0").CombinedOutput()

	 log.Println(string(ipassign))

	 if ipassignerr != nil {
		 log.Println("Ip assign error",ipassignerr)
		 return ipassignerr
	 }

	 uperr := Tun0Upstatus()

	 if uperr != nil {
		    return uperr
	 }

	 routerr := Route()

	 if routerr != nil {
		return routerr
	 }

	 return nil
}