create table beer_bot
(
    id bigserial not null
        constraint beer_pk
            primary key,
    title text,
    abv text,
    expires_at timestamp
);