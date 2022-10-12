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
    postal_code character varying(45),
    CONSTRAINT address_id PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.address
    OWNER to postgres;

INSERT INTO public.address(
    building_number, street, district, city, region, postal_code)
VALUES (1, 'polska', 'polska', 'polska', 'polska', '58000');

-- Table: public.person

CREATE TABLE IF NOT EXISTS public.person
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(45) NOT NULL,
    phone character varying(30),
    email character varying(45),
    birthday timestamp,
    address_id integer NOT NULL ,
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
    person_id integer NOT NULL,
    registration_date timestamp NOT NULL ,
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


-- Table: public.user

CREATE TABLE IF NOT EXISTS public.user
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    username character varying(45) UNIQUE NOT NULL ,
    password_hash_hint character varying(8) NOT NULL,
    check_hash character varying(128) NOT NULL ,
    customer_id integer UNIQUE NOT NULL ,
    CONSTRAINT user_id PRIMARY KEY (id),
    CONSTRAINT customer_id FOREIGN KEY (customer_id)
        REFERENCES public.customer (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user
    OWNER to postgres;

INSERT INTO public.user(
    username, password_hash_hint, check_hash, customer_id)
VALUES ('Derek', '4cbad12e', '296fd6d505f3ddf41f550a754a27541d754295fe1c125f7805e349f1d94d5330', 1);

CREATE TABLE IF NOT EXISTS public.website_session
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    user_id integer NOT NULL,
    created_at timestamp NOT NULL,
    is_repeat_session bool,
    utm_source character varying(45),
    utm_campaign character varying(45),
    utm_content character varying(45),
    device_type character varying(45),
    http_referer character varying(45),
    CONSTRAINT customer_id PRIMARY KEY (id),
    CONSTRAINT person_id FOREIGN KEY (user_id)
        REFERENCES public.user (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.website_session
    OWNER to postgres;

INSERT INTO public.website_session(
    user_id, created_at, is_repeat_session, utm_source, utm_campaign, utm_content, device_type, http_referer)
VALUES (1, '1996-12-02', false, '', '', '', '', '');

CREATE TABLE IF NOT EXISTS public.website_pageview
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    pageview_url character varying(45),
    website_session_id integer,
    CONSTRAINT website_pageviews_id PRIMARY KEY (id),
    CONSTRAINT website_session_id FOREIGN KEY (website_session_id)
        REFERENCES public.website_session (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.website_pageviews
    OWNER to postgres;

INSERT INTO public.website_pageview(
    pageview_url, website_session_id)
VALUES ('https://csca.com/home', 1);


-- +migrate Down
DROP TABLE IF EXISTS public.user;
DROP TABLE IF EXISTS public.website_pageview;
DROP TABLE IF EXISTS public.website_session;
DROP TABLE IF EXISTS public.customer;
DROP TABLE IF EXISTS public.person;
DROP TABLE IF EXISTS public.address;
