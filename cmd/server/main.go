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

	   conn,connerr := net.ListenPacket("udp",":8080")
	   
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

			//Goroutine to receive the packets from the clients

	    	go serverservices.PacketReceiver(conn,tun)

			//Goroutine to send the response to the clients

			go serverservices.Sendresp(conn,tun)

	  }

	  //make the process keep alive 

	  select{}

}