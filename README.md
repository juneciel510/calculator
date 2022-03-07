# Calculator

`Note:`<br/>
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

1. remove all the parentheses, start from innermost one first
2. remove `log, sin, cos`
3. remove `^`
4. remove `* /`
5. remove `+ -`

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

Advanced calculator input expression like `7*(9.5+9-3+9^2/9*32+(log(3+8.9*3/100)^2+3-4+7)*9+sin9-3+9^2/9)*32`
Operators: `+ - * / ^ log sin cos`

`Note:`
Stop using stack, rewrite the main algorithm which is more accurate & simpler
Nested parentheses error solved
Only allow log(number): natural logarithm

#todo:
add log(number,base)

- detect `,` first, and remove base and set the base individually
- then follow the normal procedure
- take multiple log into consideration
