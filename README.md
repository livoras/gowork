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
