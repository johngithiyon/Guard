package serverservices

import (
	"log"
	"time"
)

//Function used to check and remove the disconnected client

func Checkalive() {

	       //Wait for every 5 seconds 
	   
	       time.Sleep(5 * time.Second)

		   for _,client  := range Clients {

			    //Set the timeout for the client is 30 seconds

			    if time.Since(client.Lastseen) > 30 * time.Second {

					   log.Println("client disconnected")
					    
					   delete(Allocated,client.Clientaddr)
					   Ips = append(Ips, client.Clientaddr)
				}
		   }
}