package main

import (
	"context"
	"fmt"
	"github.com/catmullet/go-workers"
	"time"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {
	ctx := context.Background()
	rnr := workers.NewRunner(ctx, NewWorker(), 10).SetTimeout(15 * time.Second).Start()

	for i := 0; i < 100; i++ {
		rnr.Send(fmt.Sprintf("%d", i))
		time.Sleep(1 * time.Millisecond)
	}

	if err := rnr.Wait(); err != nil {
		fmt.Println(err)
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("finish")

}

type MyWorker struct {
}

func NewWorker() workers.Worker {
	return &MyWorker{}
}

func (wo *MyWorker) Work(in interface{}, out chan<- interface{}) error {
	fmt.Println(in)
	//if in.(string) == "50" {
	//	return &errorString{s: "errors " + in.(string)}
	//}
	return nil
}
