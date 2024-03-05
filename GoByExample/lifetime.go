package main

/**
Here, x must be heap-allocated because it is still reachable from the variable global after f has returned, despite being declared as a local variable; we say x escapes from f. 
Conversely, when g returns, the variable *y becomes unreachable and can be recycled. 
Since *y does not escape from g, it’s safe for the compiler to allocate *y on the stack, even though it was allo- cated with new. 

In any case, the notion of escaping is not something that you need to worry about in order to write correct code, though it’s good to keep in mind during performance optimization, since each variable that escapes requires an extra memory allocation.
**/
var global *int
var boolean * bool
func globalPtr() {
	var x int
	x = 1
	// keeping unnecessary pointers to short-lived objects within long-lived objects, especially global variables, will prevent the gc from reclaiming the short-lived objects.
	global = &x
}


func localPtr() {
	y := new(int)
	*y = 2
}
