# Go `fmt.Printf` Cheat Sheet

`fmt.Printf` formats text according to a format specifier and writes to standard output. 
**Note:** It does **not** append a newline automatically (use `\n`).

**Syntax:** `fmt.Printf("format string", args...)`

## 1. General & Debugging
Use these to inspect values of any type.

| Verb  | Description | Example Output |
| :---  | :--- | :--- |
| `%v`  | **Default** format for the value. | `{10}` |
| `%+v` | Adds **field names** (structs only). | `{x:10}` |
| `%#v` | **Go-syntax** representation. | `main.Point{x:10}` |
| `%T`  | Prints the **type** of the value. | `int` |
| `%%`  | Literal percent sign. | `%` |

## 2. Scalar Types

### Integers
| Verb | Description | Example Input (`255`) |
| :--- | :--- | :--- |
| `%d` | Base 10 (Decimal). | `255` |
| `%b` | Base 2 (Binary). | `11111111` |
| `%o` | Base 8 (Octal). | `377` |
| `%x` | Base 16 (Hex lower). | `ff` |
| `%X` | Base 16 (Hex upper). | `FF` |
| `%U` | Unicode format. | `U+00FF` |

### Floats
| Verb | Description | Example Input (`123.456`) |
| :--- | :--- | :--- |
| `%f` | Decimal (no exponent). | `123.456000` |
| `%.2f` | Precision limited to 2 decimals. | `123.46` |
| `%e` | Scientific notation. | `1.234560e+02` |
| `%g` | Compact (removes trailing zeros). | `123.456` |

### Strings & Bytes
|