-- +goose Up
-- SQL in this section is executed when the migration is applied.

--If creating the extension fails because you need to be superuser, you can do this:
--su postgres
--psql
--alter role user_name superuser;
--#then create the extension as the user in a different screen
--alter role user_name nosuperuser;
CREATE TYPE loglevel AS ENUM('debug', 'info', 'warning', 'error', 'fatal', 'panic');

CREATE TABLE log (
    id VARCHAR(150) NOT NULL,
    msg VARCHAR(1024) NOT NULL,
    level loglevel NOT NULL
);

ALTER TABLE log ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE log ADD PRIMARY KEY (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE log PURGE;
DROP TYPE loglevel;
