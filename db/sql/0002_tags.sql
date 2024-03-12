CREATE TABLE IF NOT EXISTS tags (
    id serial,
    name character varying(50),
    icon_path text,
    sort_order integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT tags_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN tags.id IS 'ID';

COMMENT ON COLUMN tags.name IS 'タグ名';

COMMENT ON COLUMN tags.icon_path IS 'タグアイコンのパス';

COMMENT ON COLUMN tags.sort_order IS '並び順';