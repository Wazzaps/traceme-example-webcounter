all: webcounter
XDG_RUNTIME_DIR ?= /run/user/$(shell id -u)

# ===== Configuration =====
# Path to the trace directory (where compressed traces are stored)
TRACE_DIR ?= $(XDG_RUNTIME_DIR)/traceme-trace
# ===== /Configuration =====

webcounter: main.go
	go build -trimpath -gcflags "-N -l" -o webcounter .

.PHONY: docker-image/webcounter
docker-image/webcounter: webcounter
	docker build -f traceme.Dockerfile -t wazzaps/webcounter-traceme .

.PHONY: start-webcounter
start-webcounter: docker-image/webcounter
	docker run --privileged --rm -it -d \
		--name webcounter-traceme \
		-v $(TRACE_DIR):/var/traceme \
		-p 8000:8000 \
		-p 8080:8080 \
		wazzaps/webcounter-traceme

.PHONY: stop-webcounter
stop-webcounter:
	docker stop webcounter-traceme
