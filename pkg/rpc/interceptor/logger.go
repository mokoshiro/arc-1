package interceptor

import (
	"context"
	"fmt"
	"time"
	"unicode/utf8"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	maxStrLen = 100
)

func ClientLoggingIntercepter(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		go func() {
			elapsed := time.Since(start).Seconds()
			code := codes.OK
			if err != nil {
				code = status.Code(err)
			}
			logger.Debug(
				"sent",
				zap.String("method", method),
				zap.String("req", shortenMessage(req)),
				zap.String("res", shortenMessage(reply)),
				zap.String("status", code.String()),
				zap.Float64("elapsed", elapsed),
				zap.Error(err),
			)
		}()
		return err
	}
}

func ServerLoggingIntercepter(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		res, err := handler(ctx, req)
		go func() {
			elapsed := time.Since(start).Seconds()
			code := codes.OK
			if err != nil {
				code = status.Code(err)
			}
			logger.Debug(
				"received",
				zap.String("method", info.FullMethod),
				zap.String("req", shortenMessage(req)),
				zap.String("res", shortenMessage(res)),
				zap.String("status", code.String()),
				zap.Float64("elapsed", elapsed),
				zap.Error(err),
			)
		}()
		return res, err
	}
}

func StreamClientLoggingIntercepter(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		go logger.Debug(
			"sent",
			zap.String("method", method),
		)
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func StreamServerLoggingIntercepter(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		go logger.Debug(
			"sent",
			zap.String("method", info.FullMethod),
		)
		return handler(srv, ss)
	}

}

func shortenMessage(obj interface{}) string {
	str := fmt.Sprint(obj)
	if utf8.RuneCountInString(str) > maxStrLen {
		str = string([]rune(str)[:maxStrLen]) + "..."
	}
	return str
}
