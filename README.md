# Ubdater

A CLI tool to manage changes to Ubity, for areas not supported by their API

## Usage

This programs expects some environment variables to be set. This is for reading  
the ubity username and password from.

```
export UBITY_LOGIN=TODO
export UBITY_PASSWORD=TODO
```

If those are set, this program  can be run like so:

```
./ubdater -e 1 -f 614300123456
```

This is saying for extension `1`, set the forward to number to `614300123456`.

```
NAME:
   ubdater - Ubity update tool

USAGE:
   ubdater [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --extension value, -e value   the extension to modify (default: 0)
   --forward-to value, -f value  the forward to number
   --help, -h                    show help (default: false)
```

## Author

Trent Schafer, (C) 2020

## Liecense

This tool is licensed under MPL-2.0: Mozilla Public Licene, 2.0
