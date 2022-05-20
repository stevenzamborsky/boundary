package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/boundary/internal/daemon/controller"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/hashicorp/boundary/internal/server"
	"github.com/hashicorp/boundary/internal/types/scope"
	"github.com/stretchr/testify/require"
)

func TestTagUpdatingListing(t *testing.T) {
	require := require.New(t)

	wrapper := db.TestWrapper(t)
	tc := controller.NewTestController(t, &controller.TestControllerOpts{
		RecoveryKms: wrapper,
	})
	defer tc.Shutdown()

	repo := tc.ServersRepo()
	srv := server.NewWorker(scope.Global.String(),
		server.WithPublicId("test_worker_1"),
		server.WithAddress("127.0.0.1"),
		server.WithWorkerTags(
			&server.Tag{
				Key:   "tag1",
				Value: "value1",
			},
			&server.Tag{
				Key:   "tag1",
				Value: "value2",
			}))

	_, _, err := repo.UpsertWorker(tc.Context(), srv, server.WithUpdateTags(true))
	require.NoError(err)

	srv = server.NewWorker(scope.Global.String(),
		server.WithPublicId("test_worker_2"),
		server.WithAddress("127.0.0.1"),
		server.WithWorkerTags(
			&server.Tag{
				Key:   "tag2",
				Value: "value1",
			},
			&server.Tag{
				Key:   "tag2",
				Value: "value2",
			}))
	_, _, err = repo.UpsertWorker(tc.Context(), srv, server.WithUpdateTags(true))
	require.NoError(err)

	tags, err := repo.ListTagsForWorkers(tc.Context(), []string{"test_worker_1", "test_worker_2"})
	require.NoError(err)

	// Base case
	exp := map[string][]*server.Tag{
		"test_worker_1": {
			{
				Key:   "tag1",
				Value: "value1",
			}, {
				Key:   "tag1",
				Value: "value2",
			},
		},
		"test_worker_2": {
			{
				Key:   "tag2",
				Value: "value1",
			},
			{
				Key:   "tag2",
				Value: "value2",
			},
		},
	}
	require.Equal(exp, tags)

	// Update without saying to update tags
	srv = server.NewWorker(scope.Global.String(),
		server.WithPublicId("test_worker_2"),
		server.WithAddress("192.168.1.1"),
		server.WithWorkerTags(
			&server.Tag{
				Key:   "tag22",
				Value: "value21",
			},
			&server.Tag{
				Key:   "tag22",
				Value: "value22",
			}))
	_, _, err = repo.UpsertWorker(tc.Context(), srv)
	require.NoError(err)
	tags, err = repo.ListTagsForWorkers(tc.Context(), []string{"test_worker_1", "test_worker_2"})
	require.NoError(err)
	require.Equal(exp, tags)

	// Update tags and test again
	_, _, err = repo.UpsertWorker(tc.Context(), srv, server.WithUpdateTags(true))
	require.NoError(err)
	tags, err = repo.ListTagsForWorkers(tc.Context(), []string{"test_worker_1", "test_worker_2"})
	require.NoError(err)
	require.NotEqual(exp, tags)
	// Update and try again
	exp["test_worker_2"] = []*server.Tag{
		{
			Key:   "tag22",
			Value: "value21",
		},
		{
			Key:   "tag22",
			Value: "value22",
		},
	}
	require.Equal(exp, tags)
}

func TestListWorkersWithLiveness(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, wrapper)
	serversRepo, err := server.NewRepository(rw, rw, kms)
	require.NoError(err)
	ctx := context.Background()

	newWorker := func(publicId string) *server.Worker {
		result := server.NewWorker(scope.Global.String(),
			server.WithPublicId(publicId),
			server.WithAddress("127.0.0.1"))
		_, rowsUpdated, err := serversRepo.UpsertWorker(ctx, result)
		require.NoError(err)
		require.Equal(1, rowsUpdated)

		return result
	}

	server1 := newWorker("test_worker_1")
	server2 := newWorker("test_worker_2")
	server3 := newWorker("test_worker_3")

	// Sleep the default liveness time (15sec currently) +1s
	time.Sleep(time.Second * 16)

	// Push an upsert to the first worker so that its status has been
	// updated.
	_, rowsUpdated, err := serversRepo.UpsertWorker(ctx, server1)
	require.NoError(err)
	require.Equal(1, rowsUpdated)

	requireIds := func(expected []string, actual []*server.Worker) {
		require.Len(expected, len(actual))
		want := make(map[string]struct{})
		for _, v := range expected {
			want[v] = struct{}{}
		}

		got := make(map[string]struct{})
		for _, v := range actual {
			got[v.PublicId] = struct{}{}
		}

		require.Equal(want, got)
	}

	// Default liveness, should only list 1
	result, err := serversRepo.ListWorkers(ctx)
	require.NoError(err)
	require.Len(result, 1)
	requireIds([]string{server1.PublicId}, result)

	// Upsert second server.
	_, rowsUpdated, err = serversRepo.UpsertWorker(ctx, server2)
	require.NoError(err)
	require.Equal(1, rowsUpdated)

	// Static liveness. Should get two, so long as this did not take
	// more than 5s to execute.
	result, err = serversRepo.ListWorkers(ctx, server.WithLiveness(time.Second*5))
	require.NoError(err)
	require.Len(result, 2)
	requireIds([]string{server1.PublicId, server2.PublicId}, result)

	// Liveness disabled, should get all three workers.
	result, err = serversRepo.ListWorkers(ctx, server.WithLiveness(-1))
	require.NoError(err)
	require.Len(result, 3)
	requireIds([]string{server1.PublicId, server2.PublicId, server3.PublicId}, result)
}
