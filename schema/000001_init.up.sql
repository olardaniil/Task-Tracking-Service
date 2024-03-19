
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) UNIQUE CHECK (username <> ''),
    balance INTEGER DEFAULT 0 CHECK ( balance >= 0 )
);

CREATE TABLE quests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    cost INTEGER DEFAULT 0 CHECK ( cost >= 0 )
);

CREATE TABLE quests_progress (
    user_id INTEGER,
    quest_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (quest_id) REFERENCES quests(id)
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    quest_id INTEGER,
    name VARCHAR(150),
    is_reusable BOOLEAN DEFAULT FALSE,
    cost INTEGER DEFAULT 0 CHECK ( cost >= 0 ),
    FOREIGN KEY (quest_id) REFERENCES quests(id)
);

CREATE TABLE tasks_progress (
    user_id INTEGER,
    task_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);

