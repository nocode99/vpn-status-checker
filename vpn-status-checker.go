package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		ReadVpnLog(os.Args[1])
	} else {
		fmt.Println("Error! Not enough arguments! Please input log file path")
	}
}

func ReadVpnLog(input string) {
	var SWITCH bool = false
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "Updated") {
			line2 := strings.Split(sc.Text(), ",")
			timestamp := line2[1]
			fmt.Println(timestamp)
		}
		if strings.HasPrefix(sc.Text(), "ROUTING TABLE") {
			SWITCH = false
		}
		if SWITCH {
			connection := strings.Split(sc.Text(), ",")
			user, ip_port, date := connection[0], connection[1], connection[4]
			ip := strings.Split(ip_port, ":")[0]
			fmt.Printf("%s\t%s\t%s\n", user, ip, date)
		}
		if strings.HasPrefix(sc.Text(), "Common") {
			SWITCH = true
		}
	}
}
