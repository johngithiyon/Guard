package services

import (
	"log"
	"net"
)

var Conn net.Conn

var Connerr error

func CreateConnection() error  {
	 
	   //make the conn object for Vpn server

	   Conn,Connerr = net.Dial("udp","127.0.0.1:8080")

	   if Connerr != nil {
		     log.Println("Connection error when send packets",Connerr)
		     return Connerr
	   }

       return nil	      
}