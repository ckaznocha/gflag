package main

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"

	"github.com/ckaznocha/gflag"
)

func main() {
	gflag.SetFlagSet(pflag.CommandLine)
	salutation := gflag.Define("salutation", "Hello", "a salutation")
	subject := gflag.Define("subject", "user", "a subject to greet")
	wait := gflag.Define("wait", 1*time.Second, "how long  to  wait before greeting")
	pflag.Parse()

	time.Sleep(*wait)

	fmt.Printf("%s, %s\n", *salutation, *subject)
}
