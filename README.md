# `remove` - A Painless File Remover

Cross-platform cli file remover written in [Go](https://go.dev/) 1.19.

Mainly written to replace wtf `del|rmdir /s /q` for easi[l]ly deleting files in Windows.

## Usage

    Usage: remove [OPTION]... [FILE]...
    -h, --help
    -i, --input file     line separated list of files
    -r, --recursive      remove files and folders recursively
    -v, --verbose        output what is being done
    [FILE]...            file or directory to remove

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

    go install github.com/akavel/rsrc@latest
    rsrc -ico go.ico
    make build

---
Copyright © 2022 Typomedia Foundation. All rights reserved.
