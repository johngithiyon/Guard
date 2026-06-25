package setup

import (
	"log"
	"os"
	"os/exec"
)

//Function Used to create a entry in the kernal for request route to the tun0 interface

func Route() error {

	//Replace the ip with your ssh server private ip
	
	sship := os.Getenv("SSH_IP")

	status,routerr := exec.Command("ip","route","add",sship,"dev","guard0").CombinedOutput()

	if routerr != nil {
		log.Println("Status of route err",status)
		return routerr
	}

	return nil 
}