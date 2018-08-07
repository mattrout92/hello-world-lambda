package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	optimus "github.com/pjebs/optimus-go"
)

var OptimusDecode = optimus.New(1322293361, 1216882833, 2078473863)

func main() {
	/*r := mux.NewRouter()

	r.Methods("GET").Path("/hello").HandlerFunc(hello)

	apr := gorillamux.New(r)

	lambda.Start(lambdaProxy(apr))*/

	fmt.Println(Encode(1))

}

func Encode(v int) int {
	return int(OptimusDecode.Encode(uint64(v)))
}

func lambdaProxy(apr *gorillamux.GorillaMuxAdapter) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		path := req.Path
		ip := req.Headers["X-Forwarded-For"]
		fmt.Printf("\n########\nrequest to path: %v [IP: %v]", path, ip)

		return apr.Proxy(req)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}
