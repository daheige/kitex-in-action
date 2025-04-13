package middleware

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"go.uber.org/zap"

	"kitex-example/internal/infras/logger"
)

// AccessLog 记录访问日志的拦截器
func AccessLog(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}, res interface{}) error {
		var requestID string
		if xRequestID := ctx.Value(logger.XRequestID.String()); xRequestID == nil {
			requestID = logger.Uuid()
		} else {
			requestID, _ = xRequestID.(string)
		}
		ctx = context.WithValue(ctx, logger.XRequestID, requestID)

		// 捕获recover
		defer func() {
			if r := recover(); r != nil {
				logger.Info(ctx, "exec panic", map[string]interface{}{
					"reply":       res,
					"trace_error": fmt.Sprintf("%v", r),
					"full_stack":  string(debug.Stack()),
				})
			}
		}()

		// 记录请求开始时间
		start := time.Now()

		// 提取请求的元数据
		ri := rpcinfo.GetRPCInfo(ctx)
		if from := ri.From(); from != nil {
			ctx = context.WithValue(ctx, logger.ReqClientIP, from.Address().String())
		}

		info := ri.Invocation()
		serviceName := info.ServiceName()
		if serviceName == "" {
			serviceName = ri.To().ServiceName()
		}
		methodName := info.MethodName()
		if methodName == "" {
			methodName = ri.To().Method()
		}

		realReq, _ := req.(utils.KitexArgs)
		logger.Info(ctx, "exec begin",
			zap.String("service_name", serviceName),
			zap.String("method_name", methodName),
			zap.Any("request", realReq),
		)

		err := next(ctx, req, res)
		execTime := time.Since(start).Seconds() // 计算请求耗时
		if err != nil {
			logger.Info(ctx, "rpc call fail", "trace_error", err.Error())
			return err
		}

		logger.Info(ctx, "exec end",
			zap.String("service_name", serviceName),
			zap.String("method_name", methodName),
			zap.String("exec_time", fmt.Sprintf("%.4f", execTime)),
		)

		return nil
	}
}
