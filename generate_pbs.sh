#!/bin/bash

PROTOS_PATH=protos
ITEM_PROTO_FILE=item.proto
COLL_PROTO_FILE=collection.proto

# Go
protoc -I $PROTOS_PATH $ITEM_PROTO_FILE --go_out=plugins=grpc:$GOPATH/src
protoc -I $PROTOS_PATH $COLL_PROTO_FILE --go_out=plugins=grpc:$GOPATH/src

# Cpp
protoc --proto_path=$PROTOS_PATH --cpp_out=cpp/build $ITEM_PROTO_FILE
protoc --proto_path=$PROTOS_PATH --cpp_out=cpp/build $COLL_PROTO_FILE