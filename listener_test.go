package listener

import (
	"errors"
	"net"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		l, err := New(nil)
		if err != nil {
			t.Fatalf("New error: %s", err)
		}
		defer func() {
			if err := l.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l1.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l2.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l1.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
		defer func() {
			if err := l2.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
		}()
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
			if err := l3.Close(); err != nil {
				t.Errorf("Close error: %s", err)
			}
			t.Fatalf("New wants error but was nil")
		}
		t.Logf("expected error: %s", err)

		var noAvailablePortErr NoAvailablePortError
		if !errors.As(err, &noAvailablePortErr) {
			t.Fatalf("error wants NoAvailablePortError but was %T", err)
		}
		causes := noAvailablePortErr.Unwrap()
		if len(causes) != 2 {
			t.Fatalf("len(causes) wants 3 but was %d", len(causes))
		}
		cause1 := causes[0]
		var ne1 *net.OpError
		if !errors.As(cause1, &ne1) {
			t.Fatalf("cause1 wants net.OpError but was %T", errors.Unwrap(cause1))
		}
		if ne1.Addr.String() != "127.0.0.1:9000" {
			t.Errorf("Addr wants 127.0.0.1:9000 but was %s", ne1.Addr)
		}
		cause2 := causes[1]
		var ne2 *net.OpError
		if !errors.As(cause2, &ne2) {
			t.Fatalf("cause1 wants net.OpError but was %T", errors.Unwrap(cause2))
		}
		if ne2.Addr.String() != "127.0.0.1:9001" {
			t.Errorf("Addr wants 127.0.0.1:9001 but was %s", ne1.Addr)
		}
	})
}
