# Go Format Verbs

| Category                   | Verb | Description                                                                                         |
|----------------------------|------|-----------------------------------------------------------------------------------------------------|
| **General Verbs**          | `%v` | The value in a default format.                                                                      |
|                            | `%+v`| The value in a default format; for structs, the plus flag adds field names.                         |
|                            | `%#v`| A Go-syntax representation of the value.                                                            |
|                            | `%T` | A Go-syntax representation of the type of the value.                                                |
|                            | `%%` | A literal percent sign; consumes no value.                                                          |
| **Boolean**                | `%t` | The word `true` or `false`.                                                                         |
| **Integer**                | `%b` | Base 2.                                                                                             |
|                            | `%c` | The character represented by the corresponding Unicode code point.                                  |
|                            | `%d` | Base 10.                                                                                            |
|                            | `%o` | Base 8.                                                                                             |
|                            | `%O` | Base 8 with 0o prefix.                                                                              |
|                            | `%q` | A single-quoted character literal safely escaped with Go syntax.                                    |
|                            | `%x` | Base 16, with lower-case letters for a-f.                                                           |
|                            | `%X` | Base 16, with upper-case letters for A-F.                                                           |
|                            | `%U` | Unicode format: U+1234; same as "U+%04X".                                                           |
| **Floating Point and Complex** | `%b` | Decimalless scientific notation with exponent a power of two.                                       |
|                            | `%e` | Scientific notation (e.g., -1.234456e+78).                                                          |
|                            | `%E` | Scientific notation (e.g., -1.234456E+78).                                                          |
|                            | `%f` | Decimal point but no exponent.                                                                      |
|                            | `%F` | Same as `%f`.                                                                                       |
|                            | `%g` | Use the shortest representation: `%e` or `%f`.                                                      |
|                            | `%G` | Use the shortest representation: `%E` or `%F`.                                                      |
|                            | `%x` | Hexadecimal notation (with decimal power of two exponent, e.g., -0x1.23abcp+20).                    |
|                            | `%X` | Upper-case hexadecimal notation (e.g., -0X1.23ABCP+20).                                             |
| **String and Slice of Bytes** | `%s` | The uninterpreted bytes of the string or slice.                                                    |
|                            | `%q` | A double-quoted string safely escaped with Go syntax.                                               |
|                            | `%x` | Base 16, lower-case, two characters per byte.                                                       |
|                            | `%X` | Base 16, upper-case, two characters per byte.                                                       |
| **Pointer**                | `%p` | Base 16 notation, with leading 0x.                                                                  |
