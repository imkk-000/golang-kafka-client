#!/bin/sh

$SRC_DIR=protobuf
$DST_DIR=model

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto

