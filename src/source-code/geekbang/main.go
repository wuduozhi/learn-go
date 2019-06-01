package main

import (
	"unsafe"
	"time"
	"fmt"
	"sync"
)

func main(){

	// for i:=0;i<20;i++{
	// 	test_counter()
	// 	test_counter_safe()
	// 	test_counter_safe_group()
	// }

	// test_service()
	// test_async_service()
	// test_select()
	test_once()
	
}

func test_counter(){
	counter := 0
	for i := 0;i<5000;i++{
		go func(){
			counter ++
		}()
	}

	time.Sleep(1*time.Second)
	fmt.Printf("counter = %d \t",counter)
}

func test_counter_safe(){
	var mut sync.Mutex
	counter := 0
	for i := 0;i<5000;i++{
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
		}()
	}
	time.Sleep(1*time.Second)
	fmt.Printf("counter = %d \t",counter)
}

func test_counter_safe_group(){
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0;i<5000;i++{
		wg.Add(1)
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("counter = %d \n",counter)
}

func service() string{
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func async_service() chan string {
	retCh := make(chan string,1)

	go func(){
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func other_task(){
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func test_service(){
	fmt.Println(service())
	other_task()
}

func test_async_service(){
	retCh := async_service()
	other_task()
	fmt.Println(<-retCh)
}

func test_select(){
	select {
	case ret := <- async_service():
		fmt.Println(ret)
	case <- time.After(time.Millisecond * 100):
		fmt.Println("time out")
	}
}

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

func get_single_obj() *Singleton{
	once.Do(func(){
		fmt.Println("creat obj")
		singleInstance = new(Singleton)
	})

	return singleInstance
}

func test_once(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go func(){
			obj := get_single_obj()
			fmt.Printf("%x \t",unsafe.Pointer(obj))
			wg.Done()
		}()
	}

	wg.Wait()
}