package serverservices

import (
	"log"
	"net"
	"github.com/songgao/water"
)

//Used to receive the packet from the connection

func PacketReceiver(conn net.PacketConn,tun *water.Interface) error {

	// Creating a Byte Slice to store the packets

	buffer := make([]byte,1024)

	for {

		// Read the packets from the connection
		 
		 length,addr,readerr := conn.ReadFrom(buffer)

		 if readerr != nil {
              log.Println("Read Error",readerr)
			  continue // I did not have to stop the loop so that other client did not suffer
		 }

		 if string(buffer[:length]) == "Ip" {
			   Ipallocator(conn,addr)
		 } else {
         		 Writetun0(tun,buffer[:length])
		 }
         		
	}
}
