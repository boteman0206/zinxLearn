package test

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func Test1() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", "127.0.0.1", 8081)
	fmt.Println("hello zinx learn..")

	v1, err := uuid.NewV1()
	fmt.Println(v1, err)
}
