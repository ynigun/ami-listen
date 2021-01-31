package main

import (
	"fmt"

	"github.com/kcasctiv/amichan"
)

func main() {
	port := 7080
	keepalive := true
 	user := "asterisk"
  	pass :=  "asterisk"
  	ip := "127.0.0.1"
	
	a := amichan.New(user,pass,ip , port, keepalive)
	a.Connect()

	for {
		select {
		case err := <-a.Err():
			fmt.Println(err)
		case event := <-a.Event():
			if event.Name() == "DTMFEnd" {
				

				num, ok := event.Field("CallerIDNum")
				if ok {
					fmt.Println(num)
				}
        
				Digit, ok := event.Field("Digit")
				if ok {
					fmt.Println(Digit)
				}
        
				ID, ok := event.Field("Uniqueid")
				if ok {
					fmt.Println(ID)
				}
				
			}
		}
	}
}
