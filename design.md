# Command Line Utility - Design
The reason for this package is to add some extra functionality simular to the flag package that are inline with the way I like to write cli tools

### Args
Args are the parsed and sorted arguments given to the tool though command line. Args come in three forms:
- Flag
- Var
- Loosie

##### Flag
A flag is a booliean style argument, if the flag is found in the arg list a corrisponding boolean value is set and can be accessed using the `Flag` function

##### Var
A var is an arg that needs a value with it, setting a var can be done in many ways using a space, colon or equal sign. If space is needed in the value wrap it with double quotes `"`
	``` Examples:
		tool /var=toolInfo
		tool -var: toolInfo
		tool --var toolinfo
		tool /var:toolInfo
		tool /str="This is a string with some spaces"
	```
##### Loosie
A loosie is any arg that isn't sorted as a flag or a var.


### App
An app is a simple was to run "commands" so giving it a simple interface to handle something like git cli
