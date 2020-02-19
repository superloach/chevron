>enter some words: >^a
^a~s>>^s

<>count
0>>^i
1>>^c
:a
^i+1>>^i
->b?^i>^s
^a,^i~c>>^l
->+2??^l=^_p
->a
^c+1>>^c
->a
:b

<>random
^c~r>>^r
^r+0.5>>^r
^r,0~d>>^r

<>index
0>>^i
1>>^n
:c
^i+1>>^i
->d?^i>^s
->d?^n=^r
^a,^i~c>>^l
->+2??^l=^_p
->c
^n+1>>^n
->c
:d

<>build
^__>^w
:e
->f?^i>^s
^a,^i~c>>^l
->f??^l=^_p
^w^l>^w
^i+1>>^i
->e
:f

<>output
><random word (^r): ^w
