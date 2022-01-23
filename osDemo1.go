package main

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
)

func main() {
	f, err := exec.LookPath("main")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)

	u, _ := user.Current()
	log.Println("user name: ", u.Username)
	log.Println("user id: ", u.Uid)
	log.Println("user dir: ", u.HomeDir)
	log.Println("main id : ", u.Gid)

	s, _ := u.GroupIds()
	log.Println("all main id: ", s)
}
