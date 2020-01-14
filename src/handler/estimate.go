/* vim: set ts=4 sts=2 sw=4 tw=0 fenc=utf-8 : */
package estimate

import (
  "bytes"
  "context"
  //"encoding/json"

  "io"
  "net/http"
  "os"

  "github.com/signintech/gopdf"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type estimate struct {}

type Response events.APIGatewayProxyResponse

func NewHandler() *estimate {
  return &estimate{}
}
func Handler(ctx context.Context) (Response, error) {

  var buf bytes.Buffer
  //var err error
  var headers map[string]string

  pdf := gopdf.GoPdf{}
  pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

  pdf.AddPage()
  pdf.SetLineWidth(1)
  pdf.Oval(100, 200, 500, 500)

  buf.Read(pdf.GetBytesPdf())

  headers["Content-Type"] =	"application/pdf"

  resp := Response{
	  StatusCode:      200,
	  IsBase64Encoded: true,
	  Body:            buf.String(),
	  Headers:		   headers,
  }

  return resp, nil
}

func DownloadFile(filepath string, url string) error {
  // Get the data
  resp, err := http.Get(url)
  if err != nil {
	  return err
  }
  
  defer resp.Body.Close() 

  // Create the file
  out, err := os.Create(filepath)
  if err != nil {
	  return err
  }

  defer out.Close()
  // Write the body to file
  _, err = io.Copy(out, resp.Body)
  return err
}

func (p estimate) Main() {
	lambda.Start(Handler)
}

