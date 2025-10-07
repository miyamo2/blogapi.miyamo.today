CREATE TABLE IF NOT EXISTS tags (
    id VARCHAR(144),
    name VARCHAR(35) NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS articles (
    id VARCHAR(26),
    tag_id VARCHAR(144),
    title VARCHAR(255) NOT NULL,
    thumbnail VARCHAR(524271),
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    PRIMARY KEY (id, tag_id)
);

CREATE TEMP TABLE tmp_articles (
    id VARCHAR(26),
    tag_id VARCHAR(144),
    title VARCHAR(255) NOT NULL,
    thumbnail VARCHAR(524271),
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    PRIMARY KEY (id, tag_id)
) ON COMMIT PRESERVE ROWS;