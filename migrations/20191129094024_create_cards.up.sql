create table card
(
    id      bigserial  not null primary key,
    word    varchar not null,
    meaning varchar
);

alter table card
    owner to postgres;

create unique index card_id_uindex
    on card (id);

create unique index card_word_uindex
    on card (word);

