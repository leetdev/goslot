package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/leetdev/goslot"
)

func main() {
	fmt.Println("Press any key to stop")

	slot := goslot.New(
		goslot.Config{
			MinDelay:   5 * time.Millisecond,
			MaxDelay:   50 * time.Millisecond,
			StopTime:   5 * time.Second,
			NumReels:   10,
			NumSymbols: 10,
		},
	)
	slot.Start()

	stopping := false
	go func() {
		_, _, _ = keyboard.GetSingleKey()
		stopping = true
	}()

	for {
		fmt.Print("\r")
		printValues(slot)

		if stopping {
			slot.Stop()
			stopping = false
		}

		if !slot.IsRunning() {
			break
		}

		time.Sleep(10 * time.Millisecond)
	}

	fmt.Print("\r")
	printValues(slot)
	fmt.Println()
}

func printValues(s *goslot.GoSlot) {
	values := s.Values()
	for _, value := range values {
		fmt.Print(value)
	}
}
