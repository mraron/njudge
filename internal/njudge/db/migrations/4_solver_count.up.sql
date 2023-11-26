alter table problem_rels
    add solver_count integer not null default 0;
update problem_rels
set solver_count = (SELECT COUNT(distinct user_id)
                    from submissions
                    where problemset = problem_rels.problemset
                      and problem = problem_rels.problem and verdict = 0);

