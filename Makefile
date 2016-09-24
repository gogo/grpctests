# Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are
# met:
#
#     * Redistributions of source code must retain the above copyright
# notice, this list of conditions and the following disclaimer.
#     * Redistributions in binary form must reproduce the above
# copyright notice, this list of conditions and the following disclaimer
# in the documentation and/or other materials provided with the
# distribution.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
# "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
# LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
# A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
# OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
# SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
# LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
# THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

all:
	make regenerate
	make gofmt
	make test

regenerate:
	(cd simple && protoc-min-version --version="3.0.0" --gogo_out=plugins=grpc:. --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ grpc.proto)
	(cd bench && protoc-gen-combo --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ --gogo_out=plugins=grpc:. bench.proto)
	(cd gofast && protoc --gofast_out=plugins=grpc:. gofast.proto)

test:
	go test -v ./...

gofmt:

	gofmt -l -s -w .

buildserverall:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/protoc-min-version
	go get github.com/gogo/protobuf/protoc-gen-combo
	go get github.com/gogo/protobuf/protoc-gen-gofast
	make all
