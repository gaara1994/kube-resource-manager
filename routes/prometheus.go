package routes

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

func Metrics()  {
	// 使用promauto创建一个自动注册的计数器metric
	requestCount := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_requests_total",
		Help: "Total number of requests processed by my application.",
	})
	// 模拟处理请求并增加计数器
	go func() {
		for {
			time.Sleep(1 * time.Second)
			requestCount.Inc()
		}
	}()
}
