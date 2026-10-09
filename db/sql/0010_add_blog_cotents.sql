ALTER TABLE
    blog_contents
ADD
    COLUMN id_sub_site smallint DEFAULT 0;

COMMENT ON COLUMN blog_contents.id_sub_site IS '紐付くサブサイトID';