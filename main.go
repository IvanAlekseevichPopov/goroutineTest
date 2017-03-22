// Тестирование возможностей горутин
package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch1 := make(chan string, 200000) //Последняя цифра - буффер для асинхронного выполнения
	ch2 := make(chan string, 200000)
	fmt.Println("Start go")
	for i := 0; i < 10; i++ {
		go thread(strconv.Itoa(i), ch1, ch2)
	}

	for i := 0; i < 100; i++ {
		ch1 <- "testString"
		fmt.Println(<-ch2)
	}
}

func thread(goSerial string, ch1 chan string, ch2 chan string) {
	i := 0
	for {
		i++
		str := <-ch1
		ch2 <- "Goroutine #" + goSerial + " " + str + " executed " + strconv.Itoa(i)
	}
}
