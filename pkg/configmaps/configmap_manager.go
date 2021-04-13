// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package configmaps

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rtcl "sigs.k8s.io/controller-runtime/pkg/client"
)

func cmapName(resourceName, kind string) string {
	return kind + "-" + resourceName + "-" + "configmap"
}

func Exists(ctx context.Context, client *rtcl.Client, resourceName string, namespace string, kind string) (bool, error) {
	cmapName := cmapName(resourceName, kind)

	var cmapObj v1.ConfigMap
	err := (*client).Get(ctx, rtcl.ObjectKey{Namespace: namespace, Name: cmapName}, &cmapObj)
	if errors.IsNotFound(err) {
		return true, nil
	}
	return false, err
}

// TODO (covariance) autogenerate for all resource types
func Put(ctx context.Context, client *rtcl.Client, resourceName string, namespace string, kind string, data map[string]string) error {
	cmapName := cmapName(resourceName, kind)

	var cmapObj v1.ConfigMap
	err := (*client).Get(ctx, rtcl.ObjectKey{Namespace: namespace, Name: cmapName}, &cmapObj)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	newState := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmapName,
			Namespace: namespace,
			Labels: map[string]string{
				"kind" : kind,
			},
		},
		Data: data,
	}
	if errors.IsNotFound(err) {
		if err := (*client).Create(ctx, &newState); err != nil {
			return err
		}
	} else {
		if err := (*client).Update(ctx, &newState); err != nil {
			return err
		}
	}
	return nil
}

func Remove(ctx context.Context, client rtcl.Client, resourceName string, namespace string, kind string) error {
	cmapName := cmapName(resourceName, kind)

	var cmapObj v1.ConfigMap
	err := client.Get(ctx, rtcl.ObjectKey{Namespace: namespace, Name: cmapName}, &cmapObj)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	if errors.IsNotFound(err) {
		return nil
	}

	if err := client.Delete(ctx, &cmapObj); err != nil {
		return err
	}

	return nil
}