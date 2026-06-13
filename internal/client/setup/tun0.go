package setup

import (
	"os/exec"
)

//It used to create a tun interfaces for the virtual NIC Card

func Createtun0() error {

	  _,runerr := exec.Command("ip", "tuntap", "add" ,"dev","tun0" ,"mode" ,"tun").CombinedOutput()

       if runerr != nil {
		 Ipallocator()
		 return runerr
	   }

	  return nil 
}