module broker

go 1.22.1

replace jsonHelpers => ../json-helpers

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.0
	jsonHelpers v0.0.0-00010101000000-000000000000
)
