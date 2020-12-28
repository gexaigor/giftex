CREATE TABLE IF NOT EXISTS usr (
    id serial PRIMARY KEY,
    login VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    role VARCHAR (20) NOT NULL,
    created_on TIMESTAMP NOT NULL
);

CREATE INDEX ON "usr" ("login");

CREATE INDEX ON "usr" ("email");

CREATE TABLE IF NOT EXISTS company (
    id serial PRIMARY KEY,
    usr_id int NOT NULL, 
    bin VARCHAR (12) UNIQUE,
    name VARCHAR UNIQUE,
    address VARCHAR (255),
    CONSTRAINT fk_usr
      FOREIGN KEY(usr_id) 
	  REFERENCES usr(id) ON DELETE CASCADE
);

CREATE INDEX ON "company" ("bin");

CREATE INDEX ON "company" ("name");

CREATE INDEX ON "company" ("usr_id");

CREATE TABLE IF NOT EXISTS company_levels (
    id serial PRIMARY KEY,
    company_id int NOT NULL,
    experience int NOT NULL,
    level int NOT NULL,
    description VARCHAR NOT NULL,
    CONSTRAINT fk_company
      FOREIGN KEY(company_id) 
	  REFERENCES company(id) ON DELETE CASCADE
);

CREATE INDEX ON "company_levels" ("company_id");

ALTER TABLE company_levels
ADD CONSTRAINT unique_company_level UNIQUE (company_id, level);

CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    usr_id int NOT NULL,
    company_id int NOT NULL,
    experience int NOT NULL,
    CONSTRAINT fk_usr
      FOREIGN KEY(usr_id) 
	  REFERENCES usr(id) ON DELETE CASCADE,
    CONSTRAINT fk_company
      FOREIGN KEY(company_id) 
	  REFERENCES company(id) ON DELETE CASCADE
);

CREATE INDEX ON "transactions" ("usr_id");

CREATE INDEX ON "transactions" ("company_id");