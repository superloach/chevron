<>fib.ch
>first n: >>^n
3>>^i
0>>^a
>^a
->x?^n=1
1>>^b
>^b
->x?^n=2
:l
^a+^b>>^c
^b>>^a
^c>>^b
>^c
^i+1>>^i
->x?^i>^n
->l
:x
