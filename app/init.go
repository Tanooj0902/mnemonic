package app

import (
	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		jsonDataBindAsActionArguments,
		// revel.SessionFilter,           // Restore and write the session cookie.
		// revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter, // Restore kept validation errors and save new ones from cookie.
		// revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,            // Add some security based headers
		revel.InterceptorFilter, // Run interceptors around the action.
		revel.CompressFilter,    // Compress the result.
		revel.ActionInvoker,     // Invoke the action.
	}

	// register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

type notJsonReqestErrRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// This will work only for the single level keys in the JSON, now this is the limitaion we can not directly overcome.
// Instead we will have to change the type of url.Values and make modifications accordingly
// Can be done but for this test task I am ignoring it.
var jsonDataBindAsActionArguments = func(c *revel.Controller, fc []revel.Filter) {

	if c.Request.ContentType != "application/json" {
		fc[0](c, fc[1:])
	} else {
		var jsonData map[string]string
		err := c.Params.BindJSON(&jsonData)

		if err != nil {
			data := notJsonReqestErrRes{Code: 442, Message: "Not a valid Json request."}
			c.Response.Status = 442
			c.Result = c.RenderJSON(data)
			return
		}

		for k, v := range jsonData {
			c.Params.Values.Set(k, v)
		}

		fc[0](c, fc[1:])
	}
}
