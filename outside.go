package main

import global_variable "golang/global-variable"

func main() {
	global_variable.Inside()
	println(global_variable.Global)
}
