package services

import (
	"log"
)

func Sendpackets(packets []byte) error {

	 _,writerr :=   Conn.Write(packets)

	 if writerr != nil {
		    log.Println("Write err in send packets",writerr)
			return writerr
	 }
       
	 return nil
}