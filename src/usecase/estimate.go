/* vim: set ts=4 sts=2 sw=4 tw=0 fenc=utf-8 : */
package usecase

import (

  "github.com/signintech/gopdf"
)

type EstimateUseCase interface {
	CreatePdfByBytes() ([]byte, error)
}

type estimateUseCase struct {}

func NewEstimateUseCase() EstimateUseCase {
	return &estimateUseCase{}
}

func (eu *estimateUseCase) CreatePdfByBytes() ([]byte, error) {

  pdf := gopdf.GoPdf{}
  pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

  pdf.AddPage()
  pdf.SetLineWidth(1)
  pdf.Oval(100, 200, 500, 500)

  return pdf.GetBytesPdf(), nil
}
