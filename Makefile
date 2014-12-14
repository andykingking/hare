default:
	bambam -o="." -p="hare" key.go
	sed -i "" s/hare.// translateCapn.go
	capnp compile -ogo ./schema.capnp
	rm -rf testdir_*
	go build
