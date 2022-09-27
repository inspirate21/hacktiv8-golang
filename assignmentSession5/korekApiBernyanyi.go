package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Korek struct {
	count  int
	player string
}

const breakpoint = 13

type Player struct {
	count int
}

func main() {
	korek := make(chan *Korek)
	done := make(chan *Korek)

	players := []string{"player 1", "player 2", "player 3", "player 4"}
	for _, player := range players {
		go play(player, korek, done)
	}

	korek <- new(Korek)

	finish(done)
}

func play(player string, korek, done chan *Korek) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	for {
		select {
		case k := <-korek:
			v := rand.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)
			k.count++
			k.player = player
			fmt.Println("korek ada di", k.player, "pada hit ke", k.count, "dengan breakpoint = ", breakpoint, "dan mempunyai nilai", v)

			if v%breakpoint == 0 {
				done <- k

				return
			}

			korek <- k
		}
	}
}

func finish(done chan *Korek) {
	for {
		select {
		case d := <-done:
			fmt.Println(d.player, "kalah pada hit ke", d.count)
			return
		}
	}
}
