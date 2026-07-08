package serverservices

import (
	"log"
	"net"
)

func Createconnection() (net.PacketConn,error) {
   
	  
	log.Println("Server Is Listening ...")     

	// Creates a Udp socket listening on the port 8080 for testing I use localhost you have give you server public Ip address

	conn,connerr := net.ListenPacket("udp",":8080")
	
	if connerr != nil {
		   log.Println("Connection err",connerr)
		   return nil,connerr
	}

	  return conn,nil 
}