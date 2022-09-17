create table if not exists tags (
    id serial not null constraint tags_pkey primary key,
    name text not null
);

insert into tags (name) values 
('implementáció'),
('matek'),
('mohó'),
('rekurzív kiszámítás'),
('dp'),
('adatszerkezetek'),
('bruteforce'),
('konstruktív'),
('gráfok'),
('rendezés'),
('bináris keresés'),
('gráfbejárás'),
('fák'),
('sztringek'),
('kombinatorika'),
('geometria'),
('bitmaszk'),
('két mutató'),
('oszd meg és uralkodj'),
('legröviebb utak'),
('játékok'),
('párosítások');

create table if not exists problem_tags (
    id serial not null constraint problem_tags_pkey primary key,
    problem_id int NOT NULL constraint problem_tags_problem_id_fk references problem_rels (id),
    tag_id int NOT NULL constraint problem_tags_tag_id_fk references tags (id),
    user_id int NOT NULL constraint problem_tags_user_id_fk references users (id),
    added timestamp not null,
    constraint unique_problem_tag unique(problem_id, tag_id)
);