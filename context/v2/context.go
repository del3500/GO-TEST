package context2

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// context has a method Done() which returns a channel which get sent a signal when the context
		// is "done" or "cancelled". We want to listen to that signal and call store.Cancel if
		//  we get it but we want to ignore it if our store manages to fetch before it.
		//
		// to manage this, we run Fetch in a goroutine and it will write the result into
		// a new channel data. we then use select to effectively race to the two asynchronous processes
		// and then we either write a response or Cancel.
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
