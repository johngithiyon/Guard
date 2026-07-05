package services

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/setup"
	"github.com/songgao/water"
)

//Function is used to read the IP packet from the tun0

func Createtun0() error {

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

	   //Goroutine for send the ip request to the server

	   go setup.ReqIpaddress(Conn)

	   //Goroutine for read the packets from the tun0 interface
    
	   go  Readtun0(tun)

	   //Goroutine for read the packets from the socket 

	   go Readresp(tun)

	   return nil 
}