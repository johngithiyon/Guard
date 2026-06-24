package serverservices

import (
	"log"

	"github.com/songgao/water"
)

//Function to write to the guard-server interface

func Writetun0(tun *water.Interface, data []byte) error {

	   n,writerr := tun.Write(data)

	   if writerr != nil {
		     log.Println("Write Error",writerr)
			 return writerr
	   }

	   log.Println("number of bytes writted",n)

	   return nil 
}