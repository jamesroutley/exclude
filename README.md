# exclude

Remove all lines in file X from file Y.

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
