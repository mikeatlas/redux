package v1_test

import (
	"fmt"

	"github.com/dannypsnl/redux/v1/action"
	"github.com/dannypsnl/redux/v1/store"
)

func counter(state interface{}, action action.Action) interface{} {
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

func Example_counter() {
	store := store.New(counter)
	store.Subscribe(func() {
		fmt.Println("Current State:", store.GetState("counter"))
	})
	store.Dispatch(action.New("INC"))
	// Output:
	// Current State: 1
}