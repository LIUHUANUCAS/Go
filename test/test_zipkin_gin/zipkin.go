package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

// 第一步: 开一个全局变量
var zkTracer opentracing.Tracer
var nativeTracer *zipkin.Tracer

func main() {
	// 第二步: 初始化 tracer
	{
		reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
		defer reporter.Close()
		endpoint, err := zipkin.NewEndpoint("main3", "localhost:80")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+vn", err)
		}
		nativeTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+vn", err)
		}
		zkTracer = zkOt.Wrap(nativeTracer)
		opentracing.SetGlobalTracer(zkTracer)
	}

	r := gin.Default()
	// 第三步: 添加一个 middleWare, 为每一个请求添加span
	r.Use(func(c *gin.Context) {
		span, ctx := nativeTracer.StartSpanFromContext(c.Request.Context(), c.FullPath()+"-http")
		// span, ctx := nativeTracer.StartSpanFromContext(c, c.FullPath()+"-http")
		defer span.Finish()
		c.Request = c.Request.WithContext(ctx) // inherit from ctx
		c.Next()
	})
	r.GET("/",
		func(c *gin.Context) {
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.GET("/app",
		func(c *gin.Context) {
			doapp(c)
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.Run(":8080")
}

//func doapp(c context.Context) {
func doapp(c *gin.Context) {
	span, pc := nativeTracer.StartSpanFromContext(c.Request.Context(), "doapp")
	defer span.Finish()
	time.Sleep(time.Duration(rand.Intn(100)+1) * time.Millisecond)
	do2(pc)
}

//func do2(c *gin.Context) {
func do2(c context.Context) {
	span, _ := nativeTracer.StartSpanFromContext(c, "do2")
	defer span.Finish()
	x := rand.Intn(50) + 10
	time.Sleep(time.Duration(x) * time.Millisecond)

}
