// Code generated by array/numeric.gen.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package array

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/arrow"
)

// A type which represents an immutable sequence of int64 values.
type Int64 struct {
	array
	values []int64
}

func NewInt64Data(data *Data) *Int64 {
	a := &Int64{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Int64) Value(i int) int64    { return a.values[i] }
func (a *Int64) Int64Values() []int64 { return a.values }

func (a *Int64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of uint64 values.
type Uint64 struct {
	array
	values []uint64
}

func NewUint64Data(data *Data) *Uint64 {
	a := &Uint64{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Uint64) Value(i int) uint64     { return a.values[i] }
func (a *Uint64) Uint64Values() []uint64 { return a.values }

func (a *Uint64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of float64 values.
type Float64 struct {
	array
	values []float64
}

func NewFloat64Data(data *Data) *Float64 {
	a := &Float64{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Float64) Value(i int) float64      { return a.values[i] }
func (a *Float64) Float64Values() []float64 { return a.values }

func (a *Float64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Float64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Float64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of int32 values.
type Int32 struct {
	array
	values []int32
}

func NewInt32Data(data *Data) *Int32 {
	a := &Int32{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Int32) Value(i int) int32    { return a.values[i] }
func (a *Int32) Int32Values() []int32 { return a.values }

func (a *Int32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of uint32 values.
type Uint32 struct {
	array
	values []uint32
}

func NewUint32Data(data *Data) *Uint32 {
	a := &Uint32{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Uint32) Value(i int) uint32     { return a.values[i] }
func (a *Uint32) Uint32Values() []uint32 { return a.values }

func (a *Uint32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of float32 values.
type Float32 struct {
	array
	values []float32
}

func NewFloat32Data(data *Data) *Float32 {
	a := &Float32{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Float32) Value(i int) float32      { return a.values[i] }
func (a *Float32) Float32Values() []float32 { return a.values }

func (a *Float32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Float32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Float32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of int16 values.
type Int16 struct {
	array
	values []int16
}

func NewInt16Data(data *Data) *Int16 {
	a := &Int16{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Int16) Value(i int) int16    { return a.values[i] }
func (a *Int16) Int16Values() []int16 { return a.values }

func (a *Int16) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int16) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int16Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of uint16 values.
type Uint16 struct {
	array
	values []uint16
}

func NewUint16Data(data *Data) *Uint16 {
	a := &Uint16{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Uint16) Value(i int) uint16     { return a.values[i] }
func (a *Uint16) Uint16Values() []uint16 { return a.values }

func (a *Uint16) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint16) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint16Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of int8 values.
type Int8 struct {
	array
	values []int8
}

func NewInt8Data(data *Data) *Int8 {
	a := &Int8{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Int8) Value(i int) int8   { return a.values[i] }
func (a *Int8) Int8Values() []int8 { return a.values }

func (a *Int8) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int8) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int8Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of uint8 values.
type Uint8 struct {
	array
	values []uint8
}

func NewUint8Data(data *Data) *Uint8 {
	a := &Uint8{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Uint8) Value(i int) uint8    { return a.values[i] }
func (a *Uint8) Uint8Values() []uint8 { return a.values }

func (a *Uint8) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint8) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint8Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of arrow.Timestamp values.
type Timestamp struct {
	array
	values []arrow.Timestamp
}

func NewTimestampData(data *Data) *Timestamp {
	a := &Timestamp{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Timestamp) Value(i int) arrow.Timestamp        { return a.values[i] }
func (a *Timestamp) TimestampValues() []arrow.Timestamp { return a.values }

func (a *Timestamp) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Timestamp) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.TimestampTraits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of arrow.Time32 values.
type Time32 struct {
	array
	values []arrow.Time32
}

func NewTime32Data(data *Data) *Time32 {
	a := &Time32{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Time32) Value(i int) arrow.Time32     { return a.values[i] }
func (a *Time32) Time32Values() []arrow.Time32 { return a.values }

func (a *Time32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Time32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Time32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of arrow.Time64 values.
type Time64 struct {
	array
	values []arrow.Time64
}

func NewTime64Data(data *Data) *Time64 {
	a := &Time64{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Time64) Value(i int) arrow.Time64     { return a.values[i] }
func (a *Time64) Time64Values() []arrow.Time64 { return a.values }

func (a *Time64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Time64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Time64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of arrow.Date32 values.
type Date32 struct {
	array
	values []arrow.Date32
}

func NewDate32Data(data *Data) *Date32 {
	a := &Date32{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Date32) Value(i int) arrow.Date32     { return a.values[i] }
func (a *Date32) Date32Values() []arrow.Date32 { return a.values }

func (a *Date32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Date32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Date32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

// A type which represents an immutable sequence of arrow.Date64 values.
type Date64 struct {
	array
	values []arrow.Date64
}

func NewDate64Data(data *Data) *Date64 {
	a := &Date64{}
	a.refCount = 1
	a.setData(data)
	return a
}

func (a *Date64) Value(i int) arrow.Date64     { return a.values[i] }
func (a *Date64) Date64Values() []arrow.Date64 { return a.values }

func (a *Date64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Date64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Date64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}
