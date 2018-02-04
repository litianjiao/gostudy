package main

import "fmt"

func send(ch chan int,exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch<-i
	}
	close(ch)
	var a struct{}
	exitChan<-a
	fmt.Println("send exit")
}

func recv(ch chan int, exitChan chan struct{}) {
	for{
		v,ok:=<-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan<-a
}

func main() {
	ch:=make(chan int,10)
	exitChan:=make(chan struct{},2)

	go send(ch,exitChan)
	go recv(ch,exitChan)

	//通过计数（chan总数）来进行同步控制
	var total=0
	for _ = range exitChan {
		total++
		if total==2 {
			break
		}
	}


}

