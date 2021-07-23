package main

import "log"

var Revision string

func main() {
	log.Println("revision:", Revision)
	Init()
	Run()
}
