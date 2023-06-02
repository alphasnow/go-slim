package app

var (
	// BUILD_DATE		= $(shell date -u '+%Y%m%d%I%M%S')
	// go build -ldflags "-w -s -X 'go-slim/pkg/app.buildTag=$(RELEASE_TAG)' -X 'go-slim/pkg/app.buildDate=prod'"
	buildTag  = "0.0.0"
	buildDate = "20221108090000"
)
