package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/openzipkin/zipkin-go"
	mw "github.com/openzipkin/zipkin-go/middleware/http"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	// 第一步: 开一个全局变量
	var nativeTracer *zipkin.Tracer

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

	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/app", nil)
	name := "middle-client"
	client, _ := mw.NewClient(nativeTracer)
	resp, err := client.DoWithAppSpan(req, name)
	if err != nil {
		log.Fatalf("err:%s", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("resp:%s", body)

}
