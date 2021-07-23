NEWRELIC_LICENSE_KEY :=

.PHONY: all install setup run update/frame
all:
	less Makefile

install:
	yum install -y curl git

setup: install
	$(MAKE) /var/run/memperf
	$(MAKE) /var/lib/memperf


/var/lib/memperf:
	mkdir -p $@
	git clone https://github.com/tetsuzawa/memperf $@

/var/run/memperf:
	tar -xvf aaa

/etc/cron.d/memperf: /var/lib/memperf
	cp -f $</cron.d/memperf $@

run: /var/run/memperf /etc/cron.d/memperf
	NEWRELIC_LICENSE_KEY=$(NEWRELIC_LICENSE_KEY) $<

update/frame:
	curl http://localhost:9999/internal/update/frame
