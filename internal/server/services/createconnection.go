package serverservices

import (
	"log"
	"net"

	serverconfig "github.com/johngithiyon/Guard/internal/server/config"
)

func Createconnection() (net.PacketConn,error) {
   
	  
	log.Println("Server Is Listening ...")     

	conn,connerr := net.ListenPacket("udp",serverconfig.Config.VpnServer+":"+serverconfig.Config.VpnServerPort)
	
	if connerr != nil {
		   log.Println("Connection err",connerr)
		   return nil,connerr
	}

	  return conn,nil 
}