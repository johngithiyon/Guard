package setup

import (
	"log"
	"os/exec"
)


func Createtun0() error {

	  _,runerr := exec.Command("ip", "tuntap", "add" ,"dev","tun0" ,"mode" ,"tun").CombinedOutput()

       if runerr != nil {
		 log.Println("runerr",runerr)
		 return runerr
	   }
	   
	  return nil 
}