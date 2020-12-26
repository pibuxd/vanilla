<div align="center">
<h1>Vanilla</h1>

[pibux.pl](https://pibux.pl) | [repos](https://pibux.pl/repos)
</div>
<div align="center">
</div>

## Objectives
There's a point in everyone's life when you feel the need to write a package manager because there are only about ∞ of them. So say hi to ∞+1.
But this is quite nice:
+ minimalism
+ fast
+ light
+
### which is not working yet
+ dependiencies and conflicts

## Installation
### manual
```sh
git clone https://github.com/pibuxd/vanilla ${HOME}/vanilla
cd vanilla
make build
```
+ in future in repositories of other package managers

<h3>If you want to remove just</h3> `make remove`
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
+ have to be compressed into tar.gz
+ Makefile has to be included where `make install` installs and `make remove` removes a package
+ the first time the package is thrown in, it needs a password
+ every update of package needs password entry
