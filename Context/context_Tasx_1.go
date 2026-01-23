package ContextContext

import (
	"context"
	"fmt"
	"time"
)

// Задача: реализовать 3 контекста с тремя уровнями градации, проэсперементировать с отключением каждого из них
// Проверены варианты закрытия разного контекста, в том числе и главного, как и ожидалось в результате закрывались группы горутин в зависимости от градации контекста

func MainTasc1() {

	parentContext, parentCancel := context.WithCancel(context.Background())
	midleContext, midleCancel := context.WithCancel(parentContext)
	childContext, childCancel := context.WithCancel(midleContext)

	go hoo(childContext)
	go foo(parentContext)
	go boo(midleContext)

	time.Sleep(500 * time.Millisecond)
	childCancel()

	time.Sleep(500 * time.Millisecond)
	midleCancel()

	time.Sleep(500 * time.Millisecond)
	parentCancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Final Main")
}

func hoo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Hoo the compited :(")
			return
		default:
			fmt.Println("Hoo не завершилась :)")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo the compited :(")
			return
		default:
			fmt.Println("Foo не завершилась :)")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo the compited :(")
			return
		default:
			fmt.Println("Boo не завершилась :)")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
