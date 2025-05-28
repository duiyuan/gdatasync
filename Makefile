
TARGETs := datasync

.PHONE: init
init:
	bash scripts/init/all.sh

.PHONE: build_datasync
build_datasync:
	go build cmd/datasync/main.go


.PHONE: start_datasync
start_datasync:
	bash scripts/start/datasync.sh