package services

import (
	"log"
	"net"
)

var ips = map[string]bool {

	 "10.0.0.1":false,
	 "10.0.0.2":false,
	 "10.0.0.3":false,
} 

func Ipallocator(conn net.PacketConn,addr net.Addr)  {

	     for index,_ := range ips{

			     if ips[index] != true{
					conn.WriteTo([]byte(index),addr)
					ips[index] = true 
					break
				 } 
		 }

		 log.Println(ips)

}