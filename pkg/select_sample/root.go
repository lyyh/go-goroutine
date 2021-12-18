package selectsample

import (
	"fmt"
	"time"
)

func ExecSelect() {
	output := make(chan string)
	go write(output)
	for data := range output {
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write")
		default:
			fmt.Println("channel full")
		}

		time.Sleep(time.Millisecond * 500)
	}
}
