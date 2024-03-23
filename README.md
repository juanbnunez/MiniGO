# Mini GO Compiler

## Overview

The Mini GO Compiler project aims to implement a compiler for a subset of GoLang known as "Mini GO." This README provides an overview of the project's goals, requirements, and implementation details.

## Project Definition

The project involves developing the syntax analysis phase of the compiler, including scanning, parsing, and AST generation. ANTLR4 is the chosen tool for this phase, and the implementation will be done using the Go programming language.

For detailed information on the Mini GO grammar, refer to the "MiniGo Grammar.docx" document provided.

## Requirements

- Implement a scanner to recognize all valid tokens in Mini GO, following the format of the original GoLang.
- Create test files that cover all syntactic capabilities of the language.
- Verify the syntactic correctness of these tests using the Mini GO parser.
- Develop a text editor capable of displaying and parsing the created tests. The editor should highlight syntax errors, indicate their location, and provide easy identification of errors by line and column.

## Developers

- Juan Núñez Parrales
- Jennifer Gonzales Solís

## Getting Started

1. Clone the repository.
2. Install ANTLR4 and any required dependencies.
3. Follow the instructions in the project's documentation for building and running the Mini GO Compiler.

## Documentation

For detailed documentation on the language analysis, implementation details, test coverage, results, conclusions, and references, refer to the project's documentation files.
