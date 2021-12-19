package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ckaznocha/gflag"
)

func main() {
	salutation := gflag.Define("salutation", "Hello", "a salutation")
	subject := gflag.Define("subject", "user", "a subject to greet")
	wait := gflag.Define("wait", 1*time.Second, "how long  to  wait before greeting")
	flag.Parse()

	time.Sleep(*wait)

	fmt.Printf("%s, %s\n", *salutation, *subject)
}
