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

-- CREATE or replace TRIGGER updated_at
--     BEFORE UPDATE
--     ON randoms
--     FOR EACH ROW
-- EXECUTE PROCEDURE
--     updated_at_column();


CREATE EXTENSION pgcrypto;
do $$
    begin
        for r in 1..10000 loop
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
