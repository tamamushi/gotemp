module github.com/tamamushi/gotemp

go 1.13

require (
	github.com/apex/gateway v1.1.1 // indirect
	github.com/aws/aws-lambda-go v1.13.3
	github.com/fnproject/fdk-go v0.0.1
	github.com/gin-gonic/gin v1.5.0 // indirect
	github.com/pkg/errors v0.9.0 // indirect
	github.com/signintech/gopdf v0.9.5 // indirect
	github.com/tencentyun/scf-go-lib v0.0.0-20190817080819-4a2819cda320
	gopkg.in/yaml.v2 v2.2.7 // indirect
	internal.pkg/gotemp v0.0.0
)

replace internal.pkg/gotemp => ./src
