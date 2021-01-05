package design

import (
	shared "github.com/sm43/calc/design"
	. "goa.design/goa/v3/dsl"
)

var _ = API("v1", func() {
	Title("Calculator Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("v1", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
		Services(
			"calc",
			"status",
		)
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("v1Add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/v1/add/{a}/{b}")
		})
	})

	Files("/v1/docs", "./v1/gen/http/openapi.json")
})

// To get documentation of APIs which are common in all versions
var SharedService = shared.SharedService
