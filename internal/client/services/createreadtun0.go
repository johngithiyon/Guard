package services

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/setup"
	"github.com/songgao/water"
)

//Function is used to read the IP packet from the tun0

func CreateReadtun0() error {

	   //Creates a config object to tell the package I am using Tun 

	   config := water.Config{
		     DeviceType: water.TUN,
	   }

	   //Name the tun interface as guard0

	   config.Name = "guard0"

	   // open the tun interface

	   tun,createrr := water.New(config)

	   if createrr != nil {
		    return createrr
	   }

	   iperr := setup.ReqIpaddress()

	   if iperr != nil {
		   return iperr 
	   }
    
	   buffer := make([]byte,1024)

	   
	   for {
           
		    _,readerr :=  tun.Read(buffer)

			if readerr != nil {
				return readerr
			}

			log.Println(buffer) 
		      
	   }
}