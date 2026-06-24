package serversetup


import (
	"net"
)

//Function to check the tun0 interface exists or not

func Tun0exists() error {
 
	 // Ask the linux kernal for the list of interfaces on go check the tun0 interface

	_,interfacerr := net.InterfaceByName("guard-server")

	   if interfacerr != nil {
            return interfacerr
	   } 

	   uperr := Tun0Upstatus()

	   if uperr != nil {
		    return uperr
	   }

       return nil 
}