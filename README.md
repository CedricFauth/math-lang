# mathematical expression interpreter

written in go

## Grammar

```
expression  := term
term        := factor ( '+'|'-' factor)*
factor      := unary ( '*'|'/' unary)*
unary       := '-' unary | primary
primary     := '(' expression ')' | number
```
