package main

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/config"
	"github.com/johngithiyon/Guard/internal/client/services"
	"github.com/johngithiyon/Guard/internal/client/setup"
)

func main() {

	loadconfigerr := config.Loadconfig()

	if loadconfigerr != nil {
		  log.Println("load config err",loadconfigerr)
		  return  
	}

	connerr := services.CreateConnection()

	if connerr != nil {
		 log.Println("Connection err",connerr)
		 return
	}

      existerr :=  setup.Tun0exists()

	  if existerr != nil {
		  createrr := services.Createtun0()

		  if createrr != nil {
			    log.Println("Create tun err",createrr)
			    return
		  }
	  }
     
	  //make the process keep alive 

	  select{}
}