-- +goose Up
-- +goose StatementBegin
BEGIN;

DROP TABLE IF EXISTS randoms;
CREATE TABLE randoms
(
    id             bigint generated always as identity,
    name           text                     not null,

    primary key (id)
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;
do $$
    begin
        for r in 1..100 loop
            insert into randoms(name) values(gen_random_uuid());
        end loop;
    end;
$$;

COMMIT;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS randoms;
-- +goose StatementEnd
