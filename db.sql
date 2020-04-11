--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)

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

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: slh_teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_teams (
    "TeamId" integer NOT NULL,
    "Email" character varying(50) NOT NULL,
    "FirstName" character varying(50) NOT NULL,
    "LastName" character varying(50) NOT NULL,
    "MobileNo" character varying(10) NOT NULL,
    "Address" character varying(50) NOT NULL,
    "CreatedAt" timestamp without time zone NOT NULL,
    "JoiningDate" timestamp without time zone,
    "Status" character varying(50) NOT NULL,
    "Password" text NOT NULL,
    "Token" text
);


ALTER TABLE public.slh_teams OWNER TO postgres;

--
-- Name: slh_teams_TeamId_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_teams_TeamId_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_teams_TeamId_seq" OWNER TO postgres;

--
-- Name: slh_teams_TeamId_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_teams_TeamId_seq" OWNED BY public.slh_teams."TeamId";


--
-- Name: team; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.team (
    "Team_ID" integer NOT NULL,
    "First_Name" character varying(50) NOT NULL,
    "Last_Name" character varying(50) NOT NULL,
    "Mobile_No." character varying(10) NOT NULL,
    "Address" character varying(50) NOT NULL,
    "Created_At" timestamp without time zone NOT NULL,
    "Joining_At" timestamp without time zone,
    "Status" character varying(50) NOT NULL,
    "Password" text NOT NULL
);


ALTER TABLE public.team OWNER TO postgres;

--
-- Name: team_Team_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."team_Team_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."team_Team_ID_seq" OWNER TO postgres;

--
-- Name: team_Team_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."team_Team_ID_seq" OWNED BY public.team."Team_ID";


--
-- Name: slh_teams TeamId; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams ALTER COLUMN "TeamId" SET DEFAULT nextval('public."slh_teams_TeamId_seq"'::regclass);


--
-- Name: team Team_ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.team ALTER COLUMN "Team_ID" SET DEFAULT nextval('public."team_Team_ID_seq"'::regclass);


--
-- Data for Name: slh_teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_teams ("TeamId", "Email", "FirstName", "LastName", "MobileNo", "Address", "CreatedAt", "JoiningDate", "Status", "Password", "Token") FROM stdin;
3	aksingh@gmail.com	ashish	kumar	09876543	A 140, Rajiv	2020-04-04 19:06:31.141501	2020-04-04 19:06:31.141501	active	$2a$08$Xzq1NdQJD16yWvqbrmOGQuXJYXMVCY/DYOjcWOWmf9Z5cILJ83jTi	\N
4	sk@gmail.com	ashish	kumar	1234567	A 140, Rajiv	2020-04-04 19:06:46.505203	2020-04-04 19:06:46.505203	active	$2a$08$sIUuLzZmqxgPh/o0W/d0Fe3r32oqov.7JUdLNOLH5WL1FepB/Pxcm	\N
5	ssk@gmail.com	ashish	kumar	0987	A 140, Rajiv	2020-04-04 19:06:54.617889	2020-04-04 19:06:54.617889	active	$2a$08$TAEp59OhrOgexGTClx3WRuxM7upHoMVTcLvsgx3yV1u7lrOpNugGS	\N
6	asfk@gmail.com	ashish	kumar	78563	A 140, Rajiv	2020-04-04 19:07:03.83426	2020-04-04 19:07:03.83426	active	$2a$08$RHoYqtsKT6.8Ye.pROE1mugZPRqRSwxhC95zNSvxxVGLgZhLIWJzO	\N
7	cat@gmail.com	ashish	kumar	215	A 140, Rajiv	2020-04-04 19:07:11.252671	2020-04-04 19:07:11.252671	active	$2a$08$guynIfhQ9rljkRJ4rqUj4uUuIWYsV7PgfqFZlBhTbzvojTRg133k2	\N
11	coll@gmail.com	ashish	kumar	236545678	A 140, Rajiv	2020-04-04 19:08:02.205407	2020-04-04 19:08:02.205407	active	$2a$08$xutQ2RwepO5JPV4//84v0ubaOueO.Kag2D37K7l/EupSg32oqYc1W	\N
13	tall@gmail.com	ashish	kumar	90785	A 140, Rajiv	2020-04-04 19:08:10.082299	2020-04-04 19:08:10.082299	active	$2a$08$tHXG6qCnjp4WVPnCh1dzF.nS5AA9WcEX3Ng1fmX4751iQuYJMkxZa	\N
1	ashish@gmail.com	ashish	kumar	2412	A 140, Rajiv	2020-04-04 19:05:36.502473	2020-04-04 19:05:36.502473	active	$2a$08$scn9SIAZWxiD.2WbHZPbc.k3XA/zWL.skGVtKIJLxKToIY63noeDy	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzaGlzaEBnbWFpbC5jb20iLCJleHAiOjE1ODY5MTE5ODJ9.RcJEvQ2IkLoPfjqRzMymY-6nLc70d6rsCQtgbnUBduA
8	lol@gmail.com	ashish	kumar	234567890	A 140, Rajiv	2020-04-04 19:07:22.425071	2020-04-04 19:07:22.425071	active	$2a$08$7ozfv6lM2IM4MHlvFHmew.7HNkol7KDwcHIDjn8vS66BvkvApDtCq	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxvbEBnbWFpbC5jb20iLCJleHAiOjE1ODcwODI0MDV9.VhtNOFra2-ncsnYIDHgO_kPA9mbFVEajwXZQF3cmAd4
16	aman@gmail.com	aman	singh	123456789	A 140, Rajiv	2020-04-06 22:50:43.74743	2020-04-06 22:50:43.74743	active	$2a$08$0.OQHgHUfjWKRqmo6/6gCOEpNa.9RQYJFuBOc6pHP4GsJ.0d5drxW	\N
2	aks@gmail.com	ashish	kumar	34567890	A 140, Rajiv	2020-04-04 19:06:21.673562	2020-04-04 19:06:21.673562	active	$2a$08$aTliue5D9Fw/RNc0IHmXh.V4D4gnGOQ2HHCCExpmadNxnW/SJttsq	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFrc0BnbWFpbC5jb20iLCJleHAiOjE1ODcxNzc0NDJ9.BmJeNtM8SF9gXAYIZH0ecHxP6WFqfw2d-4YJPWob11o
\.


--
-- Data for Name: team; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.team ("Team_ID", "First_Name", "Last_Name", "Mobile_No.", "Address", "Created_At", "Joining_At", "Status", "Password") FROM stdin;
\.


--
-- Name: slh_teams_TeamId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_teams_TeamId_seq"', 16, true);


--
-- Name: team_Team_ID_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."team_Team_ID_seq"', 1, false);


--
-- Name: slh_teams slh_teams_Email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams
    ADD CONSTRAINT "slh_teams_Email_key" UNIQUE ("Email");


--
-- Name: slh_teams slh_teams_MobileNo_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams
    ADD CONSTRAINT "slh_teams_MobileNo_key" UNIQUE ("MobileNo");


--
-- Name: slh_teams slh_teams_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams
    ADD CONSTRAINT slh_teams_pkey PRIMARY KEY ("TeamId");


--
-- Name: team team_Mobile_No._key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.team
    ADD CONSTRAINT "team_Mobile_No._key" UNIQUE ("Mobile_No.");


--
-- Name: team team_Password_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.team
    ADD CONSTRAINT "team_Password_key" UNIQUE ("Password");


--
-- Name: team team_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.team
    ADD CONSTRAINT team_pkey PRIMARY KEY ("Team_ID");


--
-- PostgreSQL database dump complete
--

