module github.com/basebytes/elastic-go

require (
	github.com/basebytes/elastic-go/client v0.0.5
	github.com/basebytes/elastic-go/service v0.0.6
	github.com/basebytes/tools v0.0.2
)

replace (
	github.com/basebytes/elastic-go/client v0.0.5 => ./client
	github.com/basebytes/elastic-go/service v0.0.6 => ./service
)

go 1.16
