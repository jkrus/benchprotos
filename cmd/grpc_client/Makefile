.PHONY: pre-push go-mod
pre-push: go-mod

go-mod:
	go mod tidy -v
	go mod download -x
