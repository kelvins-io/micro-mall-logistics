package startup

import (
	"context"
	"net/http"

	"gitee.com/cristiane/micro-mall-logistics/http_server"
	"gitee.com/cristiane/micro-mall-logistics/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-logistics/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RegisterGRPCServer 此处注册pb的Server
func RegisterGRPCServer(grpcServer *grpc.Server) error {
	order_business.RegisterOrderBusinessServiceServer(grpcServer, server.NewOrderServer())
	return nil
}

// RegisterGateway 此处注册pb的Gateway
func RegisterGateway(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error {
	if err := order_business.RegisterOrderBusinessServiceHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}
	return nil
}

// RegisterHttpRoute 此处注册http接口
func RegisterHttpRoute(serverMux *http.ServeMux) error {
	serverMux.HandleFunc("/swagger/", http_server.SwaggerHandler)
	return nil
}
