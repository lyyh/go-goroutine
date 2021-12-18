package channelrun

import "fmt"

// 关闭通道
func DoCloseChannel() {
	ch1 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	for {
		if data, ok := <-ch1; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
}

// 优雅的取值，range
func RecvValFromChannel() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for {
			i, ok := <-c1
			if !ok {
				c2 <- i * i
			}
			close(c2)
		}
	}()

	for i := range c2 {
		fmt.Println(i)
	}
}

// 写通道
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 单向通道
func SingleTrack() {
	c1 := make(chan int)
	c2 := make(chan int)

	go counter(c1)
	go square(c2, c1)
	printer(c2)
}
