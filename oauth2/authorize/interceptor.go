package authorize

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Error(codes.InvalidArgument, "missing metadata")
	errWrongFormat     = status.Error(codes.InvalidArgument, "wrong format")
)

func StreamAuthCheckInterceptor() grpc.ServerOption {
	return grpc.StreamInterceptor(streamAuthInterceptor)
}

func UnaryAuthCheckInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(unaryAuthInterceptor)
}

func streamAuthInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}
	au := md["authorization"]

	if len(au) < 1 {
		return errWrongFormat
	}

	err := authorize(au[0])
	if err != nil {
		return err
	}

	return handler(srv, ss)
}

func unaryAuthInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	au := md["authorization"]
	if len(au) < 1 {
		return nil, errWrongFormat
	}

	err := authorize(au[0])
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
