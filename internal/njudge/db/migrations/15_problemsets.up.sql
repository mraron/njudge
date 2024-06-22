CREATE TYPE code_visibility AS ENUM ('private', 'solved', 'public');

CREATE TABLE problemsets (
    name text unique primary key,
    code_visibility code_visibility not null
);

INSERT INTO public.problemsets (name, code_visibility)
VALUES ('main'::text, 'public'::code_visibility);

INSERT INTO public.problemsets (name, code_visibility)
VALUES ('tutorial'::text, 'public'::code_visibility);


alter table public.problem_rels
    add constraint problem_rels_problemsets_name_fk
        foreign key (problemset) references public.problemsets;

