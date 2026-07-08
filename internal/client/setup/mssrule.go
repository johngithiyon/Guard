package setup

import (
	"log"
	"os/exec"
)

// Function to set the MSS rule

func Mssrule() error {

	status,mssrulerr  := exec.Command(

		"iptables",
		"-t", "mangle",
		"-A", "OUTPUT",
		"-p", "tcp",
		"--tcp-flags", "SYN,RST", "SYN",
		"-j", "TCPMSS",
		"--clamp-mss-to-pmtu",

	).CombinedOutput()

	if mssrulerr != nil {
		log.Println("Msserr",status)
		return mssrulerr
	}

	return nil 

	
}