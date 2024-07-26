CREATE TABLE public.tasks (
	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	start_date date NULL,
	end_date date NULL,
	priority int4 NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (id)
);
