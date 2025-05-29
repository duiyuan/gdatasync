# godemo

### Prerequisites

- Go 1.22.3 installed
-
- Install dependency

```
$ go mod tidy
```

Dynamically generate configuration files and copy them to `$HOME/godemo/etc/godemo-<app>.yaml`. Will apply to all app in this repo.

```
$ make init
```

There are something `make init` will do:

- Make directories: `$HOME/godemo/etc`, `$HOME/godemo/logs`,`$HOME/godemo/install`
- Read environment variables and generate `godemo-<app>.yaml`, eg. `godemo-datasync.yaml` for `datasync`
- Put `godemo-<app>.yaml` into `$HOME/godemo/etc/` which will be read when the app launch

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
- Build binary `datasync`(or `datasync.exe`), put it into `$HOME/godemo/install`
- Run `datasync`(or `datasync.exe` )

## ...
