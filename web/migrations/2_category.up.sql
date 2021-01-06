create table if not exists problem_categories
(
    id serial not null
        constraint problem_categories_pk
            primary key,
    name text not null,
    parent_id int
        constraint problem_categories_problem_categories_id_fk
            references problem_categories
);


alter table problem_rels
    add category_id int;

alter table problem_rels
    add constraint problem_rels_problem_categories_id_fk
        foreign key (category_id) references problem_categories;

