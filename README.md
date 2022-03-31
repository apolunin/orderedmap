# Go Ordered Maps with Generics

This data structure works the same way as a regular map, but keeps the order in which
keys were inserted into the map. Implementation uses Go generics, so it
requires Go 1.18+ to run.

Features:
* All operations are done in constant time.
* Provides API for iterating over map entries in a keys insertion order.
* Uses Go generics.

## Installation
```bash
go get -u github.com/apolunin/orderedmap
```
## Documentation
The complete documentation is available on [godoc.org](https://pkg.go.dev/github.com/apolunin/orderedmap)