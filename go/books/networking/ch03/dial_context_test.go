package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)
	//cancel if not connected in 5 seconds
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()
	var d net.Dialer //DialContext is a method on a dialer
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		//force the error
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}
	//connecting to non routable address
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		conn.Close()
		t.Fatal("connection did not timeout")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not a timeout : %v", err)
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded ;actual : %v", ctx.Err())
	}
}
