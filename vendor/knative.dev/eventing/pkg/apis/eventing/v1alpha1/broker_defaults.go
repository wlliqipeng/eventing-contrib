/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	eventingduckv1alpha1 "knative.dev/eventing/pkg/apis/duck/v1alpha1"
	"knative.dev/pkg/apis"
)

func (b *Broker) SetDefaults(ctx context.Context) {
	withNS := apis.WithinParent(ctx, b.ObjectMeta)
	b.Spec.SetDefaults(withNS)
}

func (bs *BrokerSpec) SetDefaults(ctx context.Context) {
	if bs.Config == nil {
		// If we haven't configured the new channelTemplate,
		// then set the default channel to the new channelTemplate.
		if bs.ChannelTemplate == nil {
			// The singleton may not have been set, if so ignore it and validation will reject the Broker.
			if cd := eventingduckv1alpha1.ChannelDefaulterSingleton; cd != nil {
				channelTemplate := cd.GetDefault(apis.ParentMeta(ctx).Namespace)
				bs.ChannelTemplate = channelTemplate
			}
		}
	} else {
		if bs.Config.Namespace == "" {
			bs.Config.Namespace = apis.ParentMeta(ctx).Namespace
		}
	}
}
