NEWRELIC_LICENSE_KEY :=

run:
	NEWRELIC_LICENSE_KEY=$(NEWRELIC_LICENSE_KEY) memperf

update/frame:
	curl http://localhost:9999/internal/update/frame
