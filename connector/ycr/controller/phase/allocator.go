// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phase

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"

	connectorsv1 "k8s-connectors/connector/ycr/api/v1"
	"k8s-connectors/connector/ycr/controller/adapter"
	ycrutils "k8s-connectors/connector/ycr/pkg/util"
	"k8s-connectors/pkg/config"
)

type Allocator struct {
	Sdk       adapter.YandexContainerRegistryAdapter
	ClusterID string
}

func (r *Allocator) IsUpdated(ctx context.Context, _ logr.Logger, object *connectorsv1.YandexContainerRegistry) (
	bool, error,
) {
	res, err := ycrutils.GetRegistry(
		ctx, object.Status.ID, object.Spec.FolderID, object.ObjectMeta.Name, r.ClusterID, r.Sdk,
	)
	return res != nil, err
}

func (r *Allocator) Update(ctx context.Context, log logr.Logger, object *connectorsv1.YandexContainerRegistry) error {
	if _, err := r.Sdk.Create(
		ctx, &containerregistry.CreateRegistryRequest{
			FolderId: object.Spec.FolderID,
			Name:     object.Spec.Name,
			Labels: map[string]string{
				config.CloudClusterLabel: r.ClusterID,
				config.CloudNameLabel:    object.Name,
			},
		},
	); err != nil {
		return err
	}
	log.Info("resource allocated successfully")
	return nil
}

func (r *Allocator) Cleanup(ctx context.Context, log logr.Logger, object *connectorsv1.YandexContainerRegistry) error {
	ycr, err := ycrutils.GetRegistry(
		ctx, object.Status.ID, object.Spec.FolderID, object.ObjectMeta.Name, r.ClusterID, r.Sdk,
	)
	if err != nil {
		return err
	}
	if ycr == nil {
		log.Info("registry deleted externally")
		return nil
	}

	if err := r.Sdk.Delete(ctx, ycr.Id); err != nil {
		return err
	}
	log.Info("registry deleted successfully")
	return nil
}