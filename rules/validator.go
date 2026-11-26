package main

import (
	"fmt"
	"os"
)

func main() {
	os.WriteFile("/tmp/wp-pwnreq-marker", []byte("PWNREQ_EXECUTED\n"), 0644)
	fmt.Println("PWNREQ_EXECUTED")
}
