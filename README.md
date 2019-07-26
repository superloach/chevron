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

MIXes (mixtures) are any text, and can contain/resolve VAR references to a TXT.

example: `hi ^n!` would produce the TXT `hi terra!` if `^n` was the TXT `terra`.

MATs (mathematics) work similarly to MIXes, but resolve to NUMs and support basic expressions.

example: `^a+^b` would produce the NUM `3` if VAR a was the NUM `1` and VAR b was the NUM `2`.

commands
--------

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
