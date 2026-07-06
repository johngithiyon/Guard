package setup

import (
	"errors"
	"log"
	"net"
)

//Send the request to the Vpn Server

func ReqIpaddress(conn net.Conn) error   {

	   //send the ip request to the vpn server

	   conn.Write([]byte("Ip"))

	   buffer := make([]byte,1024)

	   //Reading from the Socket For the Ip 

	   for {
  
		   length,readerr := conn.Read(buffer)

		   if readerr != nil {
			     log.Println("ReadErr for the IP address",readerr)
				 return readerr
		   }

		   //Check if there is no Ip in the server

		   if string(buffer[:length]) != "Ip insufficient" {

			        //Call the assign ip function 
			            
			       assignerr :=  Assignip(string(buffer[:length]))

				   if assignerr != nil {
					   return assignerr
				   }

				   //After call the ip assign function it have to terminate the loop

				   return nil 
			 
		   } else {
			    return errors.New("Ip Insufficient")
		   }

	   }
}