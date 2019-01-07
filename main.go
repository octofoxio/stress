package main

import (
	"flag"
	"fmt"
	"github.com/stellar/go/keypair"
	"runtime"
	"time"
)

type Result struct {
	seed string
	addr string
}

var (
	count = 0
)

func Run(c chan Result, in string) {
	var seed string
	var addr string
	var prefix string

	for prefix != in {
		kp, _ := keypair.Random()
		addr = kp.Address()
		seed = kp.Seed()
		prefix = string([]rune(addr)[len(addr)-len(in) : len(addr)])
		count++
		//if ok, _ := regexp.MatchString("OFOX$", prefix); ok {
		//	fmt.Printf("%s\n%s\n\n\n", addr, seed)
		//}
	}
	c <- Result{
		seed: seed,
		addr: addr,
	}
}
func main() {

	prefix := flag.String("p", "UNSET", "Prefix to find")

	flag.Parse()

	found := make(chan Result)

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Printf("Running at %d hash per second\n", count)
				count = 0
			}
		}
	}()

	fmt.Printf("Finding Address end with %s\n", *prefix)
	for i := 0; i < runtime.NumCPU(); i++ {
		fmt.Printf("Running task no. %d \n", i+1)
		go Run(found, *prefix)
	}

	result := <-found
	fmt.Println("Hash found !!")
	fmt.Println(result.addr)
	fmt.Println(result.seed)

}
