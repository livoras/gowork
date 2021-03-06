Go Daily
===================

###  `make` or `new`

- `make` can only be used to create a channel, slice or map.
- `new` is used to allocate memory of a type.

### `slice` and `array`

- `arrays` is fixed-size structure.
- `slices` is underlying using arrays. When a slice is full, go will create a new array and copy all values in old array into new array (accurately it doubles its size). This is how arrays in other dynamic language like PHP、Python does.

If you declare with variable without specifying the size, you will create a slice; otherwise you create an array.

using `copy` to copy value to another slice.

### about `map`

- `len()` to checkout the size of a map
- `value, ok = map[key]` to check if a value exists in a map
- `for key, value := range aMap` to loop over a map

### Go is always copying value

If you don't using pointer, Go passes values to functions by copying it. So if you have a large structure, always using pointer to pass.

### packages
Go treats subfolders as sub-packages in you project. All variables with starts with a capital letter will be exported.

### Using `errors.New` to create an error

The built-in error interface:

```
type error interface {
  Error() string
}
```

You can create your own error type by implement the error interface.

You can also use `errors.New("error message")` to create and error

### strings and bytes array
1 byte = 8 bits
1 rune = 32 bits
1 int  = 64 bits

But if you use `[]byte` to cast string is not enough, because strings may contains unicode characters. And `[]rune` is a unicode version of `[]byte`. Usually use it to convert string to array or slice.

When `range` a string, you actually ranging a `[]rune` by default.

## function is first-class in golang
Function is first-class in golang. You use them as value, return value, parameter as normal variables.

## channels
Channels are communication pipes between goroutines which is used to pass data.

## select
`select` blocks util one of its case can run, then execute that case. It choose randomly when multiple case are ready.

If a channel has no buffer (when trying to read) or is full (when trying to write), then it's not ready.

In case that using `<- c` to read data, if there is no data received in channel, then it's not ready.

In case than using `c <- data` to write data into channel, this channel has to have been specified a buffer size when it's declared and is not full to be ready.

## os.Getenv to get environment variable


## `iota` auto increment variable
```
const (
  One = iota
  Two
  Three
)
```

## When declare a slice, you can use indexes to place.

var str = []string{
  0: "shi",
  3: "fuck"
}

## Using `sync/atomic` to manage shared state atomically by address.

https://gobyexample.com/atomic-counters
