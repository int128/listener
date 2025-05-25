package listener_test

import (
	"fmt"

	"github.com/int128/listener"
)

// ExampleNew allocates a net.Listener at port 18000 or 28000.
func ExampleNew() {
	l, err := listener.New([]string{"127.0.0.1:18000", "127.0.0.1:28000"})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := l.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Open %s", l.URL)

	// Output:
	// Open http://localhost:18000
}
