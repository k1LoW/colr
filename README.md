# colr [![Build Status](https://github.com/k1LoW/colr/workflows/build/badge.svg)](https://github.com/k1LoW/colr/actions) [![GitHub release](https://img.shields.io/github/release/k1LoW/colr.svg)](https://github.com/k1LoW/colr/releases)

`colr` colors strings, colorfully.

![screencast](doc/screencast.svg)

## Usage

``` console
$ tail -F /var/log/nginx/access.log | colr POST GET 404 500 search
```

If you want to erase colors from STDIN, you can use `--erase` option.

``` console
$ any-colorful-command | colr --erase
```

## Install

**homebrew tap:**

```console
$ brew install k1LoW/tap/colr
```

**manually:**

Download binany from [releases page](https://github.com/k1LoW/colr/releases)

**go get:**

```console
$ go get github.com/k1LoW/colr
```
