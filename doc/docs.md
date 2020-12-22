# Vanilla Documentation

## Operations

| Command                         | Description                                                                                                                                         |
| ------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `vanilla s <package>`             | Sync package |
| `vanilla r <package>`                       | Remove package|
| `vanilla up`                       | Update all packages |
| `vanilla ug`          | Upgrade all packages|
| `vanilla se <package>`                | Search for package|
| `vanilla q` | List all installed packages|
| `vanilla e <package>` | Search for already installed packages|

+ It can be also connected e.g. `vanilla upug`, `vanilla up ug`

## Uploading a package
### You can upload a binary or source:
### Package has to be compressed to tar.gz
+ To upload a binary `vanilla bin <path to binary>`
+ To upload a source `vanilla src <path to source>`