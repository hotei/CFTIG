// cftig.go (c) 2013-2014 David Rook - all rights reserved

// go get github.com/hotei/CFTIG should import other programs as side effect
// details can be found in the README.md files for each separate project


package main

// do NOT fix this error, program is not meant to compile
type errorHere


import (
	_ "github.com/hotei/mdserver" // markdown document server
	_ "github.com/hotei/mdr"  // utility function pkg
	_ "github.com/hotei/ls256"	// list files with SHA256 
	_ "github.com/hotei/prune256"	// list, count optionally delete based on RE2
)

/*

*/

func main() {

}
