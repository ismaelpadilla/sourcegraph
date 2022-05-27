ALTER TABLE changesets
    ADD COLUMN IF NOT EXISTS detached_at timestamp with time zone;

CREATE INDEX IF NOT EXISTS
    changesets_detached_at
    ON
        changesets (detached_at);

COMMENT ON COLUMN changesets.detached_at IS 'Tracks when a Changeset is detached';