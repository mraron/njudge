--
-- Replace (problemset, problem) with problem_id in submissions table
--
alter table submissions
    add problem_id integer default 0 not null;

UPDATE submissions SET problem_id = (
    SELECT id from problem_rels pr 
    WHERE pr.problemset = submissions.problemset 
    and pr.problem = submissions.problem
);

alter table submissions
    add constraint submissions_problem_rels_id_fk
        foreign key (problem_id) references public.problem_rels;

alter table submissions
    drop column problemset;

alter table submissions
    drop column problem;

--
-- Add unique constraint in problem_rels on columns (problemset, problem)
--
ALTER TABLE problem_rels
    ADD CONSTRAINT problem_rels_unique_problemset_problem UNIQUE (problemset, problem);

--
-- Modify foregin key to cascade
--
alter table forgotten_password_keys
    drop constraint forgotten_password_keys_user_fkey;

alter table forgotten_password_keys
    add constraint forgotten_password_keys_user_fkey
        foreign key (user_id) references users
            on delete cascade;

--
-- PL/PgSQL script to remove duplicates
--
BEGIN TRANSACTION;

DELETE from users 
WHERE name IN (SELECT name FROM users GROUP BY name HAVING count(*) > 1) 
and activation_key IS NOT NULL;

;DO
$$
DECLARE
   rec   record;
BEGIN
   FOR rec IN
    (SELECT u1.id as "to", u2.id as "from"
 from users u1,
      users u2
 WHERE u1.name IN (SELECT name FROM users GROUP BY name HAVING count(*) > 1)
   and u1.name = u2.name
   and u1.id < u2.id
 ORDER BY u2.id DESC) UNION (SELECT u1.id as "to", u2.id as "from"
    from users u1, users u2
    WHERE u1.email IN (SELECT email FROM users GROUP BY email HAVING count(*) > 1) and u1.email = u2.email and u1.id < u2.id ORDER BY u2.id DESC)
   LOOP
     UPDATE submissions SET user_id=rec."to" WHERE user_id=rec."from";
     DELETE FROM users WHERE id=rec."from";
   END LOOP;
END
$$;

COMMIT;

-- 
-- Add unique constraints to users
--
ALTER TABLE users ADD CONSTRAINT users_name_unique UNIQUE (name);
ALTER TABLE users ADD CONSTRAINT users_email_unique UNIQUE (email);

--
-- Add registered column
--
alter table users
    add registered timestamp with time zone default CURRENT_TIMESTAMP;

--
-- Add unique constraint on user_id in forgotten_pasword_keys
--

alter table forgotten_password_keys ADD constraint forgotten_password_keys_user_id_unique UNIQUE (user_id);

---
--- Modify problem_tags
---

BEGIN TRANSACTION;

alter table public.problem_tags
    drop constraint unique_problem_tag;

alter table public.problem_tags
    drop column id;

alter table public.problem_tags
    add primary key (problem_id, tag_id);

COMMIT;
