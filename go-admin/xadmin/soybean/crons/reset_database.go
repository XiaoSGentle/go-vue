package crons

import "fmt"

type AllCron struct {
}

func (receiver AllCron) SayHello(p1 string, p2 string) {
	fmt.Println(p1 + p2)
}
