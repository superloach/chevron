chevron
=======
an esoteric yet (somewhat) readable language.

data types
----------
there are two data types, **NUM** (numbers) and **TXT** (text).

generally these inherit properties from types in the interpreter language. (eg. `int` and `str` in the python implementation)

parsers
-------
**VAR**s (variables) are in the form `^c`, where `c` is any character. Special `VAR`s in the form `^_c` are used by the interpreter.

| usage | value |
| --- | --- |
| `^_l` | current line NUM |
| `^_u` | a literal `^` |
| `^_e` | an empty TXT |
| `^_c` | contents of the last comment |

-------

**MIX**es (mixtures) are any text, and can resolve to a **TXT** (applying **VAR** values).

example: `hi ^n!` would produce the **TXT** `hi terra!` if **VAR** `n` was the **TXT** `terra`.

-------

**MAT**s (mathematics) work similarly to **MIX**es, but resolve to **NUM**s and support basic 2-term expressions.

example: `^a+^b` would produce the **NUM** `3` if **VAR** `a` was the **NUM** `1` and **VAR** `b` was the **NUM** `2`.

| operator | purpose | example |
| --- | --- | --- |
| + | sum of operand 1 and operand 2 | `1 + ^a` |
| - | difference of operand 1 and operand 2 | `^b - 2` |
| / | integer quotient of operand 1 and operand 2 | `^n / 3` |
| * | product of operand 1 and operand 2 | `^c * ^d` |
| % | modulus of operand 1 and operand 2 | `5 % 3` |
| < | operand 1 is less than operand 2 | `0 < ^a` |
| > | operand 2 is greater than operand 2 | `^i > 1000` |
| = | operand 1 is equal to operand 2 | `^f = 3` |
| ~ | apply special operation (from operand 2) to operand 1 | `^m ~ p`

| special operation | meaning |
| --- | --- |
| p | is operand 1 prime |
| o | is operand 1 odd |
| e | is operand 1 even |

commands
--------
each line is interpreted as one of the following commands:

| name | description | syntax | parser meanings |
| --- | --- | --- | --- |
| COM | comment | `<>MIX` | contents |
| OUT | output | `>MIX` | text |
| TIN | TXT input | `>MIX>VAR` | prompt, target |
| NIN | NUM input | `>MIX>>VAR` | prompt, target |
| TAS | TXT assignment | `VAR<MIX` | target, text |
| NAS | NUM assignment | `VAR<<MAT` | target, expression |
| IDX | indexing | `VAR<MAT>MIX` | target, index, text |
| HOP | line change | `->MAT` | line |
| SKP | line change on MAT condition | `->MAT?MAT` | line, expression |
| JMP | line change on MIX comparison | `->MAT??MIX=MIX` | line, text 1, text 2 |
| DIE | exit program | `><MIX` | text |
