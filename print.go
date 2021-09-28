package cbg_notify

import (
	"fmt"
	"strings"
)

var (
	statusWorking = 0
)

func PrintProgress() {
	if statusWorking == 3 {
		statusWorking = 1
		fmt.Printf(strings.Repeat(" ", 3) + "\r")
	} else {
		statusWorking++
	}
	fmt.Printf(strings.Repeat(".", statusWorking) + "\r")
}
