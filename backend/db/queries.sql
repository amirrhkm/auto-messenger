-- name: GetScheduledMessages :many
SELECT id, phone_number, content
FROM messages
WHERE scheduled_at <= NOW()
AND status = 'pending';

-- name: UpdateMessageStatus :exec
UPDATE messages
SET status = $2
WHERE id = $1; 