alter table public.judges
    add url text ;

update public.judges set
    url = CONCAT(host, ':', port) ;

alter table public.judges
    drop column ping, drop column host, drop column port, drop column state;

alter table public.judges
    add column problem_list text[], add column language_list text[];
