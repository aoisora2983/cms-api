CREATE TABLE IF NOT EXISTS blog_tags (
    id_blog_content integer NOT NULL,
    id_branch_blog_content integer NOT NULL,
    id_tag integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT blog_tags_pkey PRIMARY KEY (id_blog_content, id_branch_blog_content, id_tag),
    FOREIGN KEY (id_tag) REFERENCES tags(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_blog_content, id_branch_blog_content) REFERENCES blog_contents(id, id_branch) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMENT ON COLUMN blog_tags.id_blog_content IS '記事ID';

COMMENT ON COLUMN blog_tags.id_branch_blog_content IS '記事枝番';

COMMENT ON COLUMN blog_tags.id_tag IS 'タグID';