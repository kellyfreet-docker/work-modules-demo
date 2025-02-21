module github.com/example/petcare-api

go 1.22.6

require (
	github.com/example/petcare-common/v2 v2.0.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/example/petcare-common/v2 => ../petcare-common
