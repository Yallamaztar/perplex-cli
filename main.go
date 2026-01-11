package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Yallamaztar/iw4m-go/iw4m"
	"github.com/Yallamaztar/iw4m-go/iw4m/server"
)

func checkStatus(s *server.Server) {
	status, err := s.Status()
	if err != nil {
		fmt.Println("   [ \033[1m\033[35mServer Not Reachable\033[0m ] (\033[31mInactive\033[0m)")
		return
	}

	for _, server := range status {
		fmt.Println("   [\033[1m\033[35m", server.Name, "\033[0m] (\033[32mActive\033[0m)")
		fmt.Printf("   \033[35m%s\033[0m (%s)\n", server.Map.Alias, server.Map.Name)
		fmt.Println("   players online:", len(server.Players), "/", server.MaxPlayers)

		for _, player := range server.Players {
			fmt.Printf("    - %-15s (%dms)\n", player.Name, player.Ping)
		}

		println()
		fmt.Printf("   \033[35mconnect\033[0m \033[1m%s:%d\033[0m\n", server.ListenAddress, server.ListenPort)
		fmt.Println("   " + strings.Repeat("─", len(server.ListenAddress)+len(fmt.Sprintf("%d", server.ListenPort))+9))
	}
}

func main() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("\n   \033[1m\033[35mPerpleX Sniper SnD Status Checker\033[0m")
	fmt.Println("   \033[1mhttps://github.com/Yallamaztar\033[0m")
	fmt.Println("   ───────────────────────────")

	primary := iw4m.NewWrapper(
		"http://193.23.160.188:1624",
		"193231601885151",
		"",
	)

	secondary := iw4m.NewWrapper(
		"http://217.217.243.33:1624",
		"217217243334242",
		"",
	)

	last := iw4m.NewWrapper(
		"http://45.146.253.176:1624",
		"451462531765151",
		"",
	)

	primaryServer := server.NewServer(primary)
	secondaryServer := server.NewServer(secondary)
	lastServer := server.NewServer(last)

	checkStatus(primaryServer)
	checkStatus(secondaryServer)
	checkStatus(lastServer)
}
