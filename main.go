package main

import (
	"log"
	"net"
	"os"
	// b "github.com/christinavanmoof/generatePasswords/passwords"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	println(lis)
	if err != nil {
		log.Printf("Error listening: %s", err.Error())
		os.Exit(1)
	}

	// c := &b.GeneratePasswords{
	// 	Length: 8,
	// }

}
