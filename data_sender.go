package godatasender

import (
	"time"
)

type DataSender[T any] struct {
	Sending func([]T) error
	queue   chan T
	timeer  *time.Timer
	opts    *Opts
}

type Opts struct {
	Interval  time.Duration // The maximum interval for each data transmission
	MaxLength int           // The maximum length of the queue
	BatchSize int           // The maximum number of data to be sent in each transmission
}

func New[T any](sending func([]T) error, opts *Opts) *DataSender[T] {
	return &DataSender[T]{
		Sending: sending,
		queue:   make(chan T, opts.MaxLength),
		timeer:  time.NewTimer(opts.Interval),
		opts:    opts,
	}
}

// Enqueue adds data to the queue.
func (ds *DataSender[T]) Enqueue(datas ...T) {
	for _, data := range datas {
		ds.queue <- data
	}
}

// Run starts the data sender.
func (ds *DataSender[T]) Run() {
	datas := make([]T, 0, ds.opts.BatchSize)

	for {
		select {
		case data := <-ds.queue:
			datas = append(datas, data)
			if len(datas) >= ds.opts.BatchSize {
				ds.Sending(datas)
				datas = datas[:0]
				ds.timeer.Reset(ds.opts.Interval)
			}
		case <-ds.timeer.C:
			if len(datas) > 0 {
				ds.Sending(datas)
				datas = datas[:0]
			}
			ds.timeer.Reset(ds.opts.Interval)
		}
	}
}
