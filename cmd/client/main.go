package main

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/services"
	"github.com/johngithiyon/Guard/internal/client/setup"
)

func main() {

	connerr := services.CreateConnection()

	if connerr != nil {
		 log.Println("Connection err",connerr)
		 return
	}

      existerr :=  setup.Tun0exists()

	  if existerr != nil {
		 services.CreateReadtun0()
	  }
}