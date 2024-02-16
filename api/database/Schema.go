package database


// ALTER TABLE messages ADD image_url text;
var schema string = `
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS user_chats;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS chats;
CREATE TABLE IF NOT EXISTS users (
    id              SERIAL NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,    
    username        VARCHAR(20) UNIQUE,
    password        VARCHAR(200),
    is_anonymous    BOOLEAN,
    profile_picture text,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS chats(
    id         SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL, 
    chat_name  VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS messages (
    id              SERIAL NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL, 
    receiver        VARCHAR(20),
    username        VARCHAR(20) NOT NULL,
    message_content TEXT,
    image_url       TEXT,
    message_date    TEXT,
    chat_id         INT,
    PRIMARY KEY (id),
    FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,
    FOREIGN KEY(receiver) REFERENCES users(username) ON DELETE CASCADE,
    FOREIGN KEY(chat_id)  REFERENCES chats(id)       ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_chats(
    id         SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL, 
    chat_id    INT,
    username   VARCHAR(20),
    PRIMARY KEY (id),
    FOREIGN KEY(chat_id)  REFERENCES chats(id)       ON DELETE CASCADE,
    FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE
);
`
func GetSchema() string{
	return schema
}