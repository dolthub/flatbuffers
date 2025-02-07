#!/bin/bash -eu
#
# Copyright 2014 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

pushd "$(dirname $0)" >/dev/null
test_dir="$(pwd)"

# Emit Go code for the example schemas in the test dir:
../flatc -g --gen-object-api -I include_test --go-module-name github.com/google/flatbuffers/tests monster_test.fbs optional_scalars.fbs
../flatc -g --gen-object-api -I include_test/sub --go-module-name github.com/google/flatbuffers/tests include_test/order.fbs
../flatc -g --gen-object-api --go-module-name github.com/google/flatbuffers/tests -o Pizza include_test/sub/no_namespace.fbs

# Run tests with necessary flags.
# Developers may wish to see more detail by appending the verbosity flag
# -test.v to arguments for this command, as in:
#   go -test -test.v ...
# Developers may also wish to run benchmarks, which may be achieved with the
# flag -test.bench and the wildcard regexp ".":
#   go -test -test.bench=. ...
CGO_ENABLED=0 go test . \
                     --coverpkg=github.com/dolthub/flatbuffers/v23/go \
                     --cpp_data=${test_dir}/monsterdata_test.mon \
                     --out_data=${test_dir}/monsterdata_go_wire.mon \
                     --bench=. \
                     --benchtime=3s \
                     --fuzz=true \
                     --fuzz_fields=4 \
                     --fuzz_objects=10000

GO_TEST_RESULT=$?
if [[ $GO_TEST_RESULT  == 0 ]]; then
    echo "OK: Go tests passed."
else
    echo "KO: Go tests failed."
    exit 1
fi

NOT_FMT_FILES=$(gofmt -l .)
if [[ ${NOT_FMT_FILES} != "" ]]; then
    echo "These files are not well gofmt'ed:"
    echo
    echo "${NOT_FMT_FILES}"
    # enable this when enums are properly formated
    # exit 1
fi

# Re-enable go modules when done tests
go env -w  GO111MODULE=on
