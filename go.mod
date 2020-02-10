module github.com/pendo-io/aws-sdk-go

require (
	github.com/aws/aws-sdk-go v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/tools v0.0.0-20200210175345-9fbd0ccf67bf
)

go 1.12

replace github.com/aws/aws-sdk-go => github.com/pendo-io/aws-sdk-go v1.99.0
