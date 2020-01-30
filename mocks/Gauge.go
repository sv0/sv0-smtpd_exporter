// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io_prometheus_client "github.com/prometheus/client_model/go"
	mock "github.com/stretchr/testify/mock"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

// Gauge is an autogenerated mock type for the Gauge type
type Gauge struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *Gauge) Add(_a0 float64) {
	_m.Called(_a0)
}

// Collect provides a mock function with given fields: _a0
func (_m *Gauge) Collect(_a0 chan<- prometheus.Metric) {
	_m.Called(_a0)
}

// Dec provides a mock function with given fields:
func (_m *Gauge) Dec() {
	_m.Called()
}

// Desc provides a mock function with given fields:
func (_m *Gauge) Desc() *prometheus.Desc {
	ret := _m.Called()

	var r0 *prometheus.Desc
	if rf, ok := ret.Get(0).(func() *prometheus.Desc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*prometheus.Desc)
		}
	}

	return r0
}

// Describe provides a mock function with given fields: _a0
func (_m *Gauge) Describe(_a0 chan<- *prometheus.Desc) {
	_m.Called(_a0)
}

// Inc provides a mock function with given fields:
func (_m *Gauge) Inc() {
	_m.Called()
}

// Set provides a mock function with given fields: _a0
func (_m *Gauge) Set(_a0 float64) {
	_m.Called(_a0)
}

// SetToCurrentTime provides a mock function with given fields:
func (_m *Gauge) SetToCurrentTime() {
	_m.Called()
}

// Sub provides a mock function with given fields: _a0
func (_m *Gauge) Sub(_a0 float64) {
	_m.Called(_a0)
}

// Write provides a mock function with given fields: _a0
func (_m *Gauge) Write(_a0 *io_prometheus_client.Metric) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*io_prometheus_client.Metric) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
