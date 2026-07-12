package servermodels

import "time"

//Create a struct to store the client info

type Client struct {
       Clientaddr string //this clientaddr is vpn sever allocated ip 
	   Lastseen time.Time   //time used to check the heartbeat of the client 
}
