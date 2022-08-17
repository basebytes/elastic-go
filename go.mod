module github.com/basebytes/elastic-go

go 1.16

require (
	github.com/basebytes/tools v0.0.1
	github.com/basebytes/elastic-go/client v0.0.2
	github.com/basebytes/elastic-go/service v0.0.2
	github.com/mitchellh/mapstructure v1.4.1
)

replace (
    github.com/basebytes/elastic-go/service v0.0.2 => ./service
)