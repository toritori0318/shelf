Shelf
=====

## Description

Shelf is for easy putting in and out of data.

## Usage

```
USAGE:
   shelf [global options] command [command options] [arguments...]

COMMANDS:
   get		get <key>
   put		put <key> <value>
   delete	delete <key>
   get-aws	get-aws desuyo
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/toritori0318/go-shelf
```

## Environment

### SHELF_HOME

   SHELF_HOME must be set only when installing to a custom location. (Default: **$HOME/.shelf** )

## Examples

```bash
# set value
$ shelf put /alpaca/money 100yen

# get key
$ shelf get /alpaca/money
100yen

# delete key
$ shelf delete /alpaca/money



# set json
$ shelf put /alpaca/json '{"hoge":"fuga"}'

# get json
$ shelf get /alpaca/json | jq .
{
    "hoge": "fuga"
}

```



## Contribution

1. Fork ([https://github.com/toritori0318/go-shelf/fork](https://github.com/toritori0318/go-shelf/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[toritori0318](https://github.com/toritori0318)
