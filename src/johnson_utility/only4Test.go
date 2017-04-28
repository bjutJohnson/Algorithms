package johnson_utility

import (
	"fmt"
	//	"time"
	"sync"
)

type TestStruct struct {
	lock *sync.Mutex
	str  string
}

var strChan chan string = make(chan string, 100)

func NewTestStruct() *TestStruct {
	return &TestStruct{new(sync.Mutex), ""}
}

func (p *TestStruct) SetString(str string) {
	p.lock.Lock()
	defer func() {
		p.lock.Unlock()
	}()

	strChan <- "start: " + str
	p.str = str
	strChan <- "end: " + str
}

func (s TestStruct) GetString() string {
	s.lock.Lock()
	defer func() {
		s.lock.Unlock()
	}()

	return s.str
}

func (s TestStruct) PrintString() {
	s.lock.Lock()
	defer func() {
		s.lock.Unlock()
	}()

	fmt.Println(s.str)
}

func Test_Channel() {

	p := NewTestStruct()

	for i := 0; i < 100; i++ {
		go p.SetString("abc")
		go p.SetString("cde")
	}

	j := 0

	for {
		select {
		case x := <-strChan:
			fmt.Println(x)
			j++
		default:
			if j >= 400 {
				return
			}
		}
	}
}

var c chan bool = make(chan bool)

func func1() {
	fmt.Println("func1")
	c <- true
}

func func2() {
	fmt.Println("func2")
	c <- true
}

func func3() {
	fmt.Println("func3")
	c <- true
}

func Test_func() {
	go func1()
	<-c
	go func2()
	<-c
	go func3()
	<-c
}
