package main

import (
	"fmt"
	"log"
	"lol_helper/lib/windows"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	if windows.MessageBoxYesNo("xx助手", "是否开启xx?") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
