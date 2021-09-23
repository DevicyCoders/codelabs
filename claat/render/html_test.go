// Copyright 2016 Google Inc. All Rights Reserved.
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

package render

import (
	"testing"

	"github.com/googlecodelabs/tools/claat/nodes"
)

func TestHTMLEnv(t *testing.T) {
	one := nodes.NewTextNode("one ")
	one.MutateEnv([]string{"one"})
	two := nodes.NewTextNode("two ")
	two.MutateEnv([]string{"two"})
	three := nodes.NewTextNode("three ")
	three.MutateEnv([]string{"one", "three"})

	tests := []struct {
		env    string
		output string
	}{
		{"", "one two three "},
		{"one", "one three "},
		{"two", "two "},
		{"three", "three "},
		{"four", ""},
	}
	for i, test := range tests {
		var ctx Context
		ctx.Env = test.env
		h, err := HTML(ctx, one, two, three)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}
		if v := string(h); v != test.output {
			t.Errorf("%d: v = %q; want %q", i, v, test.output)
		}
	}
}

// TODO: test HTML
// TODO: test writeHTML
// TODO: test ReplaceDoubleCurlyBracketsWithEntity
// TODO: test matchEnv
// TODO: test write
// TODO: test writeString
// TODO: test writeFmt
// TODO: test escape
// TODO: test writeEscape
// TODO: test text
// TODO: test image
// TODO: test url
// TODO: test button
// TODO: test code
// TODO: test list
// TODO: test onlyImages
// TODO: test itemsList
// TODO: test grid
// TODO: test infobox
// TODO: test survey
// TODO: test header
// TODO: test youtube
// TODO: test iframe
