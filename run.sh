#!/bin/sh
GRPC_ADDR=":9000" ./grpc/server/user
GRPC_ADDR=":9001" ./grpc/server/movie
# --- END --- #