package serverservices

import (
	"log"

	"github.com/songgao/water"
)

//Function to write to the guard-server interface

func Writetun0(tun *water.Interface, data []byte) error {

	   lendata,writerr := tun.Write(data)

	   if writerr != nil  {
		     log.Println("Write Error",writerr)
	   }

	   if lendata != len(data){
		    log.Println("Partial Packets",lendata,len(data))
	   }
	   
		return nil 	   
}