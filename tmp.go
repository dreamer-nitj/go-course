package main

func main() {
	go func() {
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1*time.Second) {
			log.Printf("goroutine count: %d", goroutines.count)
		}
	}

	var blockForever chan struct{}
	for i := 0; i < 10; i++ {
		go func() { <-blockForever }()
		time.Sleep(500*time.Millisecond)
	}
}
