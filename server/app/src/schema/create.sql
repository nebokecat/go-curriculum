-- public.tasks definition

-- Drop table

-- DROP TABLE public.tasks;

CREATE TABLE public.tasks (
	"name" varchar NULL,
	start_date date NULL,
	end_date date NULL,
	priority int4 NULL,
	id int8 NOT NULL,
	description varchar NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (id)
);
