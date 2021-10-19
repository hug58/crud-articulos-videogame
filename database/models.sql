

CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);


CREATE TABLE IF NOT EXISTS articles (
    id serial NOT NULL,
    name VARCHAR(150) NOT NULL,
    price INT default 0, 
    description text NOT NULL,
    user_id int NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_notes PRIMARY KEY(id),
    CONSTRAINT fk_articles_users FOREIGN KEY(user_id) REFERENCES users(id)
);