##### Data types

| Data Type | Size (bits) | Example Value |
|-----------|-------------|---------------|
| `bool`    | 1           | `true`        |
| `byte`    | 8           | `0x1F`        |
| `rune`    | 32          | `'a'`         |
| `int`     | 32 or 64    | `42`          |
| `int8`    | 8           | `-128`        |
| `int16`   | 16          | `32767`       |
| `int32`   | 32          | `2147483647`  |
| `int64`   | 64          | `9223372036854775807` |
| `uint`    | 32 or 64    | `42`          |
| `uint8`   | 8           | `255`         |
| `uint16`  | 16          | `65535`       |
| `uint32`  | 32          | `4294967295`  |
| `uint64`  | 64          | `18446744073709551615` |
| `float32` | 32          | `3.14`        |
| `float64` | 64          | `3.141592653589793` |
| `complex64` | 64        | `1 + 2i`      |
| `complex128` | 128      | `1 + 2i`      |
| `string`  | variable    | `"Hello"`     |
| `array`   | variable    | `[3]int{1, 2, 3}` |
| `slice`   | variable    | `[]int{1, 2, 3}` |
| `map`     | variable    | `map[string]int{"one": 1}` |
| `struct`  | variable    | `struct{name string}{name: "Alice"}` |
| `interface` | variable  | `interface{}` |
| `chan`    | variable    | `make(chan int)` |
| `func`    | variable    | `func(x int) int { return x * 2 }` |
