module github.com/riweston/aztx/pkg/storage

go 1.21

require (
	github.com/riweston/aztx/pkg/errors v0.0.0
	github.com/riweston/aztx/pkg/types v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace (
	github.com/riweston/aztx/pkg/errors => ../errors
	github.com/riweston/aztx/pkg/types => ../types
)
