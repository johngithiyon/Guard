package setup

import (
	"log"
	"os/exec"
)

//Function Used to create a entry in the kernal for request route to the tun0 interface

func Route(ip string) error {

	status,routerr := exec.Command("ip","route","add",ip,"dev","guard0").CombinedOutput()

	if routerr != nil {
		log.Println("Status of route err",string(status))
		return routerr
	}

	return nil 
}