sleepy
===

> sleep with progress and formated duration

# Install

```
go get -u github.com/iwittkau/sleepy
``` 

# Usage

```
sleepy duration
```

Sleep for 10 minutes:

```
$ sleepy 10m
[==>---------------------------------------------------------------------------------------------]   3.00% 09m11s
``` 

Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h` as in Go's `time` [package](https://godoc.org/time#ParseDuration).

# The Progress Bar

`sleepy`'s progress bar is created with the [cheggaaa/pb](https://github.com/cheggaaa/pb) library.