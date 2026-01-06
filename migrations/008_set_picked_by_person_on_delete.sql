-- +goose Up
-- +goose StatementBegin
ALTER TABLE entries DROP CONSTRAINT IF EXISTS entries_picked_by_person_id_fkey;
ALTER TABLE entries ADD CONSTRAINT entries_picked_by_person_id_fkey
    FOREIGN KEY (picked_by_person_id) REFERENCES persons(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries DROP CONSTRAINT IF EXISTS entries_picked_by_person_id_fkey;
ALTER TABLE entries ADD CONSTRAINT entries_picked_by_person_id_fkey
    FOREIGN KEY (picked_by_person_id) REFERENCES persons(id);
-- +goose StatementEnd
