package setup

import "net"

//Function to check the tun0 interface exists or not 

func Tun0exists() error {
 
	 // Ask the linux kernal for the list of interfaces on go check the tun0 interface

	_,interfacerr := net.InterfaceByName("guard0")

	   if interfacerr != nil {
            createrr :=  Createtun0()
			
			if createrr != nil {
                 return createrr
			}
	   } 

       return nil 
}