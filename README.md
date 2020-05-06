# reads

[![Travis CI](https://img.shields.io/travis/ajbosco/reads.svg?style=flat-square)](https://travis-ci.org/ajbosco/reads)
[![Go Report Card](https://goreportcard.com/badge/github.com/ajbosco/reads?style=flat-square)](https://goreportcard.com/report/github.com/ajbosco/reads)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/ajbosco/reads/goodreads)

Command line tool to interact with [Goodreads](https://www.goodreads.com).

![screenshot](/img/screenshot.png)

- [Installation](#installation)
    + [Binaries](#binaries)
    + [With Go](#with-go)
- [Authentication](#authentication)
- [Usage](#usage)
  * [Search for Book](#search-for-book)
  * [List Shelves](#list-shelves)
  * [Show Books on Shelf](#show-books-on-shelf)
  * [Add Book to Shelf](#add-book-to-shelf)

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/ajbosco/reads/releases).

#### With Go

```console
$ go get github.com/ajbosco/reads/cmd/reads
```

## Authentication

1. Create a Developer Key with [Goodreads](https://www.goodreads.com/api/keys)
2. Create a `config.yml` file
```console
DeveloperKey: your-developer-key
DeveloperSecret: your-developer-secret
```
3. Set the config filepath as `GOODREADS_CLI_CONFIG` environment variable.
```console
export GOODREADS_CLI_CONFIG=path/to/your/config.yml
```
 
## Usage

```console
$ reads -h
NAME:
   reads - Command line tool to interact with Goodreads

USAGE:
   reads [global options] command [command options] [arguments...]

COMMANDS:
     search   search for a book by title, author, or id
     shelves  view shelves and add books to them
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  Goodreads CLI config file [$GOODREADS_CLI_CONFIG]
   --debug, -d               enable debug logging
   --help, -h                show help
   --version, -v             print the version
```

### Search for Book

```console
$ reads search -h
NAME:
   reads search - search for a book by title, author, or id

USAGE:
   reads search [arguments...]
```

### List Shelves

```console
$ reads shelves list -h
NAME:
   reads shelves list - list your shelves

USAGE:
   reads shelves list [arguments...]
```

### Show Books on Shelf

```console
$ reads shelves show -h
NAME:
   reads shelves show - show books on shelf

USAGE:
   reads shelves show [command options] [arguments...]

OPTIONS:
   --shelf value, -s value  -s=shelf-name
```

### Add Book to Shelf

```console
$ reads shelves add -h
NAME:
   reads shelves add - add a book to shelf

USAGE:
   reads shelves add [command options] [arguments...]

OPTIONS:
   --shelf value, -s value    -s=shelf-name
   --book-id value, -b value  -b=book-id
```
