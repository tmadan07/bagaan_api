-- liquibase formated sql
--changeset SandyisGay:userCreation runInTransaction:false runAlways:true validCheckSum:any
Create TABLE public.users
(
    id serial PRIMARY KEY,
    username character varying NOT NULL UNIQUE,
    password character varying NOT NULL,
    auth_code character varying(6),
    full_name character varying  NOT NULL,
    password_verification_token character varying(6),
    email character varying  NOT NULL,
    password_changed_date timestamp without time zone  NOT NULL DEFAULT now(),
    updated_date timestamp without time zone,
    created_date timestamp without time zone  NOT NULL DEFAULT now()

);
-- rollback  drop table public.users