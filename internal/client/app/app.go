package app

import (
	"log"

	"github.com/johngithiyon/Guard/internal/client/config"
	"github.com/johngithiyon/Guard/internal/client/services"
	"github.com/johngithiyon/Guard/internal/client/setup"
)

func Run() error {
        
	    
	loadconfigerr := config.Loadconfig()

	if loadconfigerr != nil {
		  log.Println("load config err",loadconfigerr)
		  return  loadconfigerr
	}

	connerr := services.CreateConnection()

	if connerr != nil {
		 log.Println("Connection err",connerr)
		 return connerr
	}

      existerr :=  setup.Tun0exists()

	  if existerr != nil {
		  createrr := services.Createtun0()

		  if createrr != nil {
			    log.Println("Create tun err",createrr)
			    return createrr
		  }
	  }
     
	  //make the process keep alive 

	  select{}
}