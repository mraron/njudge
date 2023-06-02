create table if not exists partials
(
    name text not null constraint partials_pk primary key,
    html text not null 
);
