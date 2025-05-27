
TARGETs := datasync

.PHONE: init
init:
	for app in $(TARGETs); do\
		bash scripts/init/$$app.sh; \
	done;


.PHONE: build_datasync
build_datasync:
	go build cmd/datasync/main.go


.PHONE: start_datasync
start_datasync:
	bash scripts/start/datasync.sh