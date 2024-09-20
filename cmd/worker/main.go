package main

import (
	"github.com/troptropcontent/what_the_tide/internal/lib/worker"
)

func main() {
	mgr := worker.NewBackGroundWorker()
	mgr.Run()
}
