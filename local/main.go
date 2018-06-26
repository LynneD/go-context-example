package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ctx := context.Background()
	ctx0, cancel := context.WithCancel(ctx) //parent context ctx0
	ctx1, cancel1 := context.WithTimeout(ctx0, 2 * time.Second)
	ctx2, cancel2 := context.WithTimeout(ctx0, 1 * time.Second)
	ctx3, cancel3 := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel1()
	defer cancel2()
	defer cancel3()
	go buyTomato(ctx1)
	go buyMeat(ctx2)
	go buyVeg(ctx3)
	cancel() // when call this both ctx1 and ctx2 canceled
	time.Sleep(5 * time.Second)
}

func buyTomato(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("buy Tomato time out")
		break
	case <-time.After(2 * time.Second): // can't be default since it will not block

		fmt.Println("buy tomatos")
	}
}

func buyMeat(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("buy meat time out")
		break
	case <-time.After(5 * time.Second):

		fmt.Println("buy meat")
	}
}

func buyVeg(ctx context.Context) {
	Loop:
	for i := 0; ; i++ {
		select {
			case <-ctx.Done():
				fmt.Println("buyVeg time out")
				break Loop
			default:                        // can be default since having a for loop outside
				time.Sleep(time.Second)
				fmt.Printf("buyVeg %v\n", i)
		}
	}
}
