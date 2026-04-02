PROJECT OVERVIEW This project builds a text processing tool using Go. You read text from a file. You apply string transformations. You write the cleaned result to an output file. The goal is to practice: File handling String manipulation Working with multiple functions

HOW THE PROGRAM WORKS You provide an input file The program reads each line It applies transformation rules It writes the result to an output file Run the program like this: go run . input.txt output.txt

FEATURES Convert text to uppercase and lowercase Capitalize words Fix punctuation Convert numbers between formats Clean and format text automatically

PROJECT STRUCTURE file-pipeline/ ├── main.go // controls program flow ├── func.go // contains transformation functions ├── input.txt // raw text input ├── output.txt // processed result └── README.md

CONTRIBUTORS AND ROLES Ojonye Oyilom Caps function, Compiler Kpensalen Aondongu Cap N function, Testing Blessing Eze Hex to Binary, Testing Shaibu Abraham Punctuation Fix, Debugging Owoicho Favour Article Fix, Debugging Aondowase Ama Lowercase Conversion, Testing Emmanuella Uloko Hex to Decimal, Debugging Bashir Hassan Article Fix Innocent Ekwueme Main.go Fix, Compiler Benedict Ondoma Code Management, Project Management Joy Adikwu Lowercase Function, Testing
