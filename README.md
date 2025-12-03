# subs String substitution library

The `subs` library handles substitution of values into a string, with associated
formatting and minimal logic to handle cardinality, etc.

This package is generally used with messaging or logging subsystems for format a
message (possible drawn from a localization subsystem) that includes directives for
how/where to substitute values into the string, long with the data to be substituted.

## API

There are two ways a string can be given values. The first is using a map, where each
substitution key in the string corresponds to a string key in a `map[string]any` object.
The second from is to pass an arbitrary object, and the substitution operations will
navigate into the structure or map object passed to get the value (using the JAXON
expression package).

```go
   args := map[string]any{"count":52, "kind": "columns"}
   text := "There are {{count|%05d}} items of type {{kind}}"

   msg, err := subs.SubstituteMap(text, args)
```

This creates a map that is used to contain all value(s) needed to complete the string
substitution operations. Note that in addition to the name of the item, additional
formatting information can be given after the "|" character. In the above example, the
count will take up five characters with leading zeros. These format operations
are described later.

```go

   type User struct {
       Name string
       Age  int
   }

   user := User{"Tom", 55}
   msg, err := subs.Substitute("User {{Name}} is {{Age}} years old", user)
```

In this example, the substitution operators refer to elements of the value passed in,
in this case an instance of type `User`. To use this format, the named elements in
the structure must be exported values (i.e. start with a capital letter).

## Format Operators

Below is a table of the formatting operators that can be specified following the "|" character
in the substitution operation. If multiple format operations
are given in a substitution operator, they are processed in order specified.

| Format Operator | Description |
|-----------------|-------------|
| lines           | The item is an array, make a separate line for each array element |
| list            | The item is an array, output each item separated by "," |
| size n          | If the substitution is longer than `n` characters, truncate with `...` ellipses |
| pad "a"         | Use the value to write copies of the string "a" to the output |
| left n          | Left justify the value in a field n characters wide |
| right n.        | Right justify the value in a field n characters wide |
| center n.       | Center justify the value in a field n characters wide |
| empty "text"    | If the value is zero, an empty string, or an empty array, output "text" instead |
| nonempty "Text" | If the value is non-zero, non-empty string, or non-empty array, output "Text" instead |
| zero "text"     | If the value is numerically zero, output "text" instead of the value |
| one "text"      | If the value is numerically one, output "text" instead of the value |
| many "text"     | If the value is numerically greater than one, output "text" instead of the value |
| card "a","b"    | If the value is numerically one, output "a" else output "b" |

These can be combined as needed, and a single value from the map of values can be used multiple
times in substitution operators. Consider the following text message:

```text
There are {{count}} rows
```

In many languages (English included) both the verb and the noun are affected by the cardinality of
the value of count. Additionally, we might want to specify "no rows" when the count is zero. This
can all be done in the substitution operations.  If the message was defined as:

```text
There {{count|card is,are}} {{count|empty "no"}} {{count||card row,rows}}.
```

This uses the count value to control the verb "to be", whether a numeric value or "no" for an empty
value, and the cardinality of the row noun. Note that in the example the `card` format operator
replaces the value with the string, while the `empty` format operator formats the value of `count`
normally using the default integer output, unless the value is empty/zero in which case the string
"no" will be used instead. That is, some operators affect the formatting of the value and other
operators use the value to made decisions about what to output instead of the value itself.

```text
There are no rows.   # For count of zero
There is 1 row.      # For count of 1
There are 32 rows.   # For count of 32
```
