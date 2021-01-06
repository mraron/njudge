alter table problem_rels
    drop constraint if exists problem_rels_problem_categories_id_fk;

alter table problem_rels
    drop if exists category_id;

drop table if exists problem_categories;
