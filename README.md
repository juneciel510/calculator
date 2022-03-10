# Calculator

Commandline user interface
Inspired by [mrcaidev/calculator](https://github.com/mrcaidev/calculator), however developed a totally different scheme.<br/>
Easy to upgrade with more operators.

---

## Instruction for running

Use command line to run exe file in each sub folder.<br/>
e.g Change directory to `calculator_4` and run with command line `.\calculator_4.exe`

---

## Basic idea

Consider it as a pipeline problem and solve following the below steps:

1. handle negative numbers
2. remove all the parentheses, start from innermost one first
3. remove `log, sin, cos`
4. remove `^`
5. remove `* /`
6. remove `+ -`

---

## calculator_1

Very simple calculator which need to input operator and numbers separately<br/>
Operators: `+ - * /`<br/>
Numbers: e.g `1, 2.8, 98552`

---

## calculator_2

Simple calculator input expression like `2*3.8+9-10/4`<br/>
Operators: `+ - * /`

---

## calculator_3

Improved calculator input expression like `7*(8.7/2-3)+log(9.5-3/8)^8` <br/>
Operators: `+ - * / ^ log sin cos` <br/>
`Note:`<br/>
Nested parentheses will cause error<br/>
Accept two forms of log operator:

- log(number): natural logarithm
- log(number, base)

---

## calculator_4

Advanced calculator input expression like `-7*(9.5+9-3+9^2/9*32+(ln(3+8.9*3/100)^2+3-4+7)*9+log2(sin(-9)-3+9^2/9))*32` <br/>
Operators: `+ - * / ^ ln log sin cos` <br/>
e.g.<br/>

- ln(8): natural logarithm
- log2(5.8): 2--base
- sin(10.8)
- cos(11)

`Note:`<br/>
Stop using stack, re-design the scheme, making it more accurate & simpler<br/>
Fix nested parentheses error<br/>

`New features:` 1.allow negative numbers, 2.log base on any number (08,Mar,2022)<br/>

`To-do:`

- graphical user interface
- add more operators and constants( e.g. "π e" )
