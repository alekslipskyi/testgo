--
-- PostgreSQL database dump
--

-- Dumped from database version 10.8 (Ubuntu 10.8-0ubuntu0.18.10.1)
-- Dumped by pg_dump version 11.5 (Ubuntu 11.5-0ubuntu0.19.04.1)

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

SET default_with_oids = false;

--
-- Name: channel_users; Type: TABLE; Schema: public; Owner: deus
--

CREATE TABLE public.channel_users (
    channel_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.channel_users OWNER TO deus;

--
-- Name: channels; Type: TABLE; Schema: public; Owner: deus
--

CREATE TABLE public.channels (
    name character varying(10),
    _id integer NOT NULL
);


ALTER TABLE public.channels OWNER TO deus;

--
-- Name: channels__id_seq; Type: SEQUENCE; Schema: public; Owner: deus
--

CREATE SEQUENCE public.channels__id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.channels__id_seq OWNER TO deus;

--
-- Name: channels__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: deus
--

ALTER SEQUENCE public.channels__id_seq OWNED BY public.channels._id;


--
-- Name: ip_users; Type: TABLE; Schema: public; Owner: deus
--

CREATE TABLE public.ip_users (
    ip character varying(22) NOT NULL,
    user_id integer
);


ALTER TABLE public.ip_users OWNER TO deus;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: deus
--

CREATE TABLE public.messages (
    user_id integer NOT NULL,
    channel_id integer,
    file_url character varying(100),
    _id integer NOT NULL,
    body text
);


ALTER TABLE public.messages OWNER TO deus;

--
-- Name: messages__id_seq; Type: SEQUENCE; Schema: public; Owner: deus
--

CREATE SEQUENCE public.messages__id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.messages__id_seq OWNER TO deus;

--
-- Name: messages__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: deus
--

ALTER SEQUENCE public.messages__id_seq OWNED BY public.messages._id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: deus
--

CREATE TABLE public.users (
    _id integer NOT NULL,
    username character varying(20),
    password character varying(255),
    firstname character varying(255),
    lastname character varying(255)
);


ALTER TABLE public.users OWNER TO deus;

--
-- Name: users__id_seq; Type: SEQUENCE; Schema: public; Owner: deus
--

CREATE SEQUENCE public.users__id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users__id_seq OWNER TO deus;

--
-- Name: users__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: deus
--

ALTER SEQUENCE public.users__id_seq OWNED BY public.users._id;


--
-- Name: channels _id; Type: DEFAULT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.channels ALTER COLUMN _id SET DEFAULT nextval('public.channels__id_seq'::regclass);


--
-- Name: messages _id; Type: DEFAULT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.messages ALTER COLUMN _id SET DEFAULT nextval('public.messages__id_seq'::regclass);


--
-- Name: users _id; Type: DEFAULT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.users ALTER COLUMN _id SET DEFAULT nextval('public.users__id_seq'::regclass);


--
-- Data for Name: channel_users; Type: TABLE DATA; Schema: public; Owner: deus
--

COPY public.channel_users (channel_id, user_id) FROM stdin;
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
14	95
16	95
17	95
27	95
14	98
\.


--
-- Data for Name: channels; Type: TABLE DATA; Schema: public; Owner: deus
--

COPY public.channels (name, _id) FROM stdin;
strin2g	14
strin3g	16
strin4g	17
strin4=g	27
\.


--
-- Data for Name: ip_users; Type: TABLE DATA; Schema: public; Owner: deus
--

COPY public.ip_users (ip, user_id) FROM stdin;
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
127.0.0.1	95
127.0.0.1	98
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: deus
--

COPY public.messages (user_id, channel_id, file_url, _id, body) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: deus
--

COPY public.users (_id, username, password, firstname, lastname) FROM stdin;
95	string	b6f3cc592cd4bd571fa73b86abddac53	string	string
98	string2	b6f3cc592cd4bd571fa73b86abddac53	string	string
\.


--
-- Name: channels__id_seq; Type: SEQUENCE SET; Schema: public; Owner: deus
--

SELECT pg_catalog.setval('public.channels__id_seq', 27, true);


--
-- Name: messages__id_seq; Type: SEQUENCE SET; Schema: public; Owner: deus
--

SELECT pg_catalog.setval('public.messages__id_seq', 1, false);


--
-- Name: users__id_seq; Type: SEQUENCE SET; Schema: public; Owner: deus
--

SELECT pg_catalog.setval('public.users__id_seq', 98, true);


--
-- Name: channels channels_pk; Type: CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.channels
    ADD CONSTRAINT channels_pk PRIMARY KEY (_id);


--
-- Name: messages messages_pk; Type: CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pk PRIMARY KEY (_id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (_id);


--
-- Name: channels__id_uindex; Type: INDEX; Schema: public; Owner: deus
--

CREATE UNIQUE INDEX channels__id_uindex ON public.channels USING btree (_id);


--
-- Name: channels_name_uindex; Type: INDEX; Schema: public; Owner: deus
--

CREATE UNIQUE INDEX channels_name_uindex ON public.channels USING btree (name);


--
-- Name: messages__id_uindex; Type: INDEX; Schema: public; Owner: deus
--

CREATE UNIQUE INDEX messages__id_uindex ON public.messages USING btree (_id);


--
-- Name: users__id_uindex; Type: INDEX; Schema: public; Owner: deus
--

CREATE UNIQUE INDEX users__id_uindex ON public.users USING btree (_id);


--
-- Name: channel_users chanel_users_channels__fk; Type: FK CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.channel_users
    ADD CONSTRAINT chanel_users_channels__fk FOREIGN KEY (channel_id) REFERENCES public.channels(_id) ON DELETE CASCADE;


--
-- Name: channel_users chanel_users_users__fk; Type: FK CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.channel_users
    ADD CONSTRAINT chanel_users_users__fk FOREIGN KEY (user_id) REFERENCES public.users(_id) ON DELETE CASCADE;


--
-- Name: ip_users ip_users_users__fk; Type: FK CONSTRAINT; Schema: public; Owner: deus
--

ALTER TABLE ONLY public.ip_users
    ADD CONSTRAINT ip_users_users__fk FOREIGN KEY (user_id) REFERENCES public.users(_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

