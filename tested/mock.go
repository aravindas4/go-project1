package tested

import (
	"fmt"
	"io"
)

func CountDown(buffer io.Writer) {
	for count := 3; count > 0; count-- {
		// time.Sleep(1 * time.Second)
		fmt.Fprintln(buffer, count)
	}
	// time.Sleep(1 * time.Second)
	fmt.Fprint(buffer, "Go!")
}
