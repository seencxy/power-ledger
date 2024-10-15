package api

import (
	"PowerLedgerGo/infrastructure/config"
	"PowerLedgerGo/infrastructure/monitor"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

// NewsHandler service handler
type NewsHandler struct {
	AppConfig         *config.AppConfig          `inject:""`
	RouterHandler     *RouterHandler             `inject:""`
	PrometheusMonitor *monitor.PrometheusMonitor `inject:""`
}

// Run start services
func (s *NewsHandler) Run() {
	addr := fmt.Sprintf("0.0.0.0:%d", s.AppConfig.Port)
	// log.Println("AppConfig: ", s.AppConfig)
	log.Printf("Server running on:%s", addr)

	// register mux router
	router := s.RouterHandler.Router()

	// service prometheus monitor
	s.PrometheusMonitor.Monitor()

	// create http services
	server := &http.Server{
		// Handler: http.TimeoutHandler(router, time.Second*6, `{code:503,"message":"services timeout"}`),
		Handler:      router,
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go s.RouterHandler.MainHandler.MainService.StartWatchTradeExecuted()
	go s.RouterHandler.MainHandler.MainService.StartWatchTradePayment()
	go s.RouterHandler.MainHandler.MainService.StartWatchUserWithdraw()
	go func() {
		// 创建每分钟一个计时器
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := s.RouterHandler.MainHandler.MainService.CheckTraceExpired(); err != nil {
					break
				}
			}
		}
	}()

	// run http services in goroutine
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("exec panic", map[string]interface{}{
					"error":       err,
					"error_trace": string(debug.Stack()),
				})
			}

		}()

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("services listen error:", err)
				return
			}

			log.Println("services will exit...")
		}
	}()

	// graceful exit
	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// receive signal to exit mains goroutine
	// window signal
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	// linux signal
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2, os.Interrupt, syscall.SIGHUP)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Block until we receive our signal.
	<-ch

	// Create s deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), s.AppConfig.GracefulWait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in s goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	go func() {
		err := server.Shutdown(ctx)
		if err != nil {
			os.Exit(http.StatusOK)
		}
	}()
	<-ctx.Done()

	log.Println("services shutdown success")
}
