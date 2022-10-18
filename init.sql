CREATE TABLE category (
    id serial primary key,
    title varchar(50) NOT NULL
);

CREATE TABLE product (
    id serial primary key,
    title varchar(50) NOT NULL,
    description text NOT NULL,
    category_id integer REFERENCES category (id)
);

