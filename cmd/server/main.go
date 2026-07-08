package main

import (
	"github.com/johngithiyon/Guard/internal/server/services"
	serversetup "github.com/johngithiyon/Guard/internal/server/setup"
)

func main() {

	  conn,connerr := serverservices.Createconnection()

	  if connerr != nil {
		   return
	  }

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