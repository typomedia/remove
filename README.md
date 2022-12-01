# `remove` - A Painless File Remover

Cross-platform cli file remover written in [Go](https://go.dev/) 1.19.

Mainly written to replace wtf `del|rmdir /s /q` for easi[l]ly deleting files in Windows.

## Usage

    Usage: remove [OPTION]... [FILE]...
    -h, --help
    -i, --input file     Line separated list of files
    [FILE]...            File or directory to remove

## Examples

    remove file1 file2 dir/dir1/*.tmp

### `files.txt`

```
file1
file2
dir/dir1/*.tmp
```

    remove --input files.txt

## Build

    make build

## Requirements

Only needed for embedding a resource icon on Windows:

    go get github.com/akavel/rsrc
    rsrc -ico go.ico
    make build

---
Copyright © 2022 Typomedia Foundation. All rights reserved.
