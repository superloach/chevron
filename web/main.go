package main

func main() {
	elems()
	query()
	examples()

	loading := document.Call("getElementById", "loading")
	if loading.Truthy() {
		loading.Call("remove")
	}

	select {}
}
