package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "hello/kitex_gen/pb/greeter/gateway"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// 这个grpc微服务地址，一般来说是一个远程的ip:port，可以根据实际情况更改
	// gRPCAddress := fmt.Sprintf("0.0.0.0:%d",conf.GrpcPort)
	gRPCAddress := fmt.Sprintf("0.0.0.0:%d", 8890)

	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, gRPCAddress, opts)
	if err != nil {
		log.Fatal(ctx, "failed to register grpc endpoint", map[string]interface{}{
			"trace_error": err.Error(),
		})
	}

	router := gin.New()
	initRouter(router, mux)

	addr := fmt.Sprintf("0.0.0.0:%d", 8091)
	// 服务server设置
	server := &http.Server{
		Handler:           router,
		Addr:              addr,
		IdleTimeout:       20 * time.Second, // tcp idle time
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// 在独立携程中运行
	log.Println("server run on: ", addr)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Println("server start panic,err:", e)
			}
		}()

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(context.Background(), "server close error", map[string]interface{}{
					"trace_error": err.Error(),
				})

				log.Println("server close error:", err)
				return
			}

			log.Println("server will exit...")
		}
	}()

	// server平滑重启
	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// receive signal to exit main goroutine.
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-ch

	log.Println("exit signal: ", sig.String())
	// Create a deadline to wait for.
	ctx, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)

		_ = server.Shutdown(ctx)
	}()

	<-done
	<-ctx.Done()

	log.Println("server shutting down")
}

func initRouter(router *gin.Engine, mux *runtime.ServeMux) {
	// gateway http proxy
	// 这里将proto文件中的路由地址进行路由注册
	// 访问方式：localhost:8091/v1/hello/daheige
	router.Any("/v1/*any", gin.WrapH(mux))

	// 添加自定义路由地址
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
