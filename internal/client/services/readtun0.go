package services

import (
	"log"

	"github.com/songgao/water"
)

func Readtun0(tun *water.Interface) error {
	 
	buffer := make([]byte,65535)

	   
	for {
		
		 length,readerr :=  tun.Read(buffer)

		 if readerr != nil {
			 return readerr
		 }

		 if length > 0 {
			  log.Println("Recieved packets in client",length,string(buffer[:length]))
			  Sendpackets(tun,buffer[:length])
		 }
	   
	}   
}