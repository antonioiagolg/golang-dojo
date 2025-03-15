package main

import (
    "fmt"
    "antonioiagolg.github.io/golang-dojo/workspaces/greet"
)

// The workspaces serve as a way of multiple modules to colaborate, without the
// need to use replace statement in the go.mod definition.
// This module contains a reference to reverse module in go.mod
// without the replace statement. The same module is defined in the workspace
// So when we perform go run ./hello, the go will check that the dependency is 
// defined in the workspace, so it will use it in this module and our changes
// (the Int function) will be visible. This is a great way to work with modules without
// the need of use the replace method. To keep track of the latest versions of reverse
// module, we can use the go get command in our module
func main() {
    fmt.Println(greet.Greet("Antonio"))
}

