package setup

import (
	"log"
	"os/exec"
)

//It used to create a tun interfaces for the virtual NIC Card

func Createtun0() error {

	  runinfo,runerr := exec.Command("ip", "tuntap", "add" ,"dev","tun0" ,"mode" ,"tun").CombinedOutput()

	  log.Println(string(runinfo))

       if runerr != nil {
		log.Println("Tun0 Creation Error",runerr)
		return runerr
	   }

	  allocaterr := ReqIpaddress()

	  if allocaterr != nil {
		  return allocaterr
	  }

	  return nil 
}