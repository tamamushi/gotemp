/* vim: set ts=4 sts=2 sw=4 tw=0 fenc=utf-8 : */
package rest

import (
  "encoding/base64"
  usecase "internal.pkg/gotemp/usecase"
  "github.com/gin-gonic/gin"
)

type EstimateHandler interface {
	Create(*gin.Context)
} 

type estimateHandler struct {
	estimateUseCase	usecase.EstimateUseCase
}

func NewEstimateHandler(bu usecase.EstimateUseCase) EstimateHandler {
	return &estimateHandler{
		estimateUseCase: bu,
	}
}

func (eh *estimateHandler) Create(c *gin.Context) {

	pdf, err := eh.estimateUseCase.CreatePdfByBytes()
	if err != nil {
		c.JSON(500, gin.H { "error": err.Error() })
		return
	}

	c.Data(200, "application/pdf", pdf)
}
