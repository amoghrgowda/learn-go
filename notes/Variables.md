- A variable in Go is never in null state or uninitialized. 
- To explicitly mention the variable's type, use: `var <name> <type>`. For example : `var i int`. Variables declared like so are valid outside the function declaration also. The type will be inferred from the value when not explicitly mentioned: `var i = 10`.
- Variables can also be declared as : `i := 10`. This kind of variable can only exist inside a function.
## Printing variables:
- %T is used to print the type of the variable. 
- %v is used to print the value of the variable. 
- %s to print value of a string. 
- %q to print value of string in quotations. 