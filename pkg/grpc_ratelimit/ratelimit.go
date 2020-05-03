package grpc_ratelimit

import (
	"context"
	"net"
	"time"

	"github.com/karlseguin/ccache"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type GRPCLimiter struct {
	rateLimit rate.Limit
	burst     int

	cache *ccache.Cache
}

func NewGRPCLimiter(rateLimit rate.Limit, burst int) *GRPCLimiter {
	return &GRPCLimiter{
		rateLimit: rateLimit,
		burst:     burst,
		cache:     ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100)),
	}
}

func (l *GRPCLimiter) getVisitor(ip string) (*rate.Limiter, error) {
	limiter, err := l.cache.Fetch(ip, time.Minute*10, func() (interface{}, error) {
		return rate.NewLimiter(l.rateLimit, l.burst), nil
	})
	if err != nil {
		return nil, err
	}

	return limiter.Value().(*rate.Limiter), nil
}

// UnaryServerInterceptor returns a new unary server interceptors that performs request rate limiting.
func UnaryServerInterceptor(limiter *GRPCLimiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		p, found := peer.FromContext(ctx)
		if found {
			var srcIP string
			switch addr := p.Addr.(type) {
			case *net.UDPAddr:
				srcIP = addr.IP.String()
			case *net.TCPAddr:
				srcIP = addr.IP.String()
			}
			l, err := limiter.getVisitor(srcIP)
			if err == nil {
				if !l.Allow() {
					return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
				}
			}
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new stream server interceptor that performs rate limiting on the request.
func StreamServerInterceptor(limiter *GRPCLimiter) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		p, found := peer.FromContext(stream.Context())
		if found {
			var srcIP string
			switch addr := p.Addr.(type) {
			case *net.UDPAddr:
				srcIP = addr.IP.String()
			case *net.TCPAddr:
				srcIP = addr.IP.String()
			}
			l, err := limiter.getVisitor(srcIP)
			if err == nil {
				if !l.Allow() {
					return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
				}
			}
		}
		return handler(srv, stream)
	}
}
