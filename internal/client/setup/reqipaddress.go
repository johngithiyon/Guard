package setup

import (
	"log"
	"net"
)

//Send the request to the Vpn Server

func ReqIpaddress() error   {

	   //make the conn object for Vpn server

	   conn,connerr := net.Dial("udp","127.0.0.1:8080")

       if connerr != nil {
              log.Println("Cannot able to connect  to the server",connerr)
			  return  connerr
	   }

	   //send the ip request to the vpn server

	   conn.Write([]byte("Ip"))

	   buffer := make([]byte,1024)

	   for {
  
		   length,readerr := conn.Read(buffer)

		   if readerr != nil {
			     log.Println("ReadErr for the IP address",readerr)
				 return readerr
		   }

		   log.Println(string(buffer[:length]))

	   }
}