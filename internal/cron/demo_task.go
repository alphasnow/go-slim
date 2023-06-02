package cron

import "fmt"

type DemoTask struct {
}

func (d *DemoTask) Handle() {
	fmt.Println("Cron Demo Task")
}
