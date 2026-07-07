package serversetup

import (
	"log"

	"github.com/songgao/water"
)

//Function Used to create and write the received packets to the gurad-server interface

func Createtun0() (*water.Interface, error) {

	config := water.Config{
		DeviceType: water.TUN,
	}

	//Name the tun interface as guard0

	config.Name = "guard-server"

	// open the tun interface

	tun, createrr := water.New(config)

	if createrr != nil {
		log.Println("Create err", createrr)
		return nil, createrr
	}

   assignerr := Assignip()

   if assignerr != nil {
       return nil,assignerr
   }

	uperr := Tun0Upstatus()

	if uperr != nil {
		return nil, uperr
	}

	mtuerr := Mtulimit()

	if mtuerr != nil {
		return nil,mtuerr
	}

	return tun, nil
}