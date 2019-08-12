// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	"github.com/stretchr/testify/mock"
	"github.com/uber/cadence/client/admin"
	"github.com/uber/cadence/client/frontend"
	"github.com/uber/cadence/client/history"
	"github.com/uber/cadence/client/matching"
)

// MockClientBean is an autogenerated mock type for the MockClientBean type
type MockClientBean struct {
	mock.Mock
}

var _ Bean = (*MockClientBean)(nil)

// GetHistoryClient provides a mock function with given fields:
func (_m *MockClientBean) GetHistoryClient() history.Client {
	ret := _m.Called()

	var r0 history.Client
	if rf, ok := ret.Get(0).(func() history.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(history.Client)
		}
	}

	return r0
}

// GetMatchingClient provides a mock function with given fields: domainIDToName
func (_m *MockClientBean) GetMatchingClient(domainIDToName DomainIDToNameFunc) (matching.Client, error) {
	ret := _m.Called(domainIDToName)

	var r0 matching.Client
	if rf, ok := ret.Get(0).(func(DomainIDToNameFunc) matching.Client); ok {
		r0 = rf(domainIDToName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(matching.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(DomainIDToNameFunc) error); ok {
		r1 = rf(domainIDToName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFrontendClient provides a mock function with given fields:
func (_m *MockClientBean) GetFrontendClient() frontend.Client {
	ret := _m.Called()

	var r0 frontend.Client
	if rf, ok := ret.Get(0).(func() frontend.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(frontend.Client)
		}
	}

	return r0
}

// GetRemoteAdminClient provides a mock function with given fields: _a0
func (_m *MockClientBean) GetRemoteAdminClient(_a0 string) admin.Client {
	ret := _m.Called(_a0)

	var r0 admin.Client
	if rf, ok := ret.Get(0).(func(string) admin.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(admin.Client)
		}
	}

	return r0
}

// GetRemoteFrontendClient provides a mock function with given fields: _a0
func (_m *MockClientBean) GetRemoteFrontendClient(_a0 string) frontend.Client {
	ret := _m.Called(_a0)

	var r0 frontend.Client
	if rf, ok := ret.Get(0).(func(string) frontend.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(frontend.Client)
		}
	}

	return r0
}
