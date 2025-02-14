-> A go project can be of two types: An executable or a library/package. If the file is saying `package main`, it means that it is an executable.
-> A standard Go project structure looks like this: 
myproject/
│── go.mod          # Defines the module and dependencies
│── go.sum          # Stores checksums for module dependencies
│── main.go         # Main entry point of the application
│── /cmd            # Executable commands (if multiple binaries)
│── /pkg            # Reusable packages
│── /internal       # Private packages (restricted access)
│── /api            # API definitions
│── /configs        # Configuration files
│── /scripts        # Utility scripts
│── /docs           # Documentation
│── /test           # Additional test data
└── /vendor         # Local copies of dependencies (optional)

## Go module:
- A **Go module** is a self-contained collection of Go packages that are versioned together. It allows better dependency management.
- A go.mod file is created when `go mod init` command is run.
- It can contain : Module name, version of go used and dependencies. For example:
```
module github.com/username/myproject
go 1.20 
require ( github.com/gin-gonic/gin v1.9.0 golang.org/x/crypto v0.1.0 )
```
- Notice that the module name will be prefixed by the name of the hosting service provider such as github, etc.
## Some important commands:
- `go mod init <module_name>`: initialize new go module.
- `go mod tidy`: clean up unused dependencies.
- `go list -m all`: list all modules in the project.
- `go get <package>`: Adds a new dependency.
- `go build`: compiles the project using modules
## Some important things to remember:
- A Go package is a collection of related .go files together (basically just a module without a go.mod file and therefore, not self-contained). For example: fmt,math,etc.
- A Go project can contain multiple packages, modules, dependencies and directories.
- A Go module only manages dependencies for a specific part of the project.
- go.work file can be used to link different modules together.
- go modules are self-contained. Therefore, even within the same project, different modules may contain the exact same .go name for their files.
## Multi-module project in Go:

myproject/
│── go.work         # (Go workspace file for multiple modules)
│── /module1        # First module
│   ├── go.mod
│   ├── main.go
│── /module2        # Second module
│   ├── go.mod
│   ├── utils.go
│   ├── main.go
└── /sharedlib      # Shared package used by both modules
    ├── go.mod
    ├── helpers.go

--> Multiple modules are used in micro-services where different services within the same repo can be managed independently.
Or, sometimes, we want to write modules that contain reusable code. This code can be reused/shared between multiple other modules.

--> To use go work to link them together :
```
go work init 
go work use ./module1 
go work use ./module2 
go work use ./sharedlib
```
