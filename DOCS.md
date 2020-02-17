chevron docs
============

data types
----------
there are two data types, **NUM** (numbers) and **TXT** (text).

underneath, **NUM**s and **TXT**s are really all strings, but **NUM**s are converted to floats for operations.

**VAR**s (variables) are in the form `^c`, where `c` is any character. **VAR**s in the form `^_c` are special values.

| special var | value |
| --- | --- |
| `_#` | current line **NUM** |
| `_c` | last comment |
| `_r` | greater than `>` |
| `_l` | less than `<` |
| `_u` | caret `^` |
| `_q` | question mark `?` |
| `_d` | dash `-` |
| `_s` | underscore `_` |
| `_e` | equal `=` |
| `_o` | colon `:` |
| `_n` | newline `\n` |
| `_t` | tilde `~` |
| `_b` | backtick `` ` `` |
| `_a` | lowercase english alphabet |
| `_i` | numbers 0-9 |
| `__` | nothing |

parsers
-------

**MIX**es (mixtures) are any text, and can resolve to a **TXT** (applying **VAR** values).

example: `hi ^n!` would produce the **TXT** `hi terra!` if **VAR** `n` was the **TXT** `terra`.

-------

**MAT**s (mathematics) work similarly to **MIX**es, but use special 2-term expressions.

example: `^a+^b` would produce the `3` if **VAR** `a` was the **NUM** `1` and **VAR** `b` was the **NUM** `2`.

| type | purpose | example |
| --- | --- | --- |
| + | sum of operand 1 and operand 2 | `1 + ^a` |
| - | difference of operand 1 and operand 2 | `^b - 2` |
| / | quotient of operand 1 and operand 2 | `^n / 3` |
| * | product of operand 1 and operand 2 | `^c * ^d` |
| % | modulus of operand 1 and operand 2 | `5 % 3` |
| < | operand 1 is less than operand 2 | `0 < ^a` |
| > | operand 2 is greater than operand 2 | `^i > 1000` |
| = | operand 1 is equal to operand 2 | `^f = 3` |
| ` | operand 1 to the power of operand 2 | ``^g ` 2`` |
| ~ | apply special type | `^m ~ p` |

| special type | purpose | example |
| --- | --- | --- |
| p | is operand 1 prime | `23 ~ p` |
| o | is operand 1 odd | `^n ~ o` |
| e | is operand 1 even | `^n ~ e` |
| r | random **NUM** between 0 and operand 1 | `100 ~ r` |
| n | boolean negation of operand 1 | `1 ~ n` |
| l | lowercase of operand 1 | `^t ~ l` |
| u | uppercase of operand 1 | `^t ~ u` |
| v | reverse of operand 1 | `hello ~ v` |
| d | round operand 1 | `3.1415,2 ~ d` |
| c | cut operand 1 | `hello,1,4 ~ c`
| s | size of operand 1 | `^a ~ s` |
| f | index operand 1 | `abc,b ~ i` |

operations
----------
each line is interpreted as one of the following operations:

| name | description | syntax | parser meanings |
| --- | --- | --- | --- |
| EMP | empty line | ` ` | |
| COM | comment | `<>TXT` | contents |
| LBL | comment | `:TXT` | name |
| OUT | output | `>MIX` | text |
| TIN | TXT input | `>MIX>VAR` | prompt, target |
| NIN | NUM input | `>MIX>>VAR` | prompt, target |
| TAS | TXT assignment | `MIX>VAR` | target, text |
| NAS | NUM assignment | `MAT>>VAR` | target, expression |
| HOP | line change | `->MAT` | line |
| SKP | line change on MAT condition | `->MAT?MAT` | line, expression |
| JMP | line change on MIX comparison | `->MAT??MIX=MIX` | line, text 1, text 2 |
| DIE | exit program | `><MIX` | text |

(JMP comparisons: `=` for equality, `<` for before, `>` for after, `@` for contains)
