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
		    log.Println("Create err",createrr)
		    return createrr
	   }

	   iperr := setup.ReqIpaddress(Conn)

	   if iperr != nil {
		   log.Println("iperr",iperr)
		   return iperr 
	   }
    
	   buffer := make([]byte,1024)

	   
	   for {
           
		    length,readerr :=  tun.Read(buffer)

			if readerr != nil {
				return readerr
			}

			if length > 0 {
				 log.Println("Recieved packets in client",length,buffer[:length])
                 Sendpackets(tun,buffer[:length])
			}

			
		      
	   }
}