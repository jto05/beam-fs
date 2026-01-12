SERVER_EXECUTABLE=beam_s
CLIENT_EXECUTABLE=beam_c
BIN_PATH=build/bin
TMUX_SESSION=beam-fs

.PHONY: build run-server run-client

build:
	go build -o $(BIN_PATH)/$(SERVER_EXECUTABLE)  cmd/beam_server/main.go 
	go build -o $(BIN_PATH)/$(CLIENT_EXECUTABLE)  cmd/beam_client/main.go 

run:
	tmux new-session -d -s $(TMUX_SESSION) 'make run-server'
	tmux split-window -h $(TMUX_SESSION) 'make run-client' 
	tmux attach -t $(TMUX_SESSION)


run-server:
	$(BIN_PATH)/$(SERVER_EXECUTABLE)

run-client:
	$(BIN_PATH)/$(CLIENT_EXECUTABLE)

clean:
	rm $(BIN_PATH)/$(SERVER_EXECUTABLE)
	rm $(BIN_PATH)/$(CLIENT_EXECUTABLE)

