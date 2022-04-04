	"github.com/sourcegraph/sourcegraph/lib/errors"
	workspace := &btypes.BatchSpecWorkspace{BatchSpecID: batchSpec.ID, RepoID: repo.ID}
	executionStore := &batchSpecWorkspaceExecutionWorkerStore{Store: workStore, observationContext: &observation.TestContext, accessTokenDeleterForTX: func(tx *Store) accessTokenHardDeleter { return tx.DatabaseDB().AccessTokens().HardDeleteByID }}
		tokenID, _, err := db.AccessTokens().CreateInternal(ctx, user.ID, []string{"user:all"}, "testing", user.ID)
		_, err = db.AccessTokens().GetByID(ctx, tokenID)
		accessTokens := database.NewMockAccessTokenStore()
		accessTokens.HardDeleteByIDFunc.SetDefaultHook(func(ctx context.Context, id int64) error {
		})

		prevDeleter := executionStore.accessTokenDeleterForTX
		executionStore.accessTokenDeleterForTX = func(tx *Store) accessTokenHardDeleter {
			return accessTokens.HardDeleteByID
		t.Cleanup(func() {
			executionStore.accessTokenDeleterForTX = prevDeleter
		})
	workspace := &btypes.BatchSpecWorkspace{BatchSpecID: batchSpec.ID, RepoID: repo.ID}
	executionStore := &batchSpecWorkspaceExecutionWorkerStore{Store: workStore, observationContext: &observation.TestContext, accessTokenDeleterForTX: func(tx *Store) accessTokenHardDeleter { return tx.DatabaseDB().AccessTokens().HardDeleteByID }}
		tokenID, _, err := db.AccessTokens().CreateInternal(ctx, user.ID, []string{"user:all"}, "testing", user.ID)
	_, err = db.AccessTokens().GetByID(ctx, tokenID)