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

package automl

import (
	"iter"

	automlpb "cloud.google.com/go/automl/apiv1beta1/automlpb"
	"github.com/googleapis/gax-go/v2/iterator"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ColumnSpecIterator) All() iter.Seq2[*automlpb.ColumnSpec, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DatasetIterator) All() iter.Seq2[*automlpb.Dataset, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelEvaluationIterator) All() iter.Seq2[*automlpb.ModelEvaluation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelIterator) All() iter.Seq2[*automlpb.Model, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TableSpecIterator) All() iter.Seq2[*automlpb.TableSpec, error] {
	return iterator.RangeAdapter(it.Next)
}
