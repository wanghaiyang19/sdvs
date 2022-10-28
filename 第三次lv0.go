//为啥学者很简单，但用起来很难。崩溃大哭了！！！
package main

import (
	"fmt"
	"time"
)

func main() {
	repeater()
}

func repeater() {
	for {
		time.Sleep(time.Second)
		fmt.Println("over.")
	}
}
func main() {
	go repeater()
	go repeater()
	repeater()
}

func repeater() {
	for {
		time.Sleep(time.Second)
		fmt.Println("over.")
	}
}
func main() {
	go learnFrontend()
	go learnAndroid()
	go learnMachineLearning()
	learnBackend()
}

func learnBackend() {
	time.Sleep(10 * time.Minute)
	fmt.Println("会了！")
}

func learnFrontend() {
	time.Sleep(time.Nanosecond)
	fmt.Println("会了！")
}

func learnAndroid() {
	time.Sleep(20 * time.Minute)
	fmt.Println("悔了！")
}

func learnMachineLearning() {
	time.Sleep(114514 * time.Minute)
	fmt.Println("废了！")
}

type Item struct {
	Name  string
	Count int
}

func main() {
	pipeline := make(chan Item, 10)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pipeline <- Item{
				Name:  "螺丝",
				Count: 5,
			}
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pipeline <- Item{
				Name:  "齿轮",
				Count: 3,
			}
		}
	}()
	for {
		item := <-pipeline
		fmt.Printf("%#v\n", item)
	}
}

//1
func main() {
	go 打电动()
	好康的 := 欢迎来我家玩()
	fmt.Println(好康的)
}

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() {
	fmt.Println("输了啦，都是你害的")
}

//2
func main() {
	var 杰哥答应的好康的 = make(chan string)
	go func() {
		杰哥答应的好康的 <- 欢迎来我家玩()
	}()
	打电动()
	好康的 := <-杰哥答应的好康的
	fmt.Println(好康的)
}

func 欢迎来我家玩() string {
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() {
	fmt.Println("输了啦，都是你害的")
}
func main() {
	var (
		ch1 = make(chan struct{})
		ch2 = make(chan struct{})
		ch3 = make(chan struct{})
	)
	go handleCh1(ch1)
	go handleCh2(ch2)
	go handleCh3(ch3)
	for {
		select {
		case _ = <-ch1:
			fmt.Println("get from ch1")
		case _ = <-ch2:
			fmt.Println("get from ch2")
		case _ = <-ch3:
			fmt.Println("get from ch3")
		}
	}
}

func handleCh1(ch1 chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}
}

func handleCh2(ch2 chan struct{}) {
	for {
		time.Sleep(4 * time.Second)
		ch2 <- struct{}{}
	}
}

func handleCh3(ch3 chan struct{}) {
	for {
		time.Sleep(2 * time.Second)
		ch3 <- struct{}{}
	}
}
func main() {
	var (
		ch1  = make(chan struct{})
		stop = make(chan struct{})
	)
	go handleCh1(ch1)
	go func() {
		time.Sleep(10 * time.Second)
		stop <- struct{}{}
	}()
LOOP:
	for {
		select {
		case _ = <-ch1:
			fmt.Println("get from ch1")
		case _ = <-stop:
			break LOOP
		}
	}
}

func handleCh1(ch1 chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}
}
