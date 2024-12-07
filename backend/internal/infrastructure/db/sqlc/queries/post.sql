-- name: FuzzySearchPosts :many
SELECT *
FROM posts
WHERE title % 'search_term' OR content % 'search_term'
ORDER BY
  CASE
    WHEN title % 'search_term' THEN similarity(title, 'search_term')
    ELSE 0
  END DESC,
  CASE
    WHEN content % 'search_term' THEN similarity(content, 'search_term')
    ELSE 0
  END DESC
LIMIT 10;