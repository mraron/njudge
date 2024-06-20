alter table public.submissions
    alter column status type jsonb using status::jsonb;
