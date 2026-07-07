package serversetup

import (
	"log"
	"os/exec"
)

//Function to assign the ip to the tun0 interface

// I am just hardcode the ip address you can config any private ip range but it same to the client interface range

func Assignip() error {

	ipassign,ipassignerr := exec.Command("ip","addr","add","10.0.0.1/24","dev","guard-server").CombinedOutput()

	log.Println(string(ipassign))

	if ipassignerr != nil {
		log.Println("Ip assign error",ipassignerr)
		return ipassignerr
	}

	return nil 
}	
