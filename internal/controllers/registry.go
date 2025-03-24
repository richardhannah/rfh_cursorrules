package controllers

import "github.com/gorilla/mux"

type RouteSetter func(router *mux.Router)

var routeSetters []RouteSetter

// RegisterRouteSetter registers a route-setting function.
func RegisterRouteSetter(fn RouteSetter) {
	routeSetters = append(routeSetters, fn)
}

// SetAllRoutes calls all registered route-setting functions.
func SetAllRoutes(router *mux.Router) {
	for _, fn := range routeSetters {
		fn(router)
	}
}
