CREATE TABLE
    IF NOT EXISTS bank (
        id bigserial PRIMARY KEY,
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW(),
            first_name text NOT NULL,
            last_name text NOT NULL,
            number bigserial NOT NULL,
            balance bigserial not NULL,
            version integer NOT NULL DEFAULT 1
    );