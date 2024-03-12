CREATE TABLE IF NOT EXISTS correct_words (
    id serial,
    id_accessibility int,
    word_from text UNIQUE,
    word_to text,
    level smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT correct_words_pkey PRIMARY KEY (id),
    FOREIGN KEY (id_accessibility) REFERENCES accessibility(id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMENT ON COLUMN correct_words.id IS 'ID';

COMMENT ON COLUMN correct_words.accessibility_id IS 'アクセシビリティID';

COMMENT ON COLUMN correct_words.word_from IS '修正対象の単語';

COMMENT ON COLUMN correct_words.word_to IS '修正対象の置換推奨単語';

COMMENT ON COLUMN correct_words.level IS '警告レベル 0:チェックしない, 1:禁止, 2:警告';

INSERT INTO correct_words(id_accessibility, word_from, word_to, level) VALUES (2, 'ｱ', 'あ', 1); 