package main

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

// OVERALL APPLICATION WIRING IN vuguSetup
func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	// CREATE A NEW ROUTER INSTANCE
	router := vgrouter.New(eventEnv)

	// MAKE OUR WIRE FUNCTION POPULATE ANYTHING THAT WANTS A "NAVIGATOR".
	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
	})

	// CREATE THE ROOT COMPONENT
	root := &Root{}
	login := &Login{}
	buildEnv.WireComponent(root)  // WIRE IT
	buildEnv.WireComponent(login) // WIRE IT

	// ADD ROUTES FOR EACH PAGE.  NOTE THAT THESE ARE "EXACT" ROUTES.
	// YOU CAN ALSO ADD ROUTES THAT MATCH ANYTHING WITH THE SPECIFIED PREFIX.
	router.MustAddRouteExact("/profile",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &Profile{} // A COMPONENT WITH PAGE CONTENTS
		}))
	router.MustAddRouteExact("/",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &Login{} // A COMPONENT WITH PAGE CONTENTS
		}))
	router.SetNotFound(vgrouter.RouteHandlerFunc(
		func(rm *vgrouter.RouteMatch) {
			//	root.Body = &PageNotFound{} // A PAGE FOR THE NOT-FOUND CASE
		}))

	// TELL THE ROUTER TO LISTEN FOR THE BROWSER CHANGING URLS
	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	// GRAB THE CURRENT BROWSER URL AND PROCESS IT AS A ROUTE
	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}
