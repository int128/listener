// Package listener provides utility for allocating a net.Listener from address candidates.
package listener

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"golang.org/x/xerrors"
)

// Listener wraps a net.Listener and provides its URL.
type Listener struct {
	l net.Listener

	// URL to the listener.
	// This is always "http://localhost:PORT" regardless of the listening address.
	URL *url.URL
}

func (l *Listener) Accept() (net.Conn, error) {
	return l.l.Accept()
}

func (l *Listener) Close() error {
	return l.l.Close()
}

func (l *Listener) Addr() net.Addr {
	return l.l.Addr()
}

// New starts a Listener on one of the addresses.
// Caller should close the listener finally.
//
// If nil or an empty slice is given, it will allocate a free port at "127.0.0.1".
// If multiple address are given, it will try the addresses in order.
func New(addressCandidates []string) (*Listener, error) {
	if len(addressCandidates) == 0 {
		return NewOn("")
	}
	var errs []string
	for _, address := range addressCandidates {
		l, err := NewOn(address)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		return l, nil
	}
	return nil, xerrors.Errorf("no available port (%s)", strings.Join(errs, ", "))
}

// NewOn starts a Listener on the address.
// If an empty string is given, it defaults to "127.0.0.1:0".
// Caller should close the listener finally.
func NewOn(address string) (*Listener, error) {
	if address == "" {
		address = "127.0.0.1:0"
	}
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, xerrors.Errorf("could not listen: %w", err)
	}
	addr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return nil, xerrors.Errorf("internal error: got a unknown type of listener %T", l.Addr())
	}
	return &Listener{
		l:   l,
		URL: &url.URL{Host: fmt.Sprintf("localhost:%d", addr.Port), Scheme: "http"},
	}, nil
}
