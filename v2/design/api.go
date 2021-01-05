package design

import (
	shared "github.com/sm43/calc/design"
	. "goa.design/goa/v3/dsl"
)

var _ = API("v2", func() {
	Title("Calculator Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("v2", func() {
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

	Method("v2Add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/v2/add/{a}/{b}")
		})
	})

	Files("/v2/docs", "./v2/gen/http/openapi.json")
})

// To get documentation of APIs which are common
var SharedService = shared.SharedService
