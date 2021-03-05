package main

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

//Root : root
type Root struct {
	ShowWasm bool `vugu:"data"`
	ShowGo   bool `vugu:"data"`
	ShowVugu bool `vugu:"data"`

	// ANYTHING THAT MUST NAVIGATE NEED ONLY EMBED THIS
	vgrouter.NavigatorRef

	// THE BODY COMPONENT, GETS SET BY THE APPROPRIATE ROUTE ABOVE
	Body vugu.Builder
}
