build:
	go build

crossbuild:
	goxz -pv=v0.0.1 -os=darwin -d=pkg

.PHONY: build crossbuild
