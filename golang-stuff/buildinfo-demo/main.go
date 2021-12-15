package main

import (
	"debug/buildinfo"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("pass exactly one argument")
	}
	binPath := os.Args[1]
	bin, err := os.Open(binPath)
	if err != nil {
		panic(err)
	}

	buildInfo, err := buildinfo.Read(bin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", buildInfo)

	for _, dep := range buildInfo.Deps {
		fmt.Printf("%+v\n", dep)
	}
}
