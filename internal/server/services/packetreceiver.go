package serverservices

import (
	"log"
	"net"
	"time"

	"github.com/songgao/water"
)

//Used to receive the packet from the connection

func PacketReceiver(conn net.PacketConn,tun *water.Interface) error {

	// Creating a Byte Slice to store the packets

	buffer := make([]byte,65535)

	for {

		// Read the packets from the connection
		 
		 length,addr,readerr := conn.ReadFrom(buffer)

		 if readerr != nil {
              log.Println("Read Error",readerr)
			  continue // I did not have to stop the loop so that other client did not suffer
		 }


		 if string(buffer[:length]) == "Ip" {
			   Ipallocator(conn,addr)
			   continue
		 } else if string(buffer[:length]) == "Ping" {
                client,ok := Clients[addr.String()]

				if ok {
                    client.Lastseen = time.Now()
				}
				continue
		 } 

		 client,ok := Clients[addr.String()]

		 if ok {
			 client.Lastseen = time.Now()
		 }

         Writetun0(tun,buffer[:length])     		
	}
}
