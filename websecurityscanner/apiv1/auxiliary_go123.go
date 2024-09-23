// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package websecurityscanner

import (
	"iter"

	websecurityscannerpb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/googleapis/gax-go/v2/iterator"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *CrawledUrlIterator) All() iter.Seq2[*websecurityscannerpb.CrawledUrl, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FindingIterator) All() iter.Seq2[*websecurityscannerpb.Finding, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ScanConfigIterator) All() iter.Seq2[*websecurityscannerpb.ScanConfig, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ScanRunIterator) All() iter.Seq2[*websecurityscannerpb.ScanRun, error] {
	return iterator.RangeAdapter(it.Next)
}
