package app

import (
	"log"

	serverconfig "github.com/johngithiyon/Guard/internal/server/config"
	serverservices "github.com/johngithiyon/Guard/internal/server/services"
	serversetup "github.com/johngithiyon/Guard/internal/server/setup"
)

func Run() error {
	    
	loadconfigerr := serverconfig.Loadconfig()

	if loadconfigerr != nil {
		log.Println("Load config err",loadconfigerr)
		return loadconfigerr
	}

	conn,connerr := serverservices.Createconnection()

	if connerr != nil {
		 return connerr
	}

	existerr :=  serversetup.Tun0exists()

	if existerr != nil {
		  tun,createrr := serversetup.Createtun0()

		  if createrr != nil {
				 return createrr
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