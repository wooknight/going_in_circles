--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2 (Debian 14.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: connections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.connections (
    source integer NOT NULL,
    destination integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.connections OWNER TO postgres;

--
-- Name: diary; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.diary (
    id integer NOT NULL,
    summary character varying(255) NOT NULL,
    relations jsonb,
    tags text[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    priority integer,
    source character varying(255),
    description text
);


ALTER TABLE public.diary OWNER TO postgres;

--
-- Name: diary_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.diary_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.diary_id_seq OWNER TO postgres;

--
-- Name: diary_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.diary_id_seq OWNED BY public.diary.id;


--
-- Name: plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.plans (
    id integer NOT NULL,
    plan_name character varying(255),
    plan_amount integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.plans OWNER TO postgres;

--
-- Name: plans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.plans ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.plans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_plans (
    id integer NOT NULL,
    user_id integer,
    plan_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.user_plans OWNER TO postgres;

--
-- Name: user_plans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.user_plans ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_plans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    password character varying(60),
    user_active integer DEFAULT 0,
    is_admin integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: diary id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diary ALTER COLUMN id SET DEFAULT nextval('public.diary_id_seq'::regclass);


--
-- Data for Name: connections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.connections (source, destination, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: diary; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.diary (id, summary, relations, tags, created_at, updated_at, priority, source, description) FROM stdin;
\.


--
-- Data for Name: plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.plans (id, plan_name, plan_amount, created_at, updated_at) FROM stdin;
1	Bronze Plan	1000	2022-05-12 00:00:00	2022-05-12 00:00:00
2	Silver Plan	2000	2022-05-12 00:00:00	2022-05-12 00:00:00
3	Gold Plan	3000	2022-05-12 00:00:00	2022-05-12 00:00:00
\.


--
-- Data for Name: user_plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_plans (id, user_id, plan_id, created_at, updated_at) FROM stdin;
1	7	1	\N	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, first_name, last_name, password, user_active, is_admin, created_at, updated_at) FROM stdin;
1	admin@example.com	Admin	User	$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe	1	1	2022-03-14 00:00:00	2022-03-14 00:00:00
3	rameshnaidu@gmail.com	Ramesh	Naidu	$2a$12$owme6N6wYBjTQTDqutt1z.ASp4y9BdRw/nygbIsRKLXX1bDCAjC5m	0	0	2023-04-15 20:37:18.853045	2023-04-15 20:37:18.853045
4	me@vahii.com	rams	eam	$2a$12$aFmSSMyGKu1HEVKcohV/9.Aa2kMiN1kY7z7gcjZ0bIqidVZb4QxPO	0	0	2023-04-18 21:07:05.141887	2023-04-18 21:07:05.141888
5	m@jdaj.com	ram	nad	$2a$12$Xr6RQqhdB76u/.5mCmUIfuEtuWXJEMQ7tqkoaXdHLRUbhCfjL.RcC	1	0	2023-04-18 21:08:34.342741	2023-04-18 21:09:39.31447
6	ram@hot.com	ram	ana	$2a$12$g8DfqGwWTn218DgJET8G3.jHuToA5.19AoaRINUhUzrsBI.kPxXZG	1	0	2023-04-18 21:10:15.9457	2023-04-18 21:10:30.710916
7	me@me.com	pass	ram	$2a$12$NPnivAW4QcgpAWQCjM5J4.6htaEoPvs/UmExUIDQ33dJkuScz6glW	1	0	2023-04-19 15:48:42.613567	2023-04-19 15:48:53.844796
\.


--
-- Name: diary_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.diary_id_seq', 1, false);


--
-- Name: plans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.plans_id_seq', 3, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 7, true);


--
-- Name: user_plans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_plans_id_seq', 1, true);


--
-- Name: connections connections_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connections
    ADD CONSTRAINT connections_pkey PRIMARY KEY (source, destination);


--
-- Name: diary diary_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diary
    ADD CONSTRAINT diary_pkey PRIMARY KEY (id);


--
-- Name: plans plans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plans
    ADD CONSTRAINT plans_pkey PRIMARY KEY (id);


--
-- Name: user_plans user_plans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: connections connections_destination_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connections
    ADD CONSTRAINT connections_destination_fkey FOREIGN KEY (destination) REFERENCES public.diary(id);


--
-- Name: connections connections_source_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connections
    ADD CONSTRAINT connections_source_fkey FOREIGN KEY (source) REFERENCES public.diary(id);


--
-- Name: user_plans user_plans_plan_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_plan_id_fkey FOREIGN KEY (plan_id) REFERENCES public.plans(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: user_plans user_plans_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_plans
    ADD CONSTRAINT user_plans_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

