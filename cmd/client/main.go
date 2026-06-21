package main

import (
	"github.com/johngithiyon/Guard/internal/client/services"
	"github.com/johngithiyon/Guard/internal/client/setup"
)

func main() {

	  existerr :=  setup.Tun0exists()

	  if existerr != nil {
		 services.CreateReadtun0()
	  }
}