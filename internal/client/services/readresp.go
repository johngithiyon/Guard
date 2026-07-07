package services

import (
	"log"

	"github.com/songgao/water"
)

//Function Used to read the resp from the udp socket

func Readresp(tun *water.Interface) error  {
        
	    buffer := make([]byte,65535)

		for {

        length,readerr := Conn.Read(buffer)

		if readerr != nil {
			  log.Println("Read err in resp",readerr)
			  return readerr
		}
   
		writrerr := Writetun0(tun,buffer[:length])

		if writrerr != nil {
			log.Println("Write err in Read Resp",writrerr)
		}
	 
		}  
}