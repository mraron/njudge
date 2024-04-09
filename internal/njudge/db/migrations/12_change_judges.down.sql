
alter table public.judges
    drop column problem_list, drop column language_list;

alter table public.judges
    add ping text,
    add column host text,
    add column port text,
    add column state text;

update public.judges set
    host = split_part(url, ':', 1),
    port = split_part(url, ':', 2),
    state = '{}';

alter table public.judges alter column host set not null,
    alter column port set not null,
    alter column state set not null;

alter table public.judges
    drop column url ;


