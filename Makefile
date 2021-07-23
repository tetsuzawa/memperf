NEWRELIC_LICENSE_KEY :=
RELEASE_TAG := v0.0.3
OS := linux
ARCH := x86_64

.PHONY: all install setup run update/frame
all:
	less Makefile

install:
	yum install -y curl git wget

setup: install
	$(MAKE) /var/run/memperf
	$(MAKE) /var/lib/memperf


/var/lib/memperf:
	mkdir -p $@
	git clone https://github.com/tetsuzawa/memperf $@

/var/run/memperf_$(OS)_$(ARCH).tar.gz:
	https://github.com/tetsuzawa/memperf/releases/download/$(RELEASE_TAG)/memperf_$(OS)_$(ARCH).tar.gz


/var/run/memperf: /var/run/memperf_$(OS)_$(ARCH).tar.gz
	tar -xvf $< $@

/etc/cron.d/memperf: /var/lib/memperf
	cp -f $</cron.d/memperf $@

run: /var/run/memperf /etc/cron.d/memperf
	NEWRELIC_LICENSE_KEY=$(NEWRELIC_LICENSE_KEY) $<

update/frame:
	curl http://localhost:9999/internal/update/frame
