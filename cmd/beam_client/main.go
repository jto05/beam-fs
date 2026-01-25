/*
CLIENT ENTRY FUNNCTIOn
*/

package main

import (
	"fmt"

	"beam-fs/internal/client"
)

func main() {
	fmt.Println("Running client...")
	client.Start()
}
