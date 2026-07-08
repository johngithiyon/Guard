package serversetup

import (
	"log"
	"os/exec"
)

func Mtulimit() error {

	   status,mtuerr := exec.Command("ip","link","set","dev","guard-server","mtu","1400").CombinedOutput()
	  
	   if mtuerr != nil {
		     log.Println("Mtu err",string(status),mtuerr)
			 return mtuerr
	   }

	   return nil 

}