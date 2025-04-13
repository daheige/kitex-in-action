package middleware

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"go.uber.org/zap"

	"kitex-example/internal/infras/logger"
)

// AccessLog 记录访问日志的拦截器
func AccessLog(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}, res interface{}) error {
		var requestID string
		if logID := ctx.Value(logger.XRequestID.String()); logID == nil {
			requestID = logger.Uuid()
		} else {
			requestID, _ = logID.(string)
		}
		ctx = context.WithValue(ctx, logger.XRequestID, requestID)

		// 捕获recover
		defer func() {
			if r := recover(); r != nil {
				// the error format defined by grpc must be used here to return code, desc
				logger.Info(ctx, "exec panic", map[string]interface{}{
					"reply":       res,
					"trace_error": fmt.Sprintf("%v", r),
					"full_stack":  string(debug.Stack()),
				})
			}
		}()

		// 记录请求开始时间
		start := time.Now()
		// 获取请求的service info信息
		// 提取请求元数据
		ri := rpcinfo.GetRPCInfo(ctx)
		info := ri.Invocation()
		if from := ri.From(); from != nil {
			ctx = context.WithValue(ctx, logger.ReqClientIP, from.Address().String())
		}

		serviceName := info.ServiceName()
		if serviceName == "" {
			serviceName = ri.To().ServiceName()
		}

		methodName := info.MethodName()
		if methodName == "" {
			methodName = ri.To().Method()
		}

		logger.Info(ctx, "exec begin",
			zap.String("service_name", serviceName),
			zap.String("method_name", methodName),
			zap.Any("request", req),
		)

		err := next(ctx, req, res)
		cost := time.Since(start) // 计算请求耗时
		if err != nil {
			logger.Info(ctx, "rpc call fail", "trace_error", err.Error())
			return err
		}

		logger.Info(ctx, "exec end",
			zap.String("service_name", serviceName),
			zap.String("method_name", methodName),
			zap.Int64("exec_time", cost.Milliseconds()),
		)

		return nil
	}
}
