# GraphTheoryProject
This is a college assignment for Graph Theory , We must write a program in the Go programming language [2] that can build a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text.

## Steps to run:
1. Download GOLANG from https://golang.org.
2. Git clone this repository into desired location.
3. CMD into the location of main.go.
4. Type "go run main.go"

## Steps to use program:
1. Enter regular expression.
2. Enter string you want to test.

![alt text](https://github.com/KeithH4666/GraphTheoryProject/blob/master/Images/Capture.PNG)

## How it works:

1. First the regular expression is converted to postFix so that order doesn't matter.
2. Then the postfix regular expression is converted to an NFA which the string will be tested against.

## References: 

1. https://swtch.com/~rsc/regexp/regexp1.html
2. Lecture notes by our lecturer Ian Mcloughlin.
