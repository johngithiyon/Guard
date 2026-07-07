package services

import (
	"log"

	"github.com/songgao/water"
)

func Sendpackets(tun *water.Interface,packets []byte) error {

	 lendata,writerr := Conn.Write(packets)

	 if writerr != nil {
		    log.Println("Write err in send packets",writerr)
	 }

	 if lendata != len(packets){
		log.Println("Partial Packets",lendata,len(packets))
   }
       
	 return nil
}