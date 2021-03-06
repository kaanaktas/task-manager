CREATE TABLE IF NOT EXISTS public.task_table
(
    id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    task_id character varying(255),
    task_name character varying(255),
    task_done boolean not null,
    CONSTRAINT uniquetaskidconstraint UNIQUE (task_id)
)
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;

ALTER TABLE public.task_table OWNER to postgres;
