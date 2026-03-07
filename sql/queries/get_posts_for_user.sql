-- name: GetPostsForUser :many
SELECT
    posts.*,
    feeds.name AS feed_name
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id 
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1

ORDER BY published_at DESC
LIMIT $2;
