package main

import (
	"time"
)

const (
	idCheckTmConst = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmConst))
	println("\tidCheck ok")
	return idCheckTmConst
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tbodyCheck ok")
	return bodyCheckTmCost
}


func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\txRayCheckTmCost ok")
	return xRayCheckTmCost
}


func airportSecurityCheck() int {
	println("airportSecurityCheck...")
	total := 0

	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()

	println("airportSecurityCheck ok")
	return total
}


func main() {
	total := 0
	passengers := 30
	for i := 0; i < passengers; i++ {
		total += airportSecurityCheck()
	}

	println("total time cost:", total)
}