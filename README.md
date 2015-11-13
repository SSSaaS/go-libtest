# go-libtest
Cross-language library tester written in Go.

    Copyright (C) 2015 Alexander Scheel, Joel May, Matthew Burket  
    See Contributors.md for a complete list of contributors.  
    Licensed under the MIT License.  

## Usage
Create the directory `libs` in the root of this project. Inside libs, clone the sssa-libraries ([Ruby](https://github.com/SSSAAS/sssa-ruby), [Golang](https://github.com/SSSAAS/sssa-golang), and [Python](https://github.com/SSSAAS/sssa-python)) or copy from your development directories.

To execute all tests, run `go run ./go-libtest.go ./specs/*.json` from the project directory.

Sample output:

    loading specs: ./specs/short-chains.json
    ok @ Null string in 10.488 s
    ok @ Japanese hello world in 10.694 s
    ok @ Chinese hello world in 10.663 s
    ok @ Hello World! in 10.464 s
    ok @ Unicode across int break in 10.723 s
    ok @ Long strings in 10.672 s
    ok @ Insane strings in 10.611 s
    ok @ Middle-null string in 10.534 s
    ------
    total: 84.850 s

    loading specs: ./specs/long-chains.json
    ok @ Null string in 16.939 s
    ok @ Japanese hello world in 10.849 s
    ok @ Chinese hello world in 11.653 s
    ok @ Hello World! in 15.213 s
    ok @ Unicode across int break in 10.631 s
    ok @ Long strings in 10.867 s
    ok @ Insane strings in 11.850 s
    ok @ Middle-null string in 10.288 s
    ------
    total: 98.291 s

## Contributing
We welcome pull requests, issues, security advice on this tool, or other contributions you feel are necessary. Feel free to open an issue to discuss any questions you have about this library.

This is the core testing environment for all languages.

For security issues, send a GPG-encrypted email to <alexander.m.scheel@gmail.com> with public key [0xBDC5F518A973035E](https://pgp.mit.edu/pks/lookup?op=vindex&search=0xBDC5F518A973035E).
