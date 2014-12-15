default: build

build:
	bambam -o="." -p="hare" sequencer.go view.go document.go
	sed -i "" s/hare.// translateCapn.go
	capnp compile -ogo ./schema.capnp
	rm -rf testdir_*
	go build

test: build
	go test -v

run: build
	go run cmd/hare/hare.go
