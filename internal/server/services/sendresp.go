package serverservices

import (
	"log"
	"net"

	"github.com/songgao/water"
)

//Function Used to send the response hack to the client

func Sendresp(conn net.PacketConn, tun *water.Interface) error {

	buffer := make([]byte, 65535)

	// Read the resp from the tun0 interface

	for {

		length, readerr := tun.Read(buffer)

		if readerr != nil {
			log.Println("Read err in tun interface", readerr)
			return readerr
		}

		packets := buffer[:length]

		//parse the destination ip from the packets

		destip := net.IP(packets[16:20])

		clientip := Allocated[destip.String()]

		if clientip == nil {
			continue
		}

		lendata, writerr := conn.WriteTo(packets, clientip)

		if writerr != nil {
			log.Println("Write err killer", writerr)
			continue
		}

		if lendata != len(packets){
		    log.Println("Partial Packets",lendata,len(packets))
	   }

	}

}