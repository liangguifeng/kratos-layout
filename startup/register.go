package startup

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	google_grpc "google.golang.org/grpc"
)

// RegisterHTTPServer 此处注册http接口
func RegisterHTTPServer(httpSrv *http.Server) error {
	return nil
}

// RegisterGRPCServer 此处注册grpc接口
func RegisterGRPCServer(grpcSrv *grpc.Server) error {
	return nil
}

// RegisterInterceptor 注册拦截器
func RegisterInterceptor() []google_grpc.UnaryServerInterceptor {
	arr := make([]google_grpc.UnaryServerInterceptor, 0)
	return arr
}
