FROM gitpod/workspace-postgres

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/
# Install Golang
RUN sudo apt-get update && sudo apt-get install -y \
        golang \
    && sudo apt-get clean && sudo rm -rf /var/cache/apt/* && sudo rm -rf /var/lib/apt/lists/* && sudo rm -rf /tmp/*

RUN sudo go get -d -u github.com/golang/protobuf/protoc-gen-go && \
    sudo go install github.com/golang/protobuf/protoc-gen-go
