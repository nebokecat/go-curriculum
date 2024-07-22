CREATE TABLE public.tasks (
	id int8 GENERATED ALWAYS AS IDENTITY NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	start_date date NULL,
	end_date varchar NULL,
	priority int4 NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (id)
);
