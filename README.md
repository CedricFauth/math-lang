# mathematical expression interpreter

written in go

## Grammar

```python
declaration := statement #| fun_dec
statement   := expression

expression  := equality | var_dec
assignment  := IDENTIFIER '=' term

equality       → comparison ( '!='|'==' comparison )*
comparison     → term ( '>'|'>='|'<'|'<=' term )*

term        := factor ( '+'|'-' factor)*
factor      := unary ( '*'|'/' unary)*
unary       := '-' unary | primary
primary     := '(' expression ')' | NUMBER | IDENTIFIER
```
