package serverservices

import (
	"log"
	"net"

	"github.com/songgao/water"
)

//Function Used to send the response hack to the client

func Sendresp(conn net.PacketConn,tun *water.Interface) error {

	buffer := make([]byte,1024)

	   // Read the resp from the tun0 interface

	   for {
            
		      length,readerr := tun.Read(buffer)

			  if readerr != nil {
				   log.Println("Read err in tun interface",readerr)
				   return readerr
			  }

			  packets := buffer[:length]

			  //parse the destination ip from the packets

			  destip := net.IP(packets[16:20])

			  log.Println(destip)

			 clientip :=  Allocated[destip.String()]

			_,writerr :=  conn.WriteTo(packets,clientip)

			if writerr != nil {
				log.Println("Write err",writerr)
				return writerr
			}
            
			return nil 
	   }

}