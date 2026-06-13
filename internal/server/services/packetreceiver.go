package services

import (
	"log"
	"net"
)

//Used to receive the packet from the connection 

func PacketReceiver(conn net.PacketConn) error {

	// Creating a Byte Slice to store the packets

	buffer := make([]byte,1024)

	for {

		// Read the packets from the connection
		 
		 length,addr,readerr := conn.ReadFrom(buffer)

		 if readerr != nil {
              log.Println("Read Error",readerr)
			  continue // I did not have to stop the loop so that other client did not suffer
		 }

		 log.Println(addr,string(buffer[:length]))
	}
}
