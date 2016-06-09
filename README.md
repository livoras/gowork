Go Daily
===================

###  `make` or `new`

- `make` can only be used to create a channel, slice or map.
- `new` is used to allocate memory of a type.

### `slice` and `array`

- `arrays` is fixed-size structure.
- `slices` is underlying using arrays. When a slice is full, go will create a new array and copy all values in old array into new array (accurately it doubles its size). This is how arrays in other dynamic language like PHP、Python does.

If you declare with variable without specifying the size, you will create a slice; otherwise you create an array.
