package services

import (
	"log"
	"net"

	"github.com/johngithiyon/Guard/internal/client/config"
)

var Conn net.Conn

var Connerr error

func CreateConnection() error  {
	 
	   //make the conn object for Vpn server

	   Conn,Connerr = net.Dial("udp",config.Config.VpnServer+":"+config.Config.VpnServerPort)

	   if Connerr != nil {
		     log.Println("Connection error when send packets",Connerr)
		     return Connerr
	   }

       return nil	      
}