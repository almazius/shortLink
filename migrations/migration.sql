CREATE TABLE links (
    id bigserial primary key,
    link text,
    shortLink text,
    creationtime timestamp with time zone
);