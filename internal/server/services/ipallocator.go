package services

import (
	"log"
	"net"
)

//centralized slice for the ip list

var ips = []string {

	 "10.0.0.1",
	 "10.0.0.2",
	 "10.0.0.3",
} 

//Function for the ipallocation

//Using queue metheod for the o(1) ip allocation

func Ipallocator(conn net.PacketConn,addr net.Addr)  {

	     if len(ips) !=0 {

	     conn.WriteTo([]byte(ips[0]),addr)

		 ips = ips[1:]
	  
		 log.Println(ips)

		 } else {
			conn.WriteTo([]byte("Ip insufficient"),addr)
		 }
}