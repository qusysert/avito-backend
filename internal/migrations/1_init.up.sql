create table segment
(
    id   serial
        constraint segment_pk
            primary key,
    name varchar not null
        unique
);

create table user_segment
(
    id         serial
        constraint user_segment_pk
            primary key,
    user_id    integer not null,
    segment_id integer not null
        constraint user_segment_segment_id_fk
            references segment,
    expires    timestamp,
    constraint user_segment_pk2
        unique (user_id, segment_id)
);

create index user_segment_user_id_index
    on user_segment (user_id);

