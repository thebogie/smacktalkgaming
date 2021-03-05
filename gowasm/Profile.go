package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vugu/vugu"
	js "github.com/vugu/vugu/js"
)

//Profile : Profile
type Profile struct {
	Loading bool     `vugu:"data"`
	Items   []string `vugu:"data"`

	shortItemCount int
}

//Init : Init
func (c *Profile) Init(ctx vugu.InitCtx) {

	// kick of loading data from an endpoint
	//TODO: how to share Login struct across app...
	log.Printf("\n\n json object:::: %+v", user)

	c.Loading = true
	go func() {

		resp, err := http.Get("/some/endpoint")
		if err != nil {
			log.Printf("Error fetching: %v", err)
			return
		}
		defer resp.Body.Close()
		var items []string
		err = json.NewDecoder(resp.Body).Decode(&items)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
			return
		}

		ctx.EventEnv().Lock()
		c.Loading = false
		c.Items = items
		ctx.EventEnv().UnlockRender()
	}()
}

//Compute : Compute
func (c *Profile) Compute() {

	// recompute each render

	count := 0
	for _, item := range c.Items {
		if len(item) < 5 {
			count++
		}
	}
	c.shortItemCount = count

}

//Rendered : Rendered
func (c *Profile) Rendered(ctx vugu.RenderedCtx) {

	// if you really need to manipulate DOM directly after it is rendered, you can do it here

	if ctx.First() { // only after first render
		el := js.Global().Get("document").Call("getElementById", "some_id_here")
		_ = el // do something with an element manually after the first render
	}

}

//Destroy : Destroy
func (c *Profile) Destroy() {
	// some teardown code here
}
