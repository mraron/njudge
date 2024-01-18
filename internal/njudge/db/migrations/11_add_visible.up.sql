alter table public.problem_rels
    add visible boolean default true not null;

alter table public.problem_categories
    add visible boolean default true not null;
