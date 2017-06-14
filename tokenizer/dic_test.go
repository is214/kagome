// Copyright 2015 ikawaha
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tokenizer

import (
	"testing"
)

var testDic = "../_sample/ipa.dic"

func TestNewDic(t *testing.T) {
	d, err := NewDic(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected, c := 392126, len(d.dic.Morphs); c != expected {
		t.Errorf("got %v, expected %v\n", c, expected)
	}
	if expected, c := 392126, len(d.dic.Contents); c != expected {
		t.Errorf("got %v, expected %v\n", c, expected)
	}
}

func TestSysDic(t *testing.T) {
	a := SysDic()
	b := SysDic()
	if a.dic != b.dic {
		t.Errorf("got %p and %p, expected singleton", a.dic, b.dic)
	}
}

func TestSysDicIPA(t *testing.T) {
	a := SysDicIPA()
	b := SysDicIPA()
	if a.dic != b.dic {
		t.Errorf("got %p and %p, expected singleton", a.dic, b.dic)
	}
}

func TestSysDicIPASimple(t *testing.T) {
	a := SysDicIPASimple()
	b := SysDicIPASimple()
	if a.dic != b.dic {
		t.Errorf("got %p and %p, expected singleton", a.dic, b.dic)
	}
}

func TestSysDicUni(t *testing.T) {
	a := SysDicUni()
	b := SysDicUni()
	if a.dic != b.dic {
		t.Errorf("got %p and %p, expected singleton", a.dic, b.dic)
	}
}

func TestSysDicUniSimple(t *testing.T) {
	a := SysDicUniSimple()
	b := SysDicUniSimple()
	if a.dic != b.dic {
		t.Errorf("got %p and %p, expected singleton", a.dic, b.dic)
	}
}
