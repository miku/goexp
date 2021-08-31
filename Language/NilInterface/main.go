// go run main.go
//
// 2019/08/26 17:19:28 (*main.MyErr)(nil): some error message
// exit status 1
//
package main

import (
	"log"
)

type MyErr struct{}

func (e *MyErr) Error() string {
	return "some error message"
}

func mayFail(f float32) *MyErr {
	return nil
}

func main() {
	var err error      // interface
	err = mayFail(0.4) // err may be not nil, even though nil is the zero value for a pointer
	if err != nil {
		log.Fatalf("%#v: %s", err, err.Error()) // (*main.MyErr)(nil): some error message
	}
}
