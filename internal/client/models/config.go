package models

type Config struct {
	  
	     VpnServer string `yaml:"vpnserver_ip"`
		 VpnServerPort string  `yaml:"vpnserver_port"`
		 SshServer string `yaml:"sshserver_ip"`
}