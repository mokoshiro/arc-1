package dialer

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
)

// DialContext is wrapper function for validate destination address,
// and obscure this logic
func DialContext(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if err := validateDestinationAddr(addr); err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	return grpc.DialContext(ctx, addr, opts...)
}

func validateDestinationAddr(addr string) error {
	if addr == "" {
		return errors.New("destination address is empty")
	}
	if ipAndPortStr := strings.Split(addr, ":"); len(ipAndPortStr) != 2 || len(ipAndPortStr[1]) == 0 {
		return fmt.Errorf("invalid address, got=%s", addr)
	}
	_, err := net.ResolveTCPAddr("tcp", addr)
	return err
}
