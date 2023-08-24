create table segment
(
    id   integer not null
        constraint segment_pk
            primary key,
    name varchar not null
        unique
);

create table user_segment
(
    id         integer not null
        constraint user_segment_pk
            primary key,
    user_id    integer not null,
    segment_id integer not null
        constraint user_segment_segment_id_fk
            references segment,
    expires    date,
    constraint user_segment_pk2
        unique (user_id, segment_id, expires)
);
