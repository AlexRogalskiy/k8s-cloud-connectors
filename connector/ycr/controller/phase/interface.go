// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phase

import (
	"context"

	"github.com/go-logr/logr"

	connectorsv1 "k8s-connectors/connector/ycr/api/v1"
)

type YandexContainerRegistryPhase interface {
	IsUpdated(context.Context, logr.Logger, *connectorsv1.YandexContainerRegistry) (bool, error)
	Update(context.Context, logr.Logger, *connectorsv1.YandexContainerRegistry) error
	Cleanup(context.Context, logr.Logger, *connectorsv1.YandexContainerRegistry) error
}