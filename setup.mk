RELEASE_TAG := v0.0.6
OS := linux
ARCH := x86_64

.PHONY: all install setup run update/frame
all:
	less Makefile

install:
	yum install -y curl git wget

setup: install
	$(MAKE) -f setup.mk /usr/local/bin/memperf
	$(MAKE) -f setup.mk /var/lib/memperf
	$(MAKE) -f setup.mk /etc/cron.d/memperf


/var/lib/memperf:
	mkdir -p $@
	git clone https://github.com/tetsuzawa/memperf $@

/var/lib/memperf_$(OS)_$(ARCH).tar.gz:
	mkdir -p $(@D)
	wget -O $@ https://github.com/tetsuzawa/memperf/releases/download/$(RELEASE_TAG)/memperf_$(OS)_$(ARCH).tar.gz


/usr/local/bin/memperf: /var/lib/memperf_$(OS)_$(ARCH).tar.gz
	tar -xvf $< -C $(@D)

/etc/cron.d/memperf: /var/lib/memperf
	cp -f $</cron.d/memperf $@
