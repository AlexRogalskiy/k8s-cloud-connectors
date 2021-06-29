// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package webhook

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
)

type Mutator interface {
	Mutate(ctx context.Context, log logr.Logger, obj runtime.Object) (runtime.Object, error)
}

type ValidationError struct {
	error
}

func NewValidationError(inner error) ValidationError {
	return ValidationError{inner}
}

type Validator interface {
	ValidateCreation(ctx context.Context, log logr.Logger, obj runtime.Object) error
	ValidateUpdate(ctx context.Context, log logr.Logger, current, old runtime.Object) error
	ValidateDeletion(ctx context.Context, log logr.Logger, obj runtime.Object) error
}
