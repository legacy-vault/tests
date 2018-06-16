#!/bin/bash

protoc -I protocol protocol/protocol.proto --go_out=plugins=grpc:protocol
