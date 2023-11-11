package appdemo

import (
	. "yo/srv"
	. "yo/util"
)

func init() {
	Apis(ApiMethods{
		"_/helloName": apiHelloName.Checks(
			Fails{Err: "ExpectedName", If: HelloNameName.Equal("")},
		),
	})
}

var apiHelloName = api(func(this *ApiCtx[struct {
	Name string
}, Return[string]]) {
	this.Ret.Result = "Hello, " + this.Args.Name
})
