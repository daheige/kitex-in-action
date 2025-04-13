package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/utils"
)

// ValidatorLegacy validate interface
// 设计灵感来源于grpc-ecosystem/go-grpc-middleware
// 代码位置:https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/validator/validator.go#L26
type ValidatorLegacy interface {
	Validate() error
}

// 由于kitex将请求参数放在了Req字段中，源码实现
// type HelloArgs struct {
// 	Req *pb.HelloRequest
// }
// 因此这里需要类型断言判断是否实现了utils.KitexArgs接口，如果实现了就获取请求的真正参数req
func validate(request interface{}) error {
	if args, ok := request.(utils.KitexArgs); ok {
		req := args.GetFirstArgument() // 真正的请求结构体
		switch v := req.(type) {
		case ValidatorLegacy:
			if err := v.Validate(); err != nil {
				return fmt.Errorf("validator legacy validation failed: %s", err.Error())
			}
		default:
		}
	}

	return nil
}

// Validator 验证请求参数req是否合法
func Validator(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}, res interface{}) error {
		if err := validate(req); err != nil {
			return err
		}

		return next(ctx, req, res)
	}
}
