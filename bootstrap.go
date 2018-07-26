package main

func main() {
	go FileServ()
	go HttpServ()
	select {}
}
