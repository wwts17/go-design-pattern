package main

import (
	"fmt"
	"time"
)

type Function func(float64) float64

// the decorator function
func ProfileDecorator(fn Function) Function {
	fmt.Printf("%s log:%s \n", time.Now().Format(time.RFC822), "decorator start!")
	return func(params float64) float64 {
		start := time.Now()
		result := fn(params)
		elapsed := time.Now().Sub(start)
		fmt.Println("Function completed with time:", elapsed)
		return result
	}
}

func client() {
	decoratedSquareRoot := ProfileDecorator(func(param float64) float64 {
		time.Sleep(2 * time.Second)
		return param + 64
	})
	decoratedSquareRoot(20)
}

func main() {
	client()
}
