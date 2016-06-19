/* SQLEditor (Postgres)*/

DROP TABLE IF EXISTS user_shampoo;
DROP TABLE IF EXISTS public.user;
DROP TABLE IF EXISTS shampoo;

CREATE TABLE shampoo
(
id SERIAL NOT NULL UNIQUE ,
brand VARCHAR(255) UNIQUE ,
created_on TIMESTAMP,
updated_on TIMESTAMP,
deleted_on TIMESTAMP,
PRIMARY KEY (id)
);

CREATE TABLE "user"
(
id SERIAL NOT NULL UNIQUE ,
email VARCHAR(255) UNIQUE ,
created_on TIMESTAMP,
updated_on TIMESTAMP,
deleted_on TIMESTAMP,
PRIMARY KEY (id)
);

CREATE TABLE user_shampoo
(
id SERIAL NOT NULL UNIQUE ,
created_on TIMESTAMP,
updated_on TIMESTAMP,
deleted_on TIMESTAMP,
shampoo_id BIGSERIAL,
user_id BIGSERIAL,
PRIMARY KEY (id)
);

CREATE INDEX shampoo_name_idx ON shampoo(brand);

CREATE INDEX user_email_idx ON "user"(email);

ALTER TABLE user_shampoo ADD FOREIGN KEY (shampoo_id) REFERENCES shampoo (id);

ALTER TABLE user_shampoo ADD FOREIGN KEY (user_id) REFERENCES "user" (id);
