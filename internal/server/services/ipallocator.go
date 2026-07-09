package serverservices

import (
	"net"
	"time"

	"github.com/johngithiyon/Guard/internal/server/models"
)

//centralized slice for the ip list

var Ips = []string{
	   "10.0.0.2",
	   "10.0.0.3",
}

//Store the allocated ip and addr of the client in map

var Allocated = make(map[string]net.Addr)

//Store the heartbeat for the clients 

var Clients = make(map[string]*models.Client)

//Function for the ipallocation

//Using queue metheod for the o(1) ip allocation

func Ipallocator(conn net.PacketConn,addr net.Addr)  {

	     if len(Ips) !=0 {

	     conn.WriteTo([]byte(Ips[0]),addr)

		 Allocated[Ips[0]] = addr

		client := &models.Client{
              Clientaddr: Ips[0],
			  Lastseen: time.Now(),
		}

		//store the client udp address and give the value as models.client 

		Clients[addr.String()] = client

		Ips = Ips[1:]

		 } else {
			conn.WriteTo([]byte("Ip insufficient"),addr)
			return
		 }
}