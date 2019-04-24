sleepy
===

> sleep with progress and formated duration

# Install

```
go get -u github.com/iwittkau/sleepy
``` 

# Usage

```
sleepy duration [message]
```

Sleep for 10 minutes:

```
$ sleepy 10m
[=>----------------------------------------------------------]   3.00% 09m11s
``` 

Sleep for 10 seconds with message:

```
$ sleepy 10s "ten seconds"
 ten seconds [============================================================] 100.00%
 2019/04/23 19:55:10 finished "ten seconds"
``` 

Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h` as in Go's `time` [package](https://godoc.org/time#ParseDuration).

# The Progress Bar

`sleepy`'s progress bar is created with the [cheggaaa/pb](https://github.com/cheggaaa/pb) library.