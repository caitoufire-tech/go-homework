package homework02

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func TaskOne(num *int) {
	*num += 10
}

func TaskTwo(numbers *[]int) {
	for i := range *numbers {
		(*numbers)[i] *= 2
	}
}

func TaskThree() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				println("go one:", i)
				// time.Sleep(time.Second)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				println("go two:", i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func TaskFour(nums []int) {
	wg := sync.WaitGroup{}
	for i, num := range nums {
		wg.Add(1)
		go func(num int) {
			now := time.Now()
			task(num)
			fmt.Println("task:", i, time.Since(now))
			wg.Done()
		}(num)
	}
	wg.Wait()
}

func task(num int) {
	result := 0
	for i := 0; i < num; i++ {
		result += i
	}
}

func TaskFive() {
	circle := &Circle{Radius: 5}
	fmt.Println("Circle Area:", circle.Area())
	fmt.Println("Circle Perimeter:", circle.Perimeter())
	rectangle := &Rectangle{Width: 4, Height: 6}
	fmt.Println("Rectangle Area:", rectangle.Area())
	fmt.Println("Rectangle Perimeter:", rectangle.Perimeter())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func TaskSix() {
	employee := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}
	employee.PrintInfo()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func TaskSeven() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("Received:", num)
			default:
				time.Sleep(50 * time.Millisecond)
			}

		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("done")
}

func TaskEight() {
	bufferCh := make(chan int, 30)

	go func() {
		for i := 1; i <= 100; i++ {
			bufferCh <- i
			fmt.Println("send:", i)
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Received:", <-bufferCh)
		}
	}()

	time.Sleep(time.Second * 60)
}

func TaskNine() {
	ct := Counter{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				ct.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count:", ct.count)
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func TaskTen() {
	ca := CounterAtomic{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				ca.count.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Atomic Count:", ca.count.Load())
}

type CounterAtomic struct {
	count atomic.Int64
}
