// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phase

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	connectorsv1 "k8s-connectors/connector/ycr/api/v1"
	ycrconfig "k8s-connectors/connector/ycr/pkg/config"
	"k8s-connectors/pkg/util"
)

type FinalizerRegistrar struct {
	Client client.Client
}

func (r *FinalizerRegistrar) IsUpdated(
	_ context.Context, _ logr.Logger, registry *connectorsv1.YandexContainerRegistry,
) (bool, error) {
	return util.ContainsString(registry.Finalizers, ycrconfig.FinalizerName), nil
}

func (r *FinalizerRegistrar) Update(
	ctx context.Context, log logr.Logger, registry *connectorsv1.YandexContainerRegistry,
) error {
	registry.Finalizers = append(registry.Finalizers, ycrconfig.FinalizerName)
	if err := r.Client.Update(ctx, registry); err != nil {
		return fmt.Errorf("unable to update registry status: %v", err)
	}
	log.Info("finalizer registered successfully")
	return nil
}

func (r *FinalizerRegistrar) Cleanup(
	ctx context.Context, log logr.Logger, registry *connectorsv1.YandexContainerRegistry,
) error {
	registry.Finalizers = util.RemoveString(registry.Finalizers, ycrconfig.FinalizerName)
	if err := r.Client.Update(ctx, registry); err != nil {
		return fmt.Errorf("unable to remove finalizer: %v", err)
	}

	log.Info("finalizer removed successfully")
	return nil
}