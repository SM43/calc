package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
		Services(
			"status",
		)
	})
})

// SharedService Services which would be shared in multiple version would be defined at root
// and imported in the version only to add documentation
// APIs which are not user facing
var SharedService = Service("status", func() {
	Description("The status service provides status of the service.")

	Method("status", func() {

		Result(String)

		HTTP(func() {
			GET("/")
		})
	})

	Files("/docs", "./gen/http/openapi.json")
})

// design (shared design code, e.g. shared types)
//    types.go
//    ...
// service1
//    design (imports top level design)
//       api.go
//       ...
//    gen
//    service1.go
//    ...
// service2
//    design (imports top level design)
//       api.go
//       ...
//    gen
//    service2.go
//    ...s
