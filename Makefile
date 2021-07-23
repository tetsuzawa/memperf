NEWRELIC_LICENSE_KEY :=

all:
	less Makefile

/var/lib/memperf:
	mkdir -p $@
	git clone https://github.com/tetsuzawa/memperf $@

/var/run/memperf:
	tar -xvf aaa

/etc/cron.d/memperf: /var/lib/memperf
	cp -f $</cron.d/memperf $@

.PHONY: run
run:
	NEWRELIC_LICENSE_KEY=$(NEWRELIC_LICENSE_KEY) /var/run/memperf
