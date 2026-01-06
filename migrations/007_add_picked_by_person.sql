-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries ADD COLUMN picked_by_person_id UUID REFERENCES persons(id);

-- Index for querying entries by picker
CREATE INDEX idx_entries_picked_by ON entries(picked_by_person_id) WHERE picked_by_person_id IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_entries_picked_by;
ALTER TABLE entries DROP COLUMN IF EXISTS picked_by_person_id;
-- +goose StatementEnd
