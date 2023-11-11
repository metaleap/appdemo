package appdemo

import (
	"yo"

	. "yo/srv"
)

func init() {
	yo.AppPkgPath = appdemoPkg.PkgPath()
	AppApiUrlPrefix = "_/"
	AppSideStaticRePathFor = func(requestPath string) string {
		return "__static/appdemo.html"
	}
}

func Init() {
	// yodb.Ensure[T] calls here
}

func OnBeforeListenAndServe() {
	// yodb.Upsert[yojobs.JobDef] calls here
}
