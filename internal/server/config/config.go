package serverconfig

import (
	"os"

	servermodels "github.com/johngithiyon/Guard/internal/server/models"
	"gopkg.in/yaml.v3"
)

var Config servermodels.Config

func Loadconfig() error {

	    readata,readerr := os.ReadFile("config.yaml")

		if readerr != nil {
			 return readerr
		}
       
		unmarsherr := yaml.Unmarshal(readata,&Config)

		if unmarsherr != nil {
			return  unmarsherr
		}

	    return nil 
}