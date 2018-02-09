# exclude

Remove all lines in file X from file Y.

```
usage: exclude [FORBIDDEN file] [INPUT file]

Iterates though the lines in INPUT, and prints the lines that aren't also in
FORBIDDEN to stdout.
```

## Example

```sh
$ cat forbidden.txt
hello
$ cat input.txt
hello
world
$ exclude forbidden.txt input.txt
world
```

## Performance

This tool was written to improve the performance of [`grep -vxfF`](https://unix.stackexchange.com/questions/299462/how-to-filter-out-lines-of-a-command-output-that-occur-in-a-text-file). Here's an arbitrary performance comparison:

```sh
$ wc -l forbidden.txt
    22108
$ wc -l input.txt
    48518
$ time grep -vxfF forbidden.txt input.txt
900.88s user 2.93s system 99% cpu 15:10.36 total
$ time exclude forbidden.txt input.txt
0.05s user 0.15s system 93% cpu 0.221 total
```

A cool 4000 times faster.

## Memory

`exclude` reads the contents of `forbidden.txt` into memory. Performance may degrade for very large `forbidden.txt` files.

## Install

With `go get`:
```sh
$ go get -u github.com/jamesroutley/exclude
```
