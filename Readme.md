## Fundamentals

* #### [Data types](Fundamentals/data_types.md)
* #### [Format verbs](Fundamentals/format_verbs.md)
* #### [GoTour exercise](GoTour)

## Microservice
* #### [GET example](Microservices/main.go)


## Command 
##### go.mod
The go.mod file is used to define modules and their dependencies in a Go project. 

To initialize a Go module, 
```
go mod init <module_name>
```
## Syntax 

#### Function naming convention

Exported Methods: If a method name starts with a capital letter, it means the method is exported from the package and can be accessed by code outside of the package where it is defined.

Unexported Methods: If a method name starts with a lowercase letter, it means the method is unexported, and it is only accessible within the package where it is defined.
