module github.com/celicoo/ttest

go 1.19

require (
	github.com/celicoo/ttest/internal/function v0.0.0
	github.com/celicoo/ttest/internal/slice v0.0.0
	github.com/celicoo/ttest/internal/tabwriter v0.0.0
)

replace (
	github.com/celicoo/ttest/internal/function => ./internal/function
	github.com/celicoo/ttest/internal/slice => ./internal/slice
	github.com/celicoo/ttest/internal/tabwriter => ./internal/tabwriter
)
