// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../../vendor/go.opentelemetry.io/otel/sdk/trace/span_exporter.go

// Package myoteltest is a generated GoMock package.
package myoteltest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	trace "go.opentelemetry.io/otel/sdk/trace"
)

// MockSpanExporter is a mock of SpanExporter interface.
type MockSpanExporter struct {
	ctrl     *gomock.Controller
	recorder *MockSpanExporterMockRecorder
}

// MockSpanExporterMockRecorder is the mock recorder for MockSpanExporter.
type MockSpanExporterMockRecorder struct {
	mock *MockSpanExporter
}

// NewMockSpanExporter creates a new mock instance.
func NewMockSpanExporter(ctrl *gomock.Controller) *MockSpanExporter {
	mock := &MockSpanExporter{ctrl: ctrl}
	mock.recorder = &MockSpanExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpanExporter) EXPECT() *MockSpanExporterMockRecorder {
	return m.recorder
}

// ExportSpans mocks base method.
func (m *MockSpanExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSpans", ctx, spans)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExportSpans indicates an expected call of ExportSpans.
func (mr *MockSpanExporterMockRecorder) ExportSpans(ctx, spans interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSpans", reflect.TypeOf((*MockSpanExporter)(nil).ExportSpans), ctx, spans)
}

// Shutdown mocks base method.
func (m *MockSpanExporter) Shutdown(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockSpanExporterMockRecorder) Shutdown(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockSpanExporter)(nil).Shutdown), ctx)
}