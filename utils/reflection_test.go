//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testFunc() {
}

func testFunc1(i int) {
}

type structT struct {
}

func (structT) test() {

}

var _ = ginkgo.Describe("reflection", func() {
	ginkgo.It("GetFuncName should work", func() {
		Ω(GetFuncName(func() {})).Should(ContainSubstring("func"))
		Ω(GetFuncName(func() {})).Should(ContainSubstring("func"))
		Ω(GetFuncName(testFunc)).Should(Equal("testFunc"))
		Ω(GetFuncName(testFunc1)).Should(Equal("testFunc1"))
		t := &structT{}
		Ω(GetFuncName(t.test)).Should(Equal("test-fm"))
	})
})
