package main

import "fmt"

type PostDatabase struct{}

func (self *PostDatabase) Insert() bool {
	fmt.Println(`
	** create: client send data without giving any identifier. 
	- If the client sends data without any identifier, then we will store the data and assign/generate a new identifier.
	- If the client again sends the same data without any identifier, then we will store the data and assign/generate a new identifier.
 	- If the request contains an ID (could be a dup of existing), then better reject the request.
	Note: Duplication is allowed here.
	`)
	return true
}
func (self *PostDatabase) Delete() bool {
	return true
}

type PutDatabase struct{}

func (self *PutDatabase) Insert() bool {
	fmt.Println(`
	** Replace row: 
	- The client sends data with an identifier, then we will check whether that identifier exists. If the identifier exists, we will update the resource with the data, else we will create a resource with the data and assign/generate a new identifier.
	`)
	return true
}
func (self *PutDatabase) Delete() bool {
	return true
}

type PatchDatabase struct{}

func (self *PatchDatabase) Insert() bool {
	fmt.Println(`
	** Update row: 
	- The client sends data with an identifier, then we will check whether that identifier exists. If the identifier exists, we will update the resource with the data, else we will throw an exception.
	`)
	return true
}
func (self *PatchDatabase) Delete() bool {
	return true
}

type DAI interface {
	Insert() bool
	Delete() bool
}

// the benefit of interface is, the caller functions below accept any types
// meaning we don't have to write duplicate functions for different types of inputs
func foo(db DAI) {
	db.Insert()
	db.Delete()
}

func bar(db DAI) {
	db.Insert()
	db.Delete()
}

func baz(db DAI) {
	db.Insert()
	db.Delete()
}

func main() {
	dbPost := &PostDatabase{}
	foo(dbPost)
	bar(dbPost)
	baz(dbPost)

	dbPut := &PutDatabase{}
	foo(dbPut)
	bar(dbPut)
	baz(dbPut)

	dbPetch := &PatchDatabase{}
	foo(dbPetch)
	bar(dbPetch)
	baz(dbPetch)
}
