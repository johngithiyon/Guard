package main

import (
	"log"
	"net"

	"github.com/johngithiyon/Guard/internal/server/services"
	serversetup "github.com/johngithiyon/Guard/internal/server/setup"
)

func main() {

        log.Println("Server Is Listening ...")     

	   // Creates a Udp socket listening on the port 8080 for testing I use localhost you have give you server public Ip address

	   conn,connerr := net.ListenPacket("udp","[2409:40f4:10f7:d8bb:ee1:6a39:8301:9a37]:8080")
	   
	   if connerr != nil {
		      log.Println("Connection err",connerr)
			  return 
	   }

	   defer conn.Close()
	 
	  existerr :=  serversetup.Tun0exists()

	  if existerr != nil {
		    tun,createrr := serversetup.Createtun0()

			if createrr != nil {
                   return 
			}

	    	go serverservices.PacketReceiver(conn,tun)

			go serverservices.Sendresp(conn,tun)

	  }

}