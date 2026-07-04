package services

import (
	"log"

	"github.com/songgao/water"
)

//Function Used to read the resp from the udp socket

func Readresp(tun *water.Interface) error  {
        
	    buffer := make([]byte,1024)

		for {

        length,readerr := Conn.Read(buffer)

		if readerr != nil {
			  log.Println("Read err in resp",readerr)
			  return readerr
		}

		tun.Write(buffer[:length])

	    return nil

		}  
}