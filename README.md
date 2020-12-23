<div align="center">
<h1>Vanilla</h1>

[pibux.pl](https://pibux.pl) |
</div>
<div align="center">
</div>

## Objectives
There's a point in everyone's life when you feel the need to write a package manager because there are only about ∞ of them. So say hi to ∞+1.
But this is quite nice:
+ minimalistic
+ fast
+ light
### it just works
Vanilla works on most Linux
## Installation
```sh
git clone https://github.com/pibuxd/vanilla
cd vanilla
make build
```

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
| `vanilla cr` | Create new package|
+ It can be also connected e.g. `vanilla up ug`
