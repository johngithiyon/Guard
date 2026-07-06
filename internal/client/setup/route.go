package setup

import (
	"log"
	"os/exec"
)

//Function Used to create a entry in the kernal for request route to the tun0 interface

//Here the ip you have to put your ssh server private ip

func Route() error {

	status,routerr := exec.Command("ip","route","add","10.60.131.96","dev","guard0").CombinedOutput()

	if routerr != nil {
		log.Println("Status of route err",string(status))
		return routerr
	}

	return nil 
}