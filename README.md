## Go interface:
### Basic Structure:
1. type struct and its functions
``` type PostDatabase struct{}

func (self *PostDatabase) Insert() bool {
	fmt.Println(`
	** create: client send data without giving any identifier. 
	- If the client sends data without any identifier, then we will store the data and assign/generate a new identifier.
	- If the client again sends the same data without any identifier, then we will store the data and assign/generate a new identifier.
	Note: Duplication is allowed here.
	`)
	return true
}
func (self *PostDatabase) Delete() bool {
	return true
}
```
2. object initialization and passed to a caller function by parameter
```
func foo(db PostDatabase) {
	db.Insert()
	db.Delete()
}

dbPost := &PostDatabase{}
foo(dbPost)
```
3. If a different feature requires a different object with the same function. Then we need to create a new caller function for that object becase Go is type specific
```
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

func foo(db PostDatabase) {
	db.Insert()
	db.Delete()
}

dbPut := &PutDatabase{}
foo(dbPut)
```
4. **Now by using interface as the caller function's parameter, we can pass in any type of object to the caller function such that we don't need to write duplicate codes**
```
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

dbPost := &PostDatabase{}
foo(dbPost)

dbPut := &PutDatabase{}
foo(dbPut)
```

### What is the difference between PUT, POST and PATCH?
**Source:** https://stackoverflow.com/questions/31089221/what-is-the-difference-between-put-post-and-patch
1. POST - **send no id in the payload**
- Will store the data and assign/generate a new identifier.
- Duplication is allowed here when same data is sent again.
- If the request contains an ID (could be a dup of existing), then better [reject the request](https://stackoverflow.com/questions/33452765/what-to-do-when-rest-post-provides-an-id). 
2. PUT - **must have id in the payload**
- If the client sends data with an identifier, then we will check whether that identifier exists.
- If the identifier exists, we will update the resource with the data, else we will create a resource with the data and assign/generate a new identifier.
- "update and overwrite", replace the ENTIRE RESOURCE with the new representation provided.
3. PATCH - **must have id in the payload**
- If the client sends data with an identifier, then we will check whether that identifier exists. If the identifier exists, we will update the resource with the data, else we will throw an exception.
- "update and merge", replace parts of the source resource with the values provided AND|OR other parts of the resource are updated that you havent provided (timestamps) AND|OR updating the resource effects other resources (relationships)

#### Note: On the PUT method, we are not throwing an exception if an identifier is not found. But in the PATCH method, we are throwing an exception if the identifier is not found.
