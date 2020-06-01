// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import indexer "magma/orc8r/cloud/go/services/state/indexer"
import mock "github.com/stretchr/testify/mock"
import types "magma/orc8r/cloud/go/services/state/types"

// Indexer is an autogenerated mock type for the Indexer type
type Indexer struct {
	mock.Mock
}

// CompleteReindex provides a mock function with given fields: from, to
func (_m *Indexer) CompleteReindex(from indexer.Version, to indexer.Version) error {
	ret := _m.Called(from, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(indexer.Version, indexer.Version) error); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetID provides a mock function with given fields:
func (_m *Indexer) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSubscriptions provides a mock function with given fields:
func (_m *Indexer) GetSubscriptions() []indexer.Subscription {
	ret := _m.Called()

	var r0 []indexer.Subscription
	if rf, ok := ret.Get(0).(func() []indexer.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]indexer.Subscription)
		}
	}

	return r0
}

// GetVersion provides a mock function with given fields:
func (_m *Indexer) GetVersion() indexer.Version {
	ret := _m.Called()

	var r0 indexer.Version
	if rf, ok := ret.Get(0).(func() indexer.Version); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(indexer.Version)
	}

	return r0
}

// Index provides a mock function with given fields: networkID, states
func (_m *Indexer) Index(networkID string, states types.StatesByID) (indexer.StateErrors, error) {
	ret := _m.Called(networkID, states)

	var r0 indexer.StateErrors
	if rf, ok := ret.Get(0).(func(string, types.StatesByID) indexer.StateErrors); ok {
		r0 = rf(networkID, states)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer.StateErrors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.StatesByID) error); ok {
		r1 = rf(networkID, states)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareReindex provides a mock function with given fields: from, to, isFirstReindex
func (_m *Indexer) PrepareReindex(from indexer.Version, to indexer.Version, isFirstReindex bool) error {
	ret := _m.Called(from, to, isFirstReindex)

	var r0 error
	if rf, ok := ret.Get(0).(func(indexer.Version, indexer.Version, bool) error); ok {
		r0 = rf(from, to, isFirstReindex)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
