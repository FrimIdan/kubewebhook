// Code generated by mockery v2.7.4. DO NOT EDIT.

package validatingmock

import (
	context "context"

	model "github.com/slok/kubewebhook/v2/pkg/model"
	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	validating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: ctx, ar, obj
func (_m *Validator) Validate(ctx context.Context, ar *model.AdmissionReview, obj v1.Object) (*validating.ValidatorResult, error) {
	ret := _m.Called(ctx, ar, obj)

	var r0 *validating.ValidatorResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.AdmissionReview, v1.Object) *validating.ValidatorResult); ok {
		r0 = rf(ctx, ar, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validating.ValidatorResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.AdmissionReview, v1.Object) error); ok {
		r1 = rf(ctx, ar, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
