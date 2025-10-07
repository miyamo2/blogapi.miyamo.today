-- name: PutTag :exec
INSERT INTO "tags" (
    "id"
    ,"name"
    ,"created_at"
    ,"updated_at"
)
VALUES (
    $1
    ,$2
    ,$3
    ,$4
)
ON CONFLICT ("id") DO UPDATE
SET "name" = EXCLUDED.name
,"updated_at" = EXCLUDED.updated_at;

-- name: CreateTempArticlesTable :exec
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

-- name: PreTagArticles :copyfrom
INSERT INTO "tmp_articles" (
    "id"
    ,"tag_id"
    ,"title"
    ,"thumbnail"
    ,"created_at"
    ,"updated_at"
) VALUES (
    $1
    ,$2
    ,$3
    ,$4
    ,$5
    ,$6
);

-- name: TagArticles :exec
WITH "inserted" AS (
    INSERT INTO "articles" (
        "id"
        ,"tag_id"
        ,"title"
        ,"thumbnail"
        ,"created_at"
        ,"updated_at"
    ) SELECT * FROM "tmp_articles"
    ON CONFLICT DO NOTHING
    RETURNING "id"
)
DELETE FROM "articles" WHERE "articles"."tag_id" = $1 AND "articles"."id" NOT IN (SELECT "id" FROM "inserted");

