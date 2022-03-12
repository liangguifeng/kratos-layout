package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/liangguifeng/kratos-app"
	"github.com/liangguifeng/kratos-app/app"
	"kratos-layout/startup"
)

// 此处定义当前服务相关配置
const (
	APP_NAME      = "kratos-layout"
	APP_HTTP_POTR = "50000"
	APP_GRPC_POTR = "60000"
)

func main() {
	runer, err := app.NewRunner(&kratos.Application{
		Name:       APP_NAME,
		HTTPPort:   APP_HTTP_POTR,
		GRPCPort:   APP_GRPC_POTR,
		Type:       kratos.APP_TYPE_GRPC,
		LoadConfig: startup.LoadConfig,
		SetupVars:  startup.SetupVars,
		RegisterCallback: map[kratos.CallbackPos]func() error{
			kratos.POS_LOAD_CONFIG: startup.LoadConfigCallback,
			kratos.POS_SETUP_VARS:  startup.SetupVarsCallback,
			kratos.POS_NEW_RUNNER:  startup.RunNewRunnerCallback,
		},
	})
	if err != nil {
		log.Fatalf("app.NewRunner err: %v", err)
	}

	err = runer.ListenGRPCServer(&kratos.GRPCApplication{
		RegisterGRPCServer:      startup.RegisterGRPCServer,
		RegisterHttpRoute:       startup.RegisterHTTPServer,
		UnaryServerInterceptors: startup.RegisterInterceptor(),
	})
	if err != nil {
		log.Fatalf("runer.ListenGRPCServer err: %v", err)
	}
}
