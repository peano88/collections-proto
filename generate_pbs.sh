#!/bin/bash

PROTOS_PATH=protos
ITEM_PROTO_FILE=item.proto
COLL_PROTO_FILE=collection.proto

# Go
protoc -I $PROTOS_PATH $ITEM_PROTO_FILE --go_out=plugins=grpc:$GOPATH/src
protoc -I $PROTOS_PATH $COLL_PROTO_FILE --go_out=plugins=grpc:$GOPATH/src

# Cpp
protoc -I $PROTOS_PATH --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $ITEM_PROTO_FILE
protoc -I $PROTOS_PATH --cpp_out=cpp/build $ITEM_PROTO_FILE
protoc -I $PROTOS_PATH --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $COLL_PROTO_FILE
protoc -I $PROTOS_PATH --cpp_out=cpp/build $COLL_PROTO_FILE