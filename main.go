package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true

	eludedGuards := rand.Intn(100)
	// trespassing
	if isHeistOn && eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	// going in
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened")
	}

	leftSafely := rand.Intn(5)
	// leaving the scene
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Failed Heist")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("The floor was slippery you couldn't run away")
		case 3:
			isHeistOn = true
			fmt.Println("Turns out that you're faster than the guards")
		case 4:
			isHeistOn = true
			fmt.Println("Turns out the bank didn't care about you")
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Printf("The amount taken by this heist was $%d", amtStolen)
		}
	}

	fmt.Println(isHeistOn)
}
