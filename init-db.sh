#!/bin/bash

echo "⏳ Waiting for the database to start..."
sleep 5

echo "🚀 Running PostgreSQL commands inside the container..."
docker exec -it golabs_db psql -U postgres -d mydb <<EOF
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL
);

INSERT INTO users (name, email) VALUES ('Admin', 'admin@example.com') ON CONFLICT DO NOTHING;

SELECT * FROM users;
EOF

echo "✅ Database setup completed!"