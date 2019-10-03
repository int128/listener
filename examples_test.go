package listener_test

import (
	"fmt"

	"github.com/int128/listener"
)

// ExampleNew_Candidates allocates a port at 18000 or 28000.
func ExampleNew_Candidates() {
	l, err := listener.New([]string{"127.0.0.1:18000", "127.0.0.1:28000"})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Printf("Open %s", l.URL)

	// Output:
	// Open http://localhost:18000
}

// ExampleNew_FreePort allocates a free port.
func ExampleNew_FreePort() {
	l, err := listener.New(nil)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Printf("Hostname=%s", l.URL.Hostname())

	// Output:
	// Hostname=localhost
}
