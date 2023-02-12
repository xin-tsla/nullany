# nullany

Nullany is a library to have nullable types for Golang.

## NullAny Struct

NullAny struct is for any type.

## NullBool Struct

NullBool is a nullable bool type with handling logical `AND` and `OR`
NOTICE: in NullBool's `AND` and `OR`, if one of bool value is null, then the result of `AND` and `OR` will be the other
bool value.
For example,

```

Null AND Null = Null

True AND Null = True
False AND Null = False

True OR Null = True
False OR Null = False

```

Please be sure the above behavior is what you expect before you use this library.
If the behavior does nit fit your need, please use the NullAny.

