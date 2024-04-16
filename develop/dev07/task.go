package main

import (
	"fmt"
	"time"
)

// Функция or принимает один или более каналов и объединяет их в один канал.
// Когда один из исходных каналов закрывается, результирующий канал также закрывается.
func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})

	go func() {
		defer close(result)
		for _, ch := range channels {
			go func(ch <-chan interface{}) {
				for val := range ch {
					select {
					case <-result: // Если результирующий канал уже закрыт, просто выходим из цикла
						return
					case result <- val: // Пишем в результирующий канал
					}
				}
			}(ch)
		}
	}()

	return result
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
