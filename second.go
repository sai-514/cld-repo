package main

import (
	"fmt"
	"time"
)

func main() {
	Now := time.Now()
	hr, min, sec := Now.Clock()
	fmt.Println("Year   :=", Now.Year())
	fmt.Println("Month  :=", Now.Month())
	fmt.Println("Day    :=", Now.Day())
	fmt.Println("Hour   :=", hr)
	fmt.Println("Min    :=", min)
	fmt.Println("Sec    :=", sec)
}
