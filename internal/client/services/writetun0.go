package services

import (
	"log"

	"github.com/songgao/water"
)

//Function Used to write the packets to the tun0 interface

func Writetun0(tun *water.Interface,data []byte) error {

	lendata,writerr := tun.Write(data)
        
	if writerr != nil {
		 log.Println("Writerr in tun interface",writerr)
		 return writerr
	}

	if lendata != len(data){
		log.Println("Partial Packets",lendata,len(data))
   }
	
    return nil 
}