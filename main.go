package main

import (
	"yo"

	appdemo "appdemo/app"
)

func main() {
	appdemo.Init() // keep in `main()`, dont move to `init()`
	doListenAndServe := yo.Init(staticFs_Yo, staticFs_App)
	appdemo.OnBeforeListenAndServe()
	doListenAndServe()
}
