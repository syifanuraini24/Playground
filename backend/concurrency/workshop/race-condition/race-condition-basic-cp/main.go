package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	// TODO: answer here
	input := make(chan int)

	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//kirim 1 ke channel
			// TODO: answer here
			wg.Done()
			input <- 1
		}()
	}
	//mengubah nilai count menggunakan data dari channel

	// TODO: answer here
	for i := 0; i < 1000; i++ {
		count += <-input
	}
	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count
}
