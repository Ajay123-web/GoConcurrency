package main

import (
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	sleepTime = 0 * time.Second

	for i := 0; i < 100; i++ {
		main()
		if len(finishOrder) != 5 {
			t.Error("****ERROR****")
		}
		finishOrder = []string{}
	}
 	
}