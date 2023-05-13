-- Create an Event
INSERT INTO "events" (name, date, description, img, location, user_id)
VALUES ($ 1, $ 2, $ 3, $ 4, $ 5, $ 6)

-- Select an Event by id
SELECT id, name, date, description, img, location, user_id
FROM "events"
WHERE id = $ 1

-- Select all Events
SELECT id, name, date, description, img, location, user_id
FROM "events"
LIMIT $1 
OFFSET $2
