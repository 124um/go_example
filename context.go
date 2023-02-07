package main

import (
	"context"
	"fmt"
	"time"
)

// context.Background() - используется только на самом високом уровне (под маин)
// context.TODO() - когда не уверен какой контекст использовать
// context.WithValue - использовать как можно реже и передавать
//  только необязательние параметри потому что єто не явная передача параметров
// ctx передается всегда первим аргументом в функцию
// не передавать контекст nil - если не уверен использовать ctx.TODO
// только та функция которая порождает контекст - может отменить єтот контекст

func mainCtx() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	ctx = context.WithValue(ctx, "ID", 223323)
	// go func() {
	// 	time.Sleep(time.Millisecond * 100)
	// 	cancel()
	// }()
	parse(ctx)
}

func parse(ctx context.Context) {
	id := ctx.Value("ID")
	fmt.Println(id)

	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing complete")
			return
		case <-ctx.Done():
			fmt.Println("dedline exeded")
			return
		}
	}
}
