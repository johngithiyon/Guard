package main

import (
	"log"

	serverconfig "github.com/johngithiyon/Guard/internal/server/config"
	"github.com/johngithiyon/Guard/internal/server/services"
	serversetup "github.com/johngithiyon/Guard/internal/server/setup"
)

func main() {

	  loadconfigerr := serverconfig.Loadconfig()

	  if loadconfigerr != nil {
		  log.Println("Load config err",loadconfigerr)
		  return 
	  }

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

			//Goroutine to check the client is them alive or not 

			go serverservices.Checkalive()

	  }

	  //make the process keep alive 

	  select{}

}