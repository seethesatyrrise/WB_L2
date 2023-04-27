package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	<-or(
		sig(7*time.Second),
		sig(6*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	mu := sync.Mutex{}
	unite := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутины для прослушки done-каналов (на каждый канал
	// своя горутина). как только пришел сигнал из какого-либо канала
	// останавливаем остальные горутины и отправляем в канал
	// unite номер канала, который прислал сигнал
	for i, ch := range channels {
		go func(ctx context.Context, i int, ch <-chan interface{}) {
			select {
			case <-ch:
				mu.Lock()
				if unite != nil {
					unite <- i
					close(unite)
					unite = nil
					cancel()
				}
				mu.Unlock()
			case <-ctx.Done():
			}
		}(ctx, i, ch)
	}

	ctx2, cancel2 := context.WithCancel(context.Background())

	// горутина, прослушивающая канал unite. если что-то в него пришло
	// (значит один из done-каналов сообщил о завершении), значит пора
	// объединить остальные каналы в single-канал.
	go func() {
		iCh := <-unite
		for i, ch := range channels {
			if i == iCh {
				continue
			}
			// для каждого канала, кроме отправившего сигнал, запускаем горутину
			// в которой отслеживаем сигнал с канала и в случае его получения,
			// отправляем сигнал в single-канал, после чего останавливаем все
			// остальные горутины
			go func(ctx context.Context, ch <-chan interface{}) {
				select {
				case <-ch:
					mu.Lock()
					if single != nil {
						single <- 1
						close(single)
						single = nil
						cancel2()
					}
					mu.Unlock()
				case <-ctx.Done():
				}
			}(ctx2, ch)
		}
	}()

	return single
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
		c <- time.Now()
	}()
	return c
}
