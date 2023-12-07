all:
	go build -ldflags "-w -s -H windowsgui -extldflags --static"

