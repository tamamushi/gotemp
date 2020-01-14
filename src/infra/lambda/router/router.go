/* vim:set ts=4 sts=2 sw=4 tw=0 fenc=utf-8: */

package router

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "github.com/gin-gonic/gin"
)

type Handlers func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type Router struct {
	route	map[string]Handlers
}

func NewRouter() *Router {
	self	:= &Router{}
	self.route	= make(map[string]Handlers)
	return self 
}

func (p *Router) Get(f Handlers) {
	p.route["GET"] = f
}

func (p *Router) GetWithPathParam(f Handlers) {
	p.route["GETWithPathParam"] = f
}

func (p *Router) Post(f Handlers) {
	p.route["POST"] = f
}

func (p *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var key string

	switch req.HTTPMethod {
		case "GET":
			if len(req.PathParameters) > 0 {
				key	= "GETWithPathParam"
				fmt.Println("Request HTTP Method: GET with PathParam")
			} else {
				key	= "GET"
				fmt.Println("Request HTTP Method: GET")
			}
		case "POST": 
			fmt.Println("Request HTTP Method: POST")
	}
	fmt.Printf("req : %# v", pretty.Formatter(req))

	res, err := p.route[key](req); 
	if err != nil { return errorResponse(err) }
	return res, nil
}

func errorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, nil
}

