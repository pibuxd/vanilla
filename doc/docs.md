# Vanilla Documentation

## Introduction

Vanilla is a simple package manager written in Go that can be runned on all Linux machines.

It's similar to other package managers, like Pacman or APT.

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