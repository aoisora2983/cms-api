CREATE TABLE IF NOT EXISTS portfolios (
    id serial,
    title text NOT NULL,
    description text,
    thumbnail text,
    detail_url text,
    release_time timestamp with time zone,
    status integer,
    sort_order integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT portfolios_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN portfolios.id IS 'ID';

COMMENT ON COLUMN portfolios.title IS 'タイトル';

COMMENT ON COLUMN portfolios.description IS '概要';

COMMENT ON COLUMN portfolios.thumbnail IS 'サムネイルパス';

COMMENT ON COLUMN portfolios.detail_url IS '詳細ページへのURL';

COMMENT ON COLUMN portfolios.release_time IS 'リリース日';

COMMENT ON COLUMN portfolios.status IS '状態';
