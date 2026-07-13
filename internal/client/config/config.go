package config

import (
	"os"

	"github.com/johngithiyon/Guard/internal/client/models"
	"gopkg.in/yaml.v3"
)

var Config models.Config

func Loadconfig() error {

	    readata,readerr := os.ReadFile("../../internal/client/config/config.yaml")

		if readerr != nil {
			 return readerr
		}
       
		unmarsherr := yaml.Unmarshal(readata,&Config)

		if unmarsherr != nil {
			return  unmarsherr
		}

	    return nil 
}