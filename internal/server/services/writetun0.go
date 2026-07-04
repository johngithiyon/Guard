package serverservices

import (
	"log"
	"net"

	"github.com/songgao/water"
)

//Function to write to the guard-server interface

func Writetun0(conn net.PacketConn,tun *water.Interface, data []byte) error {

	   _,writerr := tun.Write(data)

	   if writerr != nil {
		     log.Println("Write Error",writerr)
			 return writerr
	   }
       
	  sendresperr :=  Sendresp(conn,tun)

	  if sendresperr != nil {
		   return sendresperr
	  }

		return nil 	   
}