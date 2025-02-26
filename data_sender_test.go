package godatasender

import (
	"fmt"
	"testing"
	"time"
)

func TestIntSender(t *testing.T) {
	send := func(datas []int) error {
		fmt.Println("send:", datas)
		return nil
	}
	opts := &Opts{
		Interval:  5 * time.Second,
		MaxLength: 100,
		BatchSize: 2,
	}

	sender := New(send, opts)
	go sender.Run()

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		sender.Enqueue(i, i+1000, i+2000)
	}

}
