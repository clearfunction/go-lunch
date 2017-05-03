package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("The time is now %v", time.Now())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nPress ENTER to close...")
	reader.ReadString('\n')
}
