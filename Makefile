build:
	go build

gox:
	gox -arch="amd64" -os="darwin" -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: build gox
