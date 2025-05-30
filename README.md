# gotest

### Prerequisites

- Go 1.22.3 installed
-
- Install dependency

```
$ go mod tidy
```

Dynamically generate configuration files and copy them to `$HOME/gotest/etc/gotest-<app>.yaml`. Will apply to all app in this repo.

```
$ make init
```

There are something `make init` will do:

- Make directories: `$HOME/gotest/etc`, `$HOME/gotest/logs`,`$HOME/gotest/install`
- Read environment variables and generate `gotest-<app>.yaml`, eg. `gotest-datasync.yaml` for `datasync`
- Put `gotest-<app>.yaml` into `$HOME/gotest/etc/` which will be read when the app launch

## datasync

### Run in development

```
$ make init
$ go run cmd/datasync/main.go
```

### Run in production

```
$ make start_datasync
```

There are something `make start_datasync` will do:

- `make init`
- Build binary `datasync`(or `datasync.exe`), put it into `$HOME/gotest/install`
- Run `datasync`(or `datasync.exe` )

## ...
