package main

import (
	"fmt"
	"log"
	"metasploit_minimal/rpc"
)

func main() {

	host := "192.168.1.104:55552"
	pass := "s3cr3t"

	// host := os.Getenv("MSFHOST")
	// pass := os.Getenv("MSFPASS")

	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variable MSFHOST or MSFPASS")
	}

	msf, err := rpc.New(host, user, pass)

	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.SessionList()

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Sessions:")

	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
