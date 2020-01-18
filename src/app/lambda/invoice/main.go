/* vim: set ts=4 sts=2 sw=4 tw=0 fenc=utf-8 : */
package main

import (
  "bytes"
  "context"
  //"encoding/json"
  "encoding/base64"

  "github.com/signintech/gopdf"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {

  var buf bytes
  //var err error

  pdf := gopdf.GoPdf{}
  pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

  pdf.AddPage()
  pdf.SetLineWidth(1)
  pdf.Oval(100, 200, 500, 500)

  pdf.SetLineWidth(2)
  pdf.SetLineType("dashed")
  pdf.Line(10, 30, 585, 30)


  resp := Response{
	  StatusCode:      200,
	  IsBase64Encoded: true,
	  Body:            base64.StdEncoding.EncodeToString(pdf.GetBytePdf()),
	  Headers: map[string]string {
		  "Content-Type":	"application/pdf",
	  },
  }

  return resp, nil
}

func main() {
	lambda.Start(Handler)
}

