default:
	bambam -o="." -p="hare" key.go
	sed -i "" s/hare.// translateCapn.go
	capnp compile -ogo ./schema.capnp
	go build
