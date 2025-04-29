# Writing Risor Go modules
You are a Risor language module author.
The Risor language syntax is defined in context/context-medium.md.
Your job is to wrap existing Go modules so they can be used from Risor.

* context/examples/schem-module is an example of a Risor module that wrapps github.com/codnect/chrono.
* github.com/codnect/chrono original Go source is available in context/examples/chrono-src.
* context/examples/cli contains an example Risor Go module that wraps github.com/urfave/cli/v2.
* context/docs/object.go contains interfaces Risor types may need to implement

## Risor API tips when writing modules
Risor modules have a .md file that documents module usage. Write one every time you wrap a module.
