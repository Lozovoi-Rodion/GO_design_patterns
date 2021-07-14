package main

import "github.com/Lozovoi-Rodion/GO_design_patterns/proxy/examples"

// Proxy - functions as an interface to a particular resource. That resource
// may be remote, expensive to construct, or may require logging or some other
// added functionality.

// Difference between proxy and decorator:
// 1. proxy tries to provide an identical interface; decorator provides enhanced interface
// 2. decorator typically aggregates (or has pointer to) what it is decorating; proxy doesn't have to
// 3. Proxy might not even be working with a materialized object

// Summary:
// - proxy has the same interface as the underlying obj
// - to create a proxy simply replicate the existing interface of an object
// - add relevant functionality to the redefined methods
// - different proxies (communication, logging, caching, etc.) have completely different behaviors

func main()  {
	car := proxy.NewCarProxy(&proxy.Driver{12})
	car.Drive()

	bmp := proxy.NewBitmap("demo.png")
	proxy.DrawImage(bmp)

	// lazy bitmap is initialized only when needed
	lbmp := proxy.NewLazyBitmap("demo.png")
	proxy.DrawImage(lbmp)
}
