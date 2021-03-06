// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pubsub_test

import (
	"context"
	"sync"
	"testing"

	"github.com/facebookincubator/symphony/pkg/pubsub"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gcpubsub "gocloud.dev/pubsub"
	"gocloud.dev/pubsub/mempubsub"
)

func TestTopicEmitter(t *testing.T) {
	ctx := context.Background()
	url := mempubsub.Scheme + "://" + uuid.New().String()
	emitter, err := pubsub.NewTopicEmitter(ctx, url)
	require.NoError(t, err)
	assert.Implements(t, (*pubsub.Emitter)(nil), emitter)
	subscription, err := gcpubsub.OpenSubscription(ctx, url)
	require.NoError(t, err)
	defer subscription.Shutdown(ctx)

	const tenant = "test-tenant"
	body := []byte("test-body")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		msg, err := subscription.Receive(ctx)
		require.NoError(t, err)
		defer msg.Ack()
		assert.Equal(t, tenant, msg.Metadata[pubsub.TenantHeader])
		assert.Equal(t, t.Name(), msg.Metadata[pubsub.NameHeader])
		assert.Equal(t, body, msg.Body)
	}()
	defer wg.Wait()

	err = emitter.Emit(ctx, tenant, t.Name(), body)
	require.NoError(t, err)
	err = emitter.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestNopEmitter(t *testing.T) {
	emitter := pubsub.NewNopEmitter()
	err := emitter.Emit(context.Background(), "", "", nil)
	assert.NoError(t, err)
}
