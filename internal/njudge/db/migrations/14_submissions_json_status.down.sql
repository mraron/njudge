alter table public.submissions
    alter column status type text using status::text;
