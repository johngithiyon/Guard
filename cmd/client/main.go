package main

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/services"
	"github.com/johngithiyon/Guard/internal/client/setup"
	"github.com/johngithiyon/Guard/internal/config"
)

func main() {

	connerr := services.CreateConnection()

	if connerr != nil {
		 log.Println("Connection err",connerr)
		 return
	}

	 loaderr :=  config.Loadenv()

	 if loaderr != nil {
		  log.Println("Load Error",loaderr)
		  return 
	 }

	  existerr :=  setup.Tun0exists()

	  if existerr != nil {
		 services.CreateReadtun0()
	  }
}