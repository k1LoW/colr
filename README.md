# colr [![Build Status](https://github.com/k1LoW/colr/workflows/build/badge.svg)](https://github.com/k1LoW/colr/actions) [![GitHub release](https://img.shields.io/github/release/k1LoW/colr.svg)](https://github.com/k1LoW/colr/releases)

:art: `colr` colors strings, colorfully.

![screencast](doc/screencast.svg)

## Usage

``` console
$ tail -F /var/log/nginx/access.log | colr POST GET 404 500 search
```

**Erase colors:**

If you want to erase colors from STDIN, you can use `--erase` option.

``` console
$ any-colorful-command | colr --erase
```

## Install

**deb:**

``` console
$ export COLR_VERSION=X.X.X
$ curl -o colr.deb -L https://github.com/k1LoW/colr/releases/download/v$COLR_VERSION/colr_$COLR_VERSION-1_amd64.deb
$ dpkg -i colr.deb
```

**RPM:**

``` console
$ export COLR_VERSION=X.X.X
$ yum install https://github.com/k1LoW/colr/releases/download/v$COLR_VERSION/colr_$COLR_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export COLR_VERSION=X.X.X
$ curl -o colr.apk -L https://github.com/k1LoW/colr/releases/download/v$COLR_VERSION/colr_$COLR_VERSION-1_amd64.apk
$ apk add colr.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/colr
```

**go install:**

```console
$ go install github.com/k1LoW/colr/cmd/colr@vX.X.X
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/colr/releases)

