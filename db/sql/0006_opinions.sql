CREATE TABLE IF NOT EXISTS opinions (
    id serial,
    name text NOT NULL,
    email text NOT NULL,
    content text,
    ip CIDR NOT NULL,
    send_time timestamp with time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT opinions_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN opinions.id IS 'ID';

COMMENT ON COLUMN opinions.name IS '名前';

COMMENT ON COLUMN opinions.email IS 'メールアドレス';

COMMENT ON COLUMN opinions.content IS '内容';

COMMENT ON COLUMN opinions.ip IS 'IP';

COMMENT ON COLUMN opinions.send_time IS '問合せ日時';