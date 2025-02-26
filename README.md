# Overview
This code repository contains a Go - based implementation of a queue - based data sending service. 

The service is designed to send data in batches when either the queue reaches a certain length or a specific amount of time has passed since the last data send. 

The data type and the function for actually sending the data are left to the user's definition, providing high flexibility for various use - cases.


# Quick Start
## Install
```bash
go get github.com/lunuan/go-datasender
```

## Usage
import datasender package
```go
import (
	datasender "github.com/lunuan/go-datasender"
)
```

define your sending function
```go
// Note: The data type of the parameters of the sending function determines the data type of the sending queue.
sending := func(datas []int) error {
	fmt.Println("send:", datas)
	return nil
}
```

create a new DataSender instance and start it
```go
opts := &datasender.=Opts{
	Interval:  5 * time.Second,
	MaxLength: 1000,
	BatchSize: 100,
}

intSender := datasender.New(sending, opts)
go intSender.Start()
```

enqueue data
```go
intSender.Enqueue(data1, data2, data3, ...)
```

# License
This code is released under the MIT License. Please refer to the LICENSE file in the repository for details.
