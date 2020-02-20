package main

func main() {
	setup()
	query()

	loading := document.Call("getElementById", "loading")
	if loading.Truthy() {
		loading.Call("remove")
	}

	select {}
}
