-- name: GetByID :one
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT * FROM "articles" WHERE "articles"."id" = $1) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "a"."id" = "t"."article_id"
GROUP BY "a"."id";

-- name: ListAfter :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT * FROM "articles" ORDER BY "articles"."id") AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id";

-- name: ListAfterWithLimit :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT * FROM "articles" ORDER BY "articles"."id" LIMIT $1) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id";

-- name: ListAfterWithLimitAndCursor :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT *
      FROM "articles"
      WHERE EXISTS(SELECT id
                   FROM "articles"
                   WHERE "articles"."id" = $1)
        AND "articles"."id" > $1
      ORDER BY "articles"."id" LIMIT $2) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id";

-- name: ListBefore :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT * FROM "articles" ORDER BY "articles"."id" DESC) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id" DESC;

-- name: ListBeforeWithLimit :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT * FROM "articles" ORDER BY "articles"."id" DESC LIMIT $1) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id" DESC;

-- name: ListBeforeWithLimitAndCursor :many
SELECT "a".*,
       CAST(
               COALESCE(
                       jsonb_agg(
                               json_build_object('id', t.id, 'name', t.name)
                       ) FILTER(WHERE t.id IS NOT NULL), '[]' ::json
               ) AS json
       ) AS "tags"
FROM (SELECT *
      FROM "articles"
      WHERE EXISTS(SELECT id
                   FROM "articles"
                   WHERE "articles"."id" = $1)
        AND "articles"."id" < $1
      ORDER BY "articles"."id" DESC LIMIT $2) AS "a"
         LEFT OUTER JOIN
     "tags" AS "t"
     ON
         "t"."id" = "a"."tag_id"
GROUP BY "a"."id"
ORDER BY "a"."id" DESC;