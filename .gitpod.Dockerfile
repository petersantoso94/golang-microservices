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

# install protoc
RUN sudo curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP && \
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    sudo unzip -o $PROTOC_ZIP -d /usr/local include/* && \
    rm -f $PROTOC_ZIP
