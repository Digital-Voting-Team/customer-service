-- +migrate Up
-- Table: public.address

CREATE TABLE IF NOT EXISTS public.address
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    building_number integer,
    street character varying(45),
    city character varying(45),
    district character varying(45),
    region character varying(45),
    postal_code integer,
    CONSTRAINT address_id PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.address
    OWNER to postgres;

INSERT INTO public.address(
    building_number, street, district, city, region, postal_code)
VALUES (1, 'polska', 'polska', 'polska', 'polska', 58000);

-- Table: public.person

CREATE TABLE IF NOT EXISTS public.person
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(45),
    phone character varying(30),
    email character varying(45),
    address_id integer,
    CONSTRAINT person_id PRIMARY KEY (id),
    CONSTRAINT address_id FOREIGN KEY (address_id)
        REFERENCES public.address (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.person
    OWNER to postgres;

INSERT INTO public.person(
    name, phone, email, address_id)
VALUES ('Derek', '+380435815532', 'your.funny.email@lol.tik', 1);

-- Table: public.customer

CREATE TABLE IF NOT EXISTS public.customer
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    person_id integer,
    registration_date timestamp,
    CONSTRAINT customer_id PRIMARY KEY (id),
    CONSTRAINT person_id FOREIGN KEY (person_id)
        REFERENCES public.person (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.customer
    OWNER to postgres;

INSERT INTO public.customer(
    person_id, registration_date)
VALUES (1, '1996-12-02');


-- +migrate Down
DROP TABLE IF EXISTS public.customer;
DROP TABLE IF EXISTS public.person;
DROP TABLE IF EXISTS public.address;
