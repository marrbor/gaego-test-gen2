project:=gae-go112-test
port:=8888
dse-port:=8081

export DATASTORE_EMULATOR_HOST:=localhost:$(dse-port)
export DATASTORE_PROJECT_ID:=$(project)

run:
	go run server.go

start-dse:
	gcloud beta emulators datastore start --project=$(project) --host-port=localhost:$(dse-port)

inst-dse:
	gcloud components install cloud-datastore-emulator
