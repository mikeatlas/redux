package main

import (
	"fmt"

	"github.com/dannypsnl/redux"
)

func counter(state interface{}, action redux.Action) interface{} {
	// initial state
	if state == nil {
		return 0
	}
	switch action.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func main() {
	store := redux.NewStore(counter)
	store.Subscribe(func() {
		fmt.Printf("State is %v\n", store.GetState["counter"])
	})
	for i := 0; i < 10; i++ {
		store.Dispatch(redux.SendAction("INC"))
	}
	for i := 0; i < 10; i++ {
		store.Dispatch(redux.SendAction("DEC"))
	}
}
