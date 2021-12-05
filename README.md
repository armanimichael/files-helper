# File Helper

CLI tool used to alter / read large number of textfiles.

Common use cases consist in quickly replacing patterns in files or looking for them.

## Commands

`root`: Root path (default `./`)

`pattern`: Search pattern

`replace`: What to replace the search pattern with (default)

`extensions`: Filter files extensions (separated by commas, ex. txt,css) (default `txt`)

`backup`: Backups files matching the search pattern (default `false`)

`cmd`: Util to run

- `find`: Logs files containing a `pattern`
- `replace`: Logs files containing a `pattern` and substitute with `replace`

<hr />

### Examples:

`file-helper -cmd find -pattern test` -> Finds files containing test (in the current folder)

`file-helper -cmd replace -pattern test -replace tset -backup` -> Finds files containing test and replaces it with tset (in the current folder)

`file-helper -cmd find -pattern test -root /somedir` -> Finds files containing test (in the `/somedir` directory)

<hr />  
