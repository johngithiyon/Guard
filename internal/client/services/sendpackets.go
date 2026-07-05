package services

import (
	"log"

	"github.com/songgao/water"
)

func Sendpackets(tun *water.Interface,packets []byte) error {

	 _,writerr :=   Conn.Write(packets)

	 if writerr != nil {
		    log.Println("Write err in send packets",writerr)
			return writerr
	 }
       
	 return nil
}