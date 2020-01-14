/* vim:set ts=4 sts=2 sw=4 tw=0 fenc=utf-8: */

package main

import (
  "log"
  "os"

  handler "internal.pkg/gotemp/handler/rest"
  usecase "internal.pkg/gotemp/usecase"

  "github.com/gin-gonic/gin"
  "github.com/apex/gateway"
)

var (
    Version  string
    Revision string
)

func main() {

	estimateUseCase	:= usecase.NewEstimateUseCase()
	estimateHandler	:= handler.NewEstimateHandler(estimateUseCase)

	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.POST("/estimate", estimateHandler.Create)
	r.GET("/estimate", estimateHandler.Create)

	addr := ":" + os.Getenv("PORT")
	log.Fatal(gateway.ListenAndServe(addr, r))
}

