CREATE TABLE IF NOT EXISTS accessibility (
    id serial,
    title text UNIQUE,
    message text,
    level smallint,
    is_replace smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT accessibility_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN accessibility.id IS 'ID';

COMMENT ON COLUMN accessibility.title IS 'チェックするアクセシビリティ項目名';

COMMENT ON COLUMN accessibility.message IS '警告文';

COMMENT ON COLUMN accessibility.level IS '警告レベル 0:チェックしない, 1:チェックする, 2:チェックするが警告だけ';

INSERT INTO accessibility (id, title, message, level, is_replace) VALUES (1, '画像代替文字', '画像に代替文字が設定されていません。アイコンの場合は空文字を設定してください。', 1, 0);
INSERT INTO accessibility (id, title, message, level, is_replace) VALUES (2, '非推奨文字', 'アクセシビリティに問題のある文字列があります。', 1, 1);