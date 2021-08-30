module sample.com/goGrep

go 1.13

replace sample.com/search => ./search

require (
	sample.com/file v0.0.0-00010101000000-000000000000
	sample.com/search v0.0.0-00010101000000-000000000000
)

replace sample.com/file => ./file
