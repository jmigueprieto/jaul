package main

import conductor_utils "github.com/jmigueprieto/jaul/conductor"

func main() {
	conductor_utils.RunWorkflow("greetings", 1, map[string]interface{}{"name": "Bond, James"})
}
