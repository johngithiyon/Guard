package services

import "time"

//Function Used to send the alive packets every 20 seconds to prove the client is alive

func Keepalive()  {

	for {
     
	 //It wait for the 20 seconds

     time.Sleep(20 * time.Second)

	 Conn.Write([]byte("Ping"))

	} 
	    
}

