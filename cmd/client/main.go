package main

import "github.com/johngithiyon/Guard/internal/client/app"

func main() {
         runerr := app.Run()

		 if runerr != nil {
			   return
		 }
}