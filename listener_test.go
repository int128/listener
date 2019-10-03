package listener

import (
	"fmt"
	"testing"
)

// ExampleNew_Candidates allocates a port at 18000 or 28000.
func ExampleNew_Candidates() {
	l, err := New([]string{"127.0.0.1:18000", "127.0.0.1:28000"})
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
	l, err := New(nil)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Printf("Hostname=%s", l.URL.Hostname())

	// Output:
	// Hostname=localhost
}

func TestNew(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		l, err := New(nil)
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l.Close()
		if l.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l.URL.Scheme)
		}
		if l.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l.URL.Hostname())
		}
		t.Logf("URL is %s", l.URL.String())
	})

	t.Run("Empty", func(t *testing.T) {
		l, err := New([]string{})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l.Close()
		if l.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l.URL.Scheme)
		}
		if l.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l.URL.Hostname())
		}
		t.Logf("URL is %s", l.URL.String())
	})

	t.Run("SingleAddress", func(t *testing.T) {
		l, err := New([]string{"localhost:9000"})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l.Close()
		if l.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l.URL.Scheme)
		}
		if l.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l.URL.Hostname())
		}
		if l.URL.Port() != "9000" {
			t.Errorf("Port wants 9000 but was %s", l.URL.Port())
		}
	})

	t.Run("MultipleAddressFallback", func(t *testing.T) {
		l1, err := New([]string{"localhost:9000"})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l1.Close()
		if l1.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l1.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l1.URL.Scheme)
		}
		if l1.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l1.URL.Hostname())
		}
		if l1.URL.Port() != "9000" {
			t.Errorf("Port wants 9000 but was %s", l1.URL.Port())
		}

		l2, err := New([]string{"localhost:9000", "localhost:9001"})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l2.Close()
		if l2.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l2.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l2.URL.Scheme)
		}
		if l2.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l2.URL.Hostname())
		}
		if l2.URL.Port() != "9001" {
			t.Errorf("Port wants 9001 but was %s", l2.URL.Port())
		}
	})

	t.Run("MultipleAddressFail", func(t *testing.T) {
		l1, err := New([]string{"localhost:9000"})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l1.Close()
		if l1.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l1.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l1.URL.Scheme)
		}
		if l1.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l1.URL.Hostname())
		}
		if l1.URL.Port() != "9000" {
			t.Errorf("Port wants 9000 but was %s", l1.URL.Port())
		}

		l2, err := New([]string{"localhost:9001"})
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer l2.Close()
		if l2.URL == nil {
			t.Errorf("URL wants a URL but was nil")
		}
		if l2.URL.Scheme != "http" {
			t.Errorf("Scheme wants http but was %s", l2.URL.Scheme)
		}
		if l2.URL.Hostname() != "localhost" {
			t.Errorf("Hostname wants localhost but was %s", l2.URL.Hostname())
		}
		if l2.URL.Port() != "9001" {
			t.Errorf("Port wants 9001 but was %s", l2.URL.Port())
		}

		l3, err := New([]string{"localhost:9000", "localhost:9001"})
		if err == nil {
			l3.Close()
			t.Fatalf("New wants error but was nil")
		}
		t.Logf("expected error: %s", err)
	})
}
