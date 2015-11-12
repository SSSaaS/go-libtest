# go-libtest
Cross-language library tester written in Go.

    Copyright (C) 2015 Alexander Scheel, Joel May, Matthew Burket  
    See Contributors.md for a complete list of contributors.  
    Licensed under the MIT License.  

## Usage
Create the directories `libs` and `tmp` in the root of this project. Inside libs, clone the sssa-libraries ([Ruby](https://github.com/SSSAAS/sssa-ruby), [Golang](https://github.com/SSSAAS/sssa-golang), and [Python](https://github.com/SSSAAS/sssa-python)) or copy from your development directory.

To run all tests, simply run `go run ./go-libtest.go` from the project directory.

## Contributing
We welcome pull requests, issues, security advice on this tool, or other contributions you feel are necessary. Feel free to open an issue to discuss any questions you have about this library.

This is the core testing environment for all languages.

For security issues, send a GPG-encrypted email to <alexander.m.scheel@gmail.com> with public key [0xBDC5F518A973035E](https://pgp.mit.edu/pks/lookup?op=vindex&search=0xBDC5F518A973035E).
