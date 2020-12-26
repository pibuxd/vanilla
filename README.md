<div align="center">
<h1>Vanilla</h1>

[pibux.pl](https://pibux.pl) | [repos](https://pibux.pl/repos)
</div>
<div align="center">
</div>

## Objectives
Vanilla is built for people who want to easily share projects. Also for official packages.
+ minimalism
+ fast
+ light

### which is not working yet
+ update and upgrade (will work very very soon)
+ dependencies and conflicts
+ passwords for created packages (just skip password entering)

## Installation
### manual
```sh
git clone https://github.com/pibuxd/vanilla ${HOME}/vanilla
cd vanilla
make build
```
### do not remove manually .vanilla !!

### If you want to remove just
`make remove`

## Operations

| Command                         | Description                                                                                                                                         |
| ------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `vanilla s <package(s)>`             | Sync package(s) |
| `vanilla r <package(s)>`                       | Remove package(s)|
| `vanilla up`                       | Update all packages |
| `vanilla ug`          | Upgrade all packages|
| `vanilla se <package>`                | Search for package|
| `vanilla q` | List all installed packages|
| `vanilla e <package>` | Search for already installed packages|
| `vanilla cr` | Create new package to the repositories|
+ It can be also connected e.g. `vanilla up ug`

## Adding a new package to the [repositories]("https://pibux.pl/repos")
+ it can be a binary and source :)
+ have to be compressed into tar.gz
+ Makefile has to be included where `sudo make install` installs and `sudo make remove` removes a package
+ the first time the package is thrown in, it needs a password
+ every update of package needs password entry
