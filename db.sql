--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-2.pgdg18.04+1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-2.pgdg18.04+1)

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
-- Name: slh_assign_customer_with_partner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_assign_customer_with_partner (
    "Slot_Date" character(1),
    "Slot_TIme" abstime,
    "Transaction_ID" abstime,
    "Status" abstime,
    "Commision_Plant" abstime,
    "Created_Time" abstime,
    "Updated_Time" abstime,
    "Created_By" bigint[]
);


ALTER TABLE public.slh_assign_customer_with_partner OWNER TO postgres;

--
-- Name: slh_customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_customers (
    "Customer_Id" integer NOT NULL,
    "Customer_Email" character varying(50) NOT NULL,
    "CreatedAt" timestamp without time zone NOT NULL,
    "Customer_Gender" character varying(50) NOT NULL,
    "Customer_Name" character varying(50) NOT NULL,
    "Customer_Address" character varying(200) NOT NULL,
    "Registered_Source" character varying(100) NOT NULL,
    "Last_Access_Time" timestamp without time zone,
    "Customer_Souls_Id" text,
    "Customer_Mobile_No" integer NOT NULL,
    "Status" boolean NOT NULL,
    "Pincode" integer NOT NULL
);


ALTER TABLE public.slh_customers OWNER TO postgres;

--
-- Name: slh_customers_Customer_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_customers_Customer_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_customers_Customer_Id_seq" OWNER TO postgres;

--
-- Name: slh_customers_Customer_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_customers_Customer_Id_seq" OWNED BY public.slh_customers."Customer_Id";


--
-- Name: slh_customers_pending_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_customers_pending_orders (
    "Customer_Order_Id" integer NOT NULL,
    "Customer_Id" integer NOT NULL,
    "Number_Of_Therapists_Required" integer NOT NULL,
    "Therapist_Gender" character varying(50) NOT NULL,
    "Massage_For" character varying(50) NOT NULL,
    "Slot_Time" timestamp without time zone NOT NULL,
    "Slot_Date" timestamp without time zone NOT NULL,
    "Customer_Address" character varying(200) NOT NULL,
    "Latitude" character varying(100) NOT NULL,
    "Longitude" character varying(100) NOT NULL,
    "Massage_Duration" text NOT NULL,
    "CreatedAt" timestamp without time zone NOT NULL,
    "Customer_Souls_Id" text,
    "Pincode" integer,
    "Is_Order_Confirmed" boolean,
    "Customer_Name" text,
    "Total_Order_Amount" integer,
    "Merchant_Transaction_Id" text
);


ALTER TABLE public.slh_customers_pending_orders OWNER TO postgres;

--
-- Name: slh_customers_pending_orders_Customer_Primary_Order_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq" OWNER TO postgres;

--
-- Name: slh_customers_pending_orders_Customer_Primary_Order_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq" OWNED BY public.slh_customers_pending_orders."Customer_Order_Id";


--
-- Name: slh_partners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_partners (
    "Partner_Id" integer NOT NULL,
    "Partner_Name" text,
    "Partner_Gender" text,
    "Partner_Address" text,
    "Partner_Mobile_No" integer,
    "Partner_Email" text,
    "Pincode" integer,
    "Latitude" text,
    "Longitude" text,
    "Onboard_Date" timestamp without time zone,
    "CreatedAt" timestamp without time zone,
    "UpdatedAt" timestamp without time zone,
    "Last_Updated_By" text,
    "Per_Visit_Price_Commission" integer,
    "Commission_Type" text,
    "CreatedBy" text
);


ALTER TABLE public.slh_partners OWNER TO postgres;

--
-- Name: slh_partners_Partner_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_partners_Partner_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_partners_Partner_Id_seq" OWNER TO postgres;

--
-- Name: slh_partners_Partner_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_partners_Partner_Id_seq" OWNED BY public.slh_partners."Partner_Id";


--
-- Name: slh_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_roles (
    "Role_Id" integer NOT NULL,
    "Role_Name" text NOT NULL,
    "Role_Status" boolean NOT NULL
);


ALTER TABLE public.slh_roles OWNER TO postgres;

--
-- Name: slh_roles_RoleId_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_roles_RoleId_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2000
    CACHE 1;


ALTER TABLE public."slh_roles_RoleId_seq" OWNER TO postgres;

--
-- Name: slh_roles_Role_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_roles_Role_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_roles_Role_Id_seq" OWNER TO postgres;

--
-- Name: slh_roles_Role_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_roles_Role_Id_seq" OWNED BY public.slh_roles."Role_Id";


--
-- Name: slh_team_has_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_team_has_role (
    "Team_Has_Role_Id" integer NOT NULL,
    "CreatedAt" timestamp without time zone NOT NULL,
    "UpdatedAt" timestamp without time zone,
    "Team_Id" integer NOT NULL,
    "Status" text NOT NULL,
    "FirstName" text,
    "LastName" text
);


ALTER TABLE public.slh_team_has_role OWNER TO postgres;

--
-- Name: slh_team_has_role_Team_Has_Role_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_team_has_role_Team_Has_Role_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_team_has_role_Team_Has_Role_Id_seq" OWNER TO postgres;

--
-- Name: slh_team_has_role_Team_Has_Role_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_team_has_role_Team_Has_Role_Id_seq" OWNED BY public.slh_team_has_role."Team_Has_Role_Id";


--
-- Name: slh_teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_teams (
    "FirstName" character varying(50) NOT NULL,
    "LastName" character varying(50) NOT NULL,
    "Email" text NOT NULL,
    "Address" text NOT NULL,
    "JoiningDate" timestamp without time zone DEFAULT now() NOT NULL,
    "TeamId" integer NOT NULL,
    "CreatedAt" timestamp without time zone DEFAULT now() NOT NULL,
    "Password" text NOT NULL,
    "Token" text,
    "Status" text NOT NULL,
    "MobileNo" text NOT NULL,
    "Role" text
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
-- Name: slh_transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_transactions (
    "Customer_Order_Id" integer,
    "Customer_Id" integer,
    "Therapist_Gender" character varying,
    "Massage_For" character varying,
    "Slot_Time" timestamp without time zone,
    "Slot_Date" date,
    "Customer_Address" character varying,
    "Pincode" integer,
    "Transaction_Mode" character varying,
    "Bank_Type" character varying,
    "Merchant_Transaction_Id" text,
    "Massage_Duration" text,
    "Customer_Name" text,
    "CreatedAt" timestamp without time zone,
    "Total_Order_Amount" integer,
    "Customer_Souls_Id" text,
    "Payment_Gateway_Id" text,
    "Latitude" text,
    "Longitude" text,
    "Number_Of_Therapist_Required" integer,
    "Payment_Gateway_Mode" text
);


ALTER TABLE public.slh_transactions OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    "Name" text,
    "Email" text NOT NULL,
    "Password" text NOT NULL,
    "Token" text,
    "CreatedAt" timestamp without time zone,
    "UpdatedAt" timestamp without time zone,
    "Role" text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: slh_customers Customer_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers ALTER COLUMN "Customer_Id" SET DEFAULT nextval('public."slh_customers_Customer_Id_seq"'::regclass);


--
-- Name: slh_customers_pending_orders Customer_Order_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers_pending_orders ALTER COLUMN "Customer_Order_Id" SET DEFAULT nextval('public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq"'::regclass);


--
-- Name: slh_partners Partner_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_partners ALTER COLUMN "Partner_Id" SET DEFAULT nextval('public."slh_partners_Partner_Id_seq"'::regclass);


--
-- Name: slh_roles Role_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_roles ALTER COLUMN "Role_Id" SET DEFAULT nextval('public."slh_roles_Role_Id_seq"'::regclass);


--
-- Name: slh_team_has_role Team_Has_Role_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_team_has_role ALTER COLUMN "Team_Has_Role_Id" SET DEFAULT nextval('public."slh_team_has_role_Team_Has_Role_Id_seq"'::regclass);


--
-- Name: slh_teams TeamId; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams ALTER COLUMN "TeamId" SET DEFAULT nextval('public."slh_teams_TeamId_seq"'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: slh_assign_customer_with_partner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_assign_customer_with_partner ("Slot_Date", "Slot_TIme", "Transaction_ID", "Status", "Commision_Plant", "Created_Time", "Updated_Time", "Created_By") FROM stdin;
\.


--
-- Data for Name: slh_customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers ("Customer_Id", "Customer_Email", "CreatedAt", "Customer_Gender", "Customer_Name", "Customer_Address", "Registered_Source", "Last_Access_Time", "Customer_Souls_Id", "Customer_Mobile_No", "Status", "Pincode") FROM stdin;
39	kalim@1234	2020-04-08 01:05:21.630388	poiuytrdes	kalim	ertyj	website	2020-04-08 01:05:21.630388	2020040839	172345615	t	9854334
40	kalim@1234	2020-04-08 01:06:39.704405	poiuytrdes	kalim	ertyj	website	2020-04-08 01:06:39.704405	2020040840	1725615	t	9854334
34	owaish@1234	2020-04-07 16:44:13.338694	male	ashish	ertyj	sdfghj	2020-04-07 16:44:24.024427	2020040734	9876543	f	123
41	owaish@1234	2020-04-08 10:51:30.046757	male	ayush	ertyj	sdfghj	2020-04-08 10:51:30.046757	2020040841	987656543	f	123
36	owaish@1234	2020-04-07 18:43:46.476663	male	laptop	ertyj	sdfghj	2020-04-07 19:37:19.496711	2020040736	956543	f	123
42	kalim@1234	2020-04-08 14:26:05.92923	poiuytrdes	kumar	ertyj	website	2020-04-08 14:26:05.92923	2020040842	121	t	9854334
38	owaish@1234	2020-04-07 22:57:58.287051	male	laptop	ertyj	sdfghj	2020-04-07 22:57:58.287052	2020040738	5335	f	123
43	kalim@1234	2020-04-08 21:31:13.75261	poiuytrdes	kumar	ertyj	website	2020-04-08 21:31:13.75261	20200408-43	12561	t	9854334
44	kalim@1234	2020-04-08 21:31:34.472219	poiuytrdes	kumar	ertyj	website	2020-04-08 21:31:34.472219	20200408-44	1256155555	t	9854334
45	kalim@1234	2020-04-08 21:43:13.069007	poiuytrdes	kumar	ertyj	website	2020-04-08 21:43:13.069007	20200408-45	125615545	t	9854334
46	kalim@1234	2020-04-08 22:07:26.431891	poiuytrdes	kumar	ertyj	website	2020-04-08 22:07:26.431891	20200408-46	125545	t	9854334
47	Johnw@1234	2020-04-09 23:06:05.839646	Male	John-Wick	Ny	website	2020-04-10 15:05:11.371335	2020040947	999	t	135
\.


--
-- Data for Name: slh_customers_pending_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers_pending_orders ("Customer_Order_Id", "Customer_Id", "Number_Of_Therapists_Required", "Therapist_Gender", "Massage_For", "Slot_Time", "Slot_Date", "Customer_Address", "Latitude", "Longitude", "Massage_Duration", "CreatedAt", "Customer_Souls_Id", "Pincode", "Is_Order_Confirmed", "Customer_Name", "Total_Order_Amount", "Merchant_Transaction_Id") FROM stdin;
9	36	5	zxcvb	zxcvbn	2020-04-07 19:37:32.242377	2020-04-07 19:37:32.242377	wqewrtydfghjk	qwerty	asdfg	asdfghj	2020-04-07 19:37:32.242376	2020040736	987654	f	owasih	10000	\N
10	36	5	zxcvb	zxcvbn	2020-04-07 21:33:13.790638	2020-04-07 21:33:13.790637	wqewrtydfghjk	qwerty	asdfg	asdfghj	2020-04-07 21:33:13.790637	2020040736	987654	f	owasih	10000	158627539336
11	36	5	zxcvb	zxcvbn	2020-04-07 21:34:24.432404	2020-04-07 21:34:24.432404	wqewrtydfghjk	qwerty	asdfg	asdfghj	2020-04-07 21:34:24.432403	2020040736	987654	f	owasih	10000	158627546436
12	36	5	zxcvb	zxcvbn	2020-04-07 21:34:36.322148	2020-04-07 21:34:36.322148	wqewrtydfghjk	qwerty	asdfg	asdfghj	2020-04-07 21:34:36.322148	2020040736	987654	f	owasih	10000	158627547636
13	36	5	zxcvb	zxcvbn	2020-04-07 22:27:27.526831	2020-04-07 22:27:27.526831	wqewrtydfghjk	qwerty	asdfg	asdfghj	2020-04-07 22:27:27.526831	2020040736	987654	f	owasih	10000	158627864736
14	38	2	male	cbn	2020-04-07 22:58:35.691487	2020-04-07 22:58:35.691486	ertyj	qwerty	asdfg	trew	2020-04-07 22:58:35.691486	2020040738	9854334	f	kalim	10000	158628051538
15	38	2	male	cbn	2020-04-07 22:58:41.269672	2020-04-07 22:58:41.269672	ertyj	qwerty	asdfg	trew	2020-04-07 22:58:41.269672	2020040738	9854334	f	kalim	10000	158628052138
16	38	2	male	cbn	2020-04-07 22:58:44.053138	2020-04-07 22:58:44.053138	ertyj	qwerty	asdfg	trew	2020-04-07 22:58:44.053138	2020040738	9854334	f	kalim	10000	158628052438
17	38	2	male	cbn	2020-04-07 22:58:48.846522	2020-04-07 22:58:48.846522	ertyj	qwerty	asdfg	trew	2020-04-07 22:58:48.846522	2020040738	9854334	f	kalim	10000	158628052838
18	38	2	male	cbn	2020-04-07 22:58:54.26731	2020-04-07 22:58:54.26731	ertyj	qwerty	asdfg	trew	2020-04-07 22:58:54.26731	2020040738	9854334	f	kalim	10000	158628053438
19	39	2	male	cbn	2020-04-08 01:05:35.177064	2020-04-08 01:05:35.177064	ertyj	qwerty	asdfg	trew	2020-04-08 01:05:35.177064	2020040839	9854334	f	kalim	10000	158628813539
20	40	2	male	cbn	2020-04-08 01:06:47.71179	2020-04-08 01:06:47.71179	ertyj	qwerty	asdfg	trew	2020-04-08 01:06:47.71179	2020040840	9854334	f	kalim	10000	158628820740
21	40	2	male	cbn	2020-04-08 15:59:05.800142	2020-04-08 15:59:05.800142	ertyj	qwerty	asdfg	trew	2020-04-08 15:59:05.800141	2020040840	9854334	f	kalim	10000	158634174540
22	40	2	male	cbn	2020-04-10 15:05:25.65068	2020-04-10 15:05:25.65068	ertyj	qwerty	asdfg	trew	2020-04-10 15:05:25.65068	2020040840	9854334	f	kalim	10000	1586511325-40
\.


--
-- Data for Name: slh_partners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_partners ("Partner_Id", "Partner_Name", "Partner_Gender", "Partner_Address", "Partner_Mobile_No", "Partner_Email", "Pincode", "Latitude", "Longitude", "Onboard_Date", "CreatedAt", "UpdatedAt", "Last_Updated_By", "Per_Visit_Price_Commission", "Commission_Type", "CreatedBy") FROM stdin;
3	buzurgbuddha	male	poiu	1095	ow@12	10909	LL	pp	2020-04-10 18:55:21.808417	2020-04-10 18:55:21.808417	2020-04-10 18:55:21.808417	ll	190	oiioi	uu
4	owaishkalim	male	poiu	1090	owa@12	10909	LL	pp	2020-04-10 21:15:57.539975	2020-04-10 21:15:57.539975	2020-04-10 21:15:57.539975	ll	190	oiioi	uu
5	owaishkalimlll	male	poiu	1090	owa@1hhh2	10909	LL	pp	2020-04-10 23:27:25.090668	2020-04-10 23:27:25.090668	2020-04-10 23:27:25.090668	ll	190	oiioi	uu
\.


--
-- Data for Name: slh_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_roles ("Role_Id", "Role_Name", "Role_Status") FROM stdin;
1	Admin	t
2	Accountant	t
3	Customer Service	t
\.


--
-- Data for Name: slh_team_has_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_team_has_role ("Team_Has_Role_Id", "CreatedAt", "UpdatedAt", "Team_Id", "Status", "FirstName", "LastName") FROM stdin;
1	2020-04-09 14:36:40.066848	\N	185	sdfbhm	\N	\N
1	2020-04-09 14:46:07.803057	2020-04-09 14:46:08.151185	188	sdfbhm	\N	\N
1	2020-04-09 15:18:25.656218	2020-04-09 15:18:25.753062	193	Active	\N	\N
1	2020-04-09 15:20:15.389856	2020-04-09 15:20:15.453307	194	Active	\N	\N
2	2020-04-09 15:21:21.13906	2020-04-09 15:21:21.406727	195	Inactive	\N	\N
1	2020-04-09 18:52:41.931823	2020-04-09 18:52:41.975515	202	sdfbhm	\N	\N
2	2020-04-09 18:52:56.861374	2020-04-09 18:52:57.306618	203	Deleted	\N	\N
2	2020-04-09 20:13:22.809635	2020-04-09 20:13:23.394341	204	Inactive	\N	\N
1	2020-04-10 10:56:05.138526	2020-04-10 10:56:05.184158	206	sdfbhm	\N	\N
2	2020-04-10 17:18:09.019004	2020-04-10 17:18:10.14642	208	Inactive	\N	\N
2	2020-04-09 23:09:52.93417	2020-04-09 23:09:52.973462	205	sdfbhm	\N	\N
1	2020-04-11 00:45:12.344031	2020-04-11 00:45:13.515938	213	sdfbhm	John	Wick
1	2020-04-11 00:51:01.930495	2020-04-11 00:51:01.990551	214	sdfbhm	John	Wick
1	2020-04-11 00:55:26.335268	2020-04-11 00:55:26.564514	215	sdfbhm	John	Wick
1	2020-04-11 00:59:06.519941	2020-04-11 00:59:06.623231	216	sdfbhm	John	Wick
1	2020-04-11 12:33:05.495498	2020-04-11 12:33:05.538835	218	sdfbhm	John	Wick
\.


--
-- Data for Name: slh_teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_teams ("FirstName", "LastName", "Email", "Address", "JoiningDate", "TeamId", "CreatedAt", "Password", "Token", "Status", "MobileNo", "Role") FROM stdin;
AShish	kumar	ak@asf.as	Qwerty	2020-04-04 22:47:53.604904	141	2020-04-04 22:47:53.604904	$2a$08$tsYFF7bP/noieJBR14.rWeXMR5Gg9bUvkYQFHWz7xXLdQSW4kgE8m	\N	Inactive	234567898	\N
Ashish	Singh	asingh2@ch.iitr.ac.in	Rajiv Bhwan, IIT Roorkee	2020-04-04 12:19:26.929737	120	2020-04-04 12:19:26.929737	$2a$08$Odz87MnbLSN4evI12sVqPecm.YRBRrPhi1PSUooU0onWmBB3BJAHa	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzaW5naDJAY2guaWl0ci5hYy5pbiIsImV4cCI6MTU4Njg4NzI0OH0.b1S7c_-yt2QIXYwRfSsQe0JaJ0F_HCby5ObiTcQQIcY	Inactive	+918769304216	Accountant
saSdfa	test	sh@gmail.com	dfghjk	2020-04-03 17:14:24.273031	114	2020-04-03 17:14:24.273031	$2a$08$XWC2NdF8/5a4fv9uiBvD2ujxHNY9laSh89ERnbJ.W.cxPOdfjc0JK	\N	active	94112121212	admin
Kartik 	Aryan	k@kk.kk	qwertyui	2020-04-04 22:53:39.313767	142	2020-04-04 22:53:39.313767	$2a$08$r1N2fi5YGNgGKUgelFBA7u.os5AkqHe.YaJRG5VCKmC3fG1TF0PmC	\N	Inactive	23456789	\N
sdvf	sdfcvbn	as@gmail.com	v	2020-04-04 20:50:12.84326	139	2020-04-04 20:50:12.84326	$2a$08$oNfiL1o0bskZJZQJFuHaZu3Kev9kBGS3E.5VVFpSXIASut7vkT2KK	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtYWlsLmNvbSIsImV4cCI6MTU4NjkxMzgxOH0.0XI66C4bCqeC7YMUvxe4akls1ECjkwo875kZM6nalNU	accountant	94112121212	\N
AShish	kumar	ahirnotia@ch.iitr.ac.in	dfghkj	2020-04-04 17:47:10.247661	132	2020-04-04 17:47:10.247662	$2a$08$QsKrZ4Vsgsj.fRw4eLo4jOZ0D4PDQF7BgC9sCoeZ/OnHxHriyjDka	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFoaXJub3RpYUBjaC5paXRyLmFjLmluIiwiZXhwIjoxNTg3NDk1MzQzfQ.VeoX0JaokyJ43dGP5S4buJ6o6O5xjXbqqarLLSbzG78	Active	1234567892	Admin
AShish	kumar	sd@af.af	asdgfhgjk	2020-04-03 17:46:45.846419	115	2020-04-03 17:46:45.84642	$2a$08$mHcOXlXlOaAdv5rsZ9O8B.8LyBm/zwwbQXzr11Gt5kICor/SoCRGW	\N	inactive	1234567789	admin
AShish	kumar	sds@af.af	asdgfhgjk	2020-04-03 17:48:09.661473	116	2020-04-03 17:48:09.661473	$2a$08$EfOW8CYMoHX3rQIgcWEsROUVCndwe2yHElpTzJ07ZAD1OkVcolCRW	\N	inactive	1234567789	admin
	cvbn	sh@gghjsdertfvghmail.com	v	2020-04-04 17:38:11.467879	125	2020-04-04 17:38:11.467879	$2a$08$grUDESzu7oc4kPvGn3zFmuPh6RLoIVq1Orwu19RtKCEP5K8VIu54C	\N		94112121212	
saSdfa	test	sh@gghjmail.com	dfghjk	2020-04-03 22:07:32.341728	117	2020-04-03 22:07:32.341728	$2a$08$KsiRMhJslypxxptG3H9AM.9dvEewCFZE7SrS/Wk3GuJkpFpKrPiti	\N	active	94112121212	admin
bn	test	sh@gghjvghmail.com	v	2020-04-03 22:13:11.395077	119	2020-04-03 22:13:11.395077	$2a$08$1kG4naPdgil9YQW3WhrdfuxS64HJz.wtEqtG4Ow6zJ7YepfIYiE06	\N	fgh	94112121212	admin
sdvf	sdfcvbn	sh@gghjsdertfvail.com	v	2020-04-04 17:38:43.610279	127	2020-04-04 17:38:43.610279	$2a$08$XVV/PmNuQNFyKLK0bFdu9uw2nL1E6EqLphlzRRrnjitT.vjp1uYmO	\N	sdfbgnhm	94112121212	zxcv 
sdvf	sdfcvbn	sh@gghjsdertfvartyil.com	v	2020-04-04 17:38:54.665595	128	2020-04-04 17:38:54.665595	$2a$08$PHaTWixqJLnXxufOGSsEIeYPc2iGSe7IGKqdc7vuonfi99YVIK1ZW	\N	sdfbgnhm	94112121212	zxcv 
sdvf	sdfcvbn	sh@gghjsdefrgthsdertfvartyil.com	v	2020-04-04 17:39:07.276873	130	2020-04-04 17:39:07.276873	$2a$08$CqJKFGCslUEL4v0Vt.z.rOLtIbk7mOHOVtiU8uF2GWSh3SN11RJBS	\N	sdfbgnhm	94112121212	zxcv
asfa	test	owaish@gmail.com	dfghjk	2020-04-03 17:02:41.174742	113	2020-04-03 17:02:41.174743	$2a$08$xQ7ctow8rGKAUrcRWw.LBetRFGK.GwHOU2M7GM.jcKfpchDDNhXx2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzaEBnbWFpbC5jb20iLCJleHAiOjE1ODY5MTI5MDV9.zyfDAqPo4AdYNo-V2kWB6gmEWXLR7t3--H9atmOczkM	accountant	94112121212	admin
bn	cvbn	sh@gghjsdfvghmail.com	v	2020-04-04 17:11:45.913296	123	2020-04-04 17:11:45.913296	$2a$08$XIZ./xeRmlr55ZDqHbfXF.fPCuJRLcHW3uCjJycv0FPYB0PB6Hk8.	\N	fgh	94112121212	admin
AShish	kumar	a@a.cc	asdfghj	2020-04-04 17:44:27.115112	131	2020-04-04 17:44:27.115112	$2a$08$bgExze5662U3E8v3pPtlE.AmoJ2WDIQkur7135857Z5M1..8.qdxS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYS5jYyIsImV4cCI6MTU4NjkwMjU0MX0.ruT9B7wYQXksugTi6FCSggdSqVvXk_I5H9bbiI0WKo4	Active	1234567894	Admin
sdvf	sdfcvbn	sh@gghjssdcfvdefrgthsdertfvartyil.com	v	2020-04-04 17:48:01.449996	134	2020-04-04 17:48:01.449996	$2a$08$sc1O84ampL5zG/A2yLggq.6UNosQlxINgZPS3C4KOYrBi7aduJuPi	\N	sdfbgnhm	94112121212	zxcv
sdvf	sdfcvbn	sh@gghjssdcfasdfgvdefrgthsdertfvartyil.com	v	2020-04-04 17:48:05.526691	135	2020-04-04 17:48:05.526691	$2a$08$.q5yPJhyOto60AZdWjNrjemK2VxkgxN0NrAJGI8cFbDt1GxqYs9sG	\N	sdfbgnhm	94112121212	zxcv
sdvf	sdfcvbn	sh@gghjssdcfwertyasdfgvdefrgthsdertfvartyil.com	v	2020-04-04 17:49:15.821892	137	2020-04-04 17:49:15.821892	$2a$08$rv11WiK/voyPw7aqH588NeELy/DW6psUv.V7TYDR.DDaISJMfZ0Zi	\N	sdfbgnhm	94112121212	zxcv
AShish	kumar	ashutoshhhirnotia23@gmail.com	Gali no 16 B,Lajjapuri ,chamri road ,Hapur	2020-04-04 22:03:29.997268	140	2020-04-04 22:03:29.997268	$2a$08$XDRmXktAKpz2UFT1QzgXjeKvzvtRGW1iAQS/WjNYVZkCgkf3SnuCe	\N	Inactive	09084768046	\N
sdvf	sdfcvbn	as@gmvbnail.com	v	2020-04-05 16:12:01.726735	145	2020-04-05 16:12:01.726736	$2a$08$JfZvVtQjv4asM1vBYhU3AewUtysWC75fgCsMGasJPDviR7XDAHVBi	\N	sdfbgnhm	94112121212	\N
sdvf	sdfcvbn	as@gmvbvbnail.com	v	2020-04-05 16:49:54.542103	148	2020-04-05 16:49:54.542103	$2a$08$6GwLCpmKrIjIpT1ok27Aw.5VzV.WcqMwYWpVkox7fjqWk5OfITjD2	\N	sdfbgnhm	94112121212	\N
sdvf	sdfcvbn	as@gmvbvbbnail.com	v	2020-04-05 17:11:34.525998	152	2020-04-05 17:11:34.525998	$2a$08$u6u0N.HdLZBDRY3IE2zIRuuuYufeiGrqcglXeBHxz9fkX/G5zGWJq	\N	sdfbgnhm	94112121212	\N
sdvf	sdfcvbn	as@gmvbvbbxcnail.com	jv	2020-04-05 19:42:54.833455	155	2020-04-05 19:42:54.833455	$2a$08$isp0ldIIBQD5MeiZNgkkgelFa88Lch/PVBpGYhUpFa6.6/DXI.3eG	\N	sdfbgnhm	94112121212	\N
sdvf	ghjk	as@gmvbvbbbnmxcnail.com	jv	2020-04-05 22:24:23.858635	158	2020-04-05 22:24:23.858636	$2a$08$pSK3bO90lI4qQg.VV1bV6.F694BjTelTDyjYCMZnFwcrKGLwUBaju	\N	sdfbgnhm	94112121212	\N
owaish	kalim	np@kliakashlodum	s	2020-04-09 14:43:25.896769	186	2020-04-09 14:43:25.896769	$2a$08$BeWp.U7AhqE0hQnvJpQ1EODfU8c25Ehw6JLEmsw6Mr0ijbf/mN3ca	\N	sdfbhm	94121212	Admin
ayush	bhai	ayush@123456	rtdt	2020-04-09 18:17:04.429046	196	2020-04-09 18:17:04.429046	$2a$08$xHj/41rSidIlKdiFD8hymOSMbqDKGsuDeEwWbfx14bePeyrTwq/Gm	\N	Inactive	7894561235	Accountant
sdvf	ghjk	owaish@kalim	jv	2020-04-05 22:29:49.415256	161	2020-04-05 22:29:49.415256	$2a$08$4gIPpUW6mHOPJtXqfyg3.ecAKUidtoaMtOTn0.LCts.g3XuL8Rc2q	\N	sdfbgnhm	94112121212	\N
owaish	gheem	a@a.aa	12qwertbnbnmyuil	2020-04-05 21:32:10.470454	157	2020-04-05 21:32:10.470455	$2a$08$l.jy1pudS6lWL/Q.4/ctjOj0sGqYz.OoigzgvRanMb1/QiNast0Ye	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtdmJ2YmJibm14Y25haWwuY29tIiwiZXhwIjoxNTg3MDAyOTQ0fQ.LVug_MIsMthyLgRp5VzPI606bGGMOWuY6g3Wbi79S5U	actddive	9411212fdf1212	\N
AShish	kumar	a@ab.aa	rtfghjk	2020-04-06 15:17:00.334879	166	2020-04-06 15:17:00.334879	$2a$08$bOLJhUs59z/6YRI7MuEPq.N1Pm6kmsQU1DRqnV1kJSVlkzSwXuqAG	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYWIuYWEiLCJleHAiOjE1ODcwNzg2MDZ9.3nRVUyyOE6bbrbBkazdmWuwmaSSipU6fS9fmgmmZm_0	Active	1234567890	\N
aakash123456	sdfgh	bnm@ayush.com	vbn	2020-04-06 15:28:35.30426	169	2020-04-06 15:28:35.30426	$2a$08$.LhaCSpeamAif75sGZrSfeeNYoJReNEN9jW/KFZsQDgRDRsPddDqC	\N	Inactive	234567890	\N
vinay	tyu	a@ayush.com	sdfghjk	2020-04-06 15:26:33.656613	168	2020-04-06 15:26:33.656613	$2a$08$vQAW3vcZ6nE.cSZm7LDDGeob7CrTRCBqQDU2ekbcsDAosB5Iyie0C	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYXl1c2guY29tIiwiZXhwIjoxNTg3MDc0MjQ0fQ.OaEib0e9TgjrutcisGJq5U0syem2ccFJJdRcaZNOPDY	Active	1234567890	\N
owa1212123fghj456ish	gheqwesdfgrtyem	owaish@kalim1	1342qwertbnbnmyuil	2020-04-05 22:24:35.820247	160	2020-04-05 22:24:35.820247	fgh		sdfgh	123494112121234fdf1212	\N
sdvf	ghjk	owaish@sdkalim	jv	2020-04-05 22:59:01.383303	163	2020-04-05 22:59:01.383303	$2a$08$A8ACFvKdrJBStZ9SzM5q1OG5cxN18hZStb98iKcz1y8QNrkfsw/Ku	\N	sdfbgnhm	94112121212	\N
z	ghjk	owaisqwrhcvbn@sdkalim	s	2020-04-08 18:53:07.594736	173	2020-04-08 18:53:07.594736	$2a$08$vTbwmnIIswLiShzX7452j.A5jU.ntV1bMezQ0FL3iZyeIERtNypqG	\N	Active	94112121212	\N
sdvf	ghjk	owaisqwhcvbn@sdkalim	jv	2020-04-06 02:14:07.737403	165	2020-04-06 02:14:07.737403	$2a$08$K/rpaPrRRAE4srC68hUzDusxgnb45nKJle30BvIldlDDjijlJEfCK	\N	sdfbgnhm	94112121212	\N
owaish	kalim	np@kliakashmahalodum	s	2020-04-09 14:46:07.803057	188	2020-04-09 14:46:07.803057	$2a$08$GLGbSWyM0/9vKzhynEessee2We0Fenh8Gi6It.x2lLvY66FrM5RCG	\N	sdfbhm	94121212	Admin
aakash	hirnotia	a@f	etrdtfyguh	2020-04-06 20:05:24.357311	170	2020-04-06 20:05:24.357311	$2a$08$p.p22Arzrb05lkHieShEUO6NqD16lWny50/IC8WzwSJTPogZkTCXm	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAZiIsImV4cCI6MTU4NzQ4NTE2NX0.Yb6834l5-lfC9AjC0kt6TqsHund7e7YdfcnZL-wsb4I	Active	9258656829	\N
asdfghj	ssdfghj	hfgdjj@ggvjg	fgjgh	2020-04-09 18:22:04.282123	197	2020-04-09 18:22:04.282124	$2a$08$QVJEgq8zhxPYrHG6bYOB9.P.i5AsWvY1QaT/JRz23IhInBYZK72kS	\N	Active	gcfgvhj	Accountant
z	ghjk	n@sdkalim	s	2020-04-09 12:26:14.555658	174	2020-04-09 12:26:14.555658	$2a$08$7us5YCW2QC3W4WdPWrZ2nu3AeNUe8fMsI8q.F6WFgI2feAJChDwFq	\N	sdfbgnhm	94112121212	\N
z	ghjk	np@sdkalim	s	2020-04-09 12:27:35.515558	176	2020-04-09 12:27:35.515558	$2a$08$r77le6tMFzcNUOiQALLg5uC5.yCbxTHur/SWyGueK6Ii7l/TG8hIq	\N	sdfbgnhm	94112121212	\N
vkgoyal	goyal	vgoyal@j	asbhscvj	2020-04-06 21:47:33.145885	171	2020-04-06 21:47:33.145885	$2a$08$GijssIzO5YKEf7gEyLlliuQF2qwWtHY0qcBK404n30wX3gYD3mDG.	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZnb3lhbEBqIiwiZXhwIjoxNTg3MDkwMzE4fQ.__ADOMKyHofZt4TL1ng4aSFPjBFWcChp6OlN3bizzuE	Deleted	456789255	\N
vinay	kumar	a@erty	fjf	2020-04-06 23:54:23.783927	172	2020-04-06 23:54:23.783927	$2a$08$2AWPjq/d9S5bCH5NfVVdEe7bQUBdjXV6aCJwrUzGwst7sV5dspDm6	\N	Inactive	12345678	\N
vinay	ashish	v@goyaljgj	sdfghj	2020-04-09 14:49:06.855978	189	2020-04-09 14:49:06.855979	$2a$08$9mxXssayLfjFGd4S3aXyYO.DT3zfCiRUA6eVmtUeyd2Ao3KmLtrci	\N	Inactive	7894561235	Customer Care
z	ghjk	np@sdcfddddfckalim	s	2020-04-09 12:30:33.583335	177	2020-04-09 12:30:33.583335	$2a$08$amHkKxFTmTCMIKiD.2AMpuNjuYaWdO9EILDnNoboAlZPTIUcRPEWG	\N	sdfbgnhm	94112121212	accountant
owaish	kalim	np@sfddfckalim	s	2020-04-09 13:39:30.551052	178	2020-04-09 13:39:30.551052	$2a$08$Ho1VZrpDxP5ajYotr/3VQOsj4KBr8AKxeqrt/LozxUl.D/CaTzbMu	\N	sdfbhm	94121212	admin
aghvgvjg	gvghvgh	hfc@hvg	ffc	2020-04-09 15:08:34.814264	190	2020-04-09 15:08:34.814264	$2a$08$K5NbJMf73uABWasUQ0tb5efYA4ae.N/V7nEBi/.W4Gc65b9xMqtMy	\N	Inactive	hgghv	Customer Care
owaish	kalim	np@sfddfjhjhjckalim	s	2020-04-09 14:33:35.765078	180	2020-04-09 14:33:35.765078	$2a$08$dH.Ezopc0x0Q6tM0O6koK.Fuarvo70ZqUF/birDo2EzvRcr3DUXeO	\N	sdfbhm	94121212	admin
owaish	kalim	np@sfddfjhjhjcddkalim	s	2020-04-09 14:34:45.690596	181	2020-04-09 14:34:45.690596	$2a$08$8szFCR9vCFIc3Aqq6HRKnOWY5/88snYHjdhza43YZL53NTiHypjB.	\N	sdfbhm	94121212	admin
owaish	kalim	np@kllalim	s	2020-04-09 14:35:11.138258	182	2020-04-09 14:35:11.138258	$2a$08$Zz6SVw5PctywgPCZ6.rmSuk8QNjaZOgTlyWQD/dQUbcFDFPgjfkLO	\N	sdfbhm	94121212	admin
owaish	kalim	np@klim	s	2020-04-09 14:36:40.066848	185	2020-04-09 14:36:40.066848	$2a$08$.FWDqvQJEwMEoTEpggmyBOd3Q8j88bWCd3Q1bFh2kFRsMmfImhYwq	\N	sdfbhm	94121212	Admin
aajka	gftff	a@hgfghh	sdfgyu	2020-04-09 15:18:25.656217	193	2020-04-09 15:18:25.656218	$2a$08$z9xz53HTkj0ZZAgqWb4IrObh5KabyhMbV5x7a1pqWx8I0tyBk4h7C	\N	Active	2345678	Admin
12345678	esdyuj	s@fh	gvgvg	2020-04-09 15:20:15.389856	194	2020-04-09 15:20:15.389856	$2a$08$QVhIY3hoIMixi48elmUsGum596rYH0sb25GBSJdwYEFXTEk3aJ93.	\N	Active	697653	Admin
svbn	fgbhnjm	bhgb@jhhj	rfvev	2020-04-09 15:21:21.13906	195	2020-04-09 15:21:21.13906	$2a$08$UlHZOOrn54efiw.BG7EZxerxiSrLWzKvAiXlbM5iY2KrpsKhHIkfG	\N	Inactive	gbhn	Accountant
abhishek	kumar	dfds@nvv	greg	2020-04-09 20:13:22.809635	204	2020-04-09 20:13:22.809635	$2a$08$bZoSVNxJFF3FxV/Nr9UgzO75pnSeuudOlEBw9fWji75yuNpVRtPHW	\N	Inactive	sdfghj	Accountant
owaish	kkalimllll	a@owaish	ghvgh	2020-04-09 18:45:39.19051	198	2020-04-09 18:45:39.190511	$2a$08$k0Xm/Ej19k.B.CewVv98CeVzZ9kHdzeTowSswesSUnKc73n.wuDn2	\N	Active	88945562	Accountant
ABD	KHAN	np@mqqqahalodum	s	2020-04-09 18:50:17.800002	199	2020-04-09 18:50:17.800002	$2a$08$x2oMufuPOjCFRtstSn6rFuab7LvDAB8CmxsozpA9NoVpyXDO64joy	\N	sdfbhm	94121212	Admin
ABD	KHAN	np@mqqqahadum	hhs	2020-04-09 18:51:30.303936	200	2020-04-09 18:51:30.303936	$2a$08$eMW.4SyZJYnixdkCm1Ct..oK8Rpin8lgqCaVLiDVUPVYgq1MMbt1G	\N	sdfbhm	94121212	Admin
ABD	KHAN	np@mqqqahaoosodum	hhs	2020-04-09 18:52:41.931823	202	2020-04-09 18:52:41.931823	$2a$08$JF1o1L/uuc8tYrORdnFoN.4GQ1BHfndNjXy1mKEBzNSqKpBW8VHCu	\N	sdfbhm	94121212	Admin
ayush	sdfghj	efewf@jhbjh	sdgsgvfd	2020-04-09 18:52:56.861374	203	2020-04-09 18:52:56.861374	$2a$08$TzaalUBBDWVO8WSoABQOZeFiKyO/.5UzwICQncaIJ8aWn8sQzjQgC	\N	Inactive	sdfghefewf	Accountant
John	Wick	np@johnaoosodum	hhs	2020-04-09 23:09:52.93417	205	2020-04-09 23:09:52.93417	$2a$08$9pt7v29QzIBXBTq4KrhPDuPFe.nL2USvLLSISstUCmslZp5um0u16	\N	sdfbhm	94121212	Accountant
John	Wick	np@johnaohnosodum	hhs	2020-04-10 10:56:05.138525	206	2020-04-10 10:56:05.138526	$2a$08$Z/oiiVa4HAeB6umKQxvJD.I1SG6R1TUI94ccqgJJsMOd6gV0Kcst.	\N	sdfbhm	94121212	Admin
owaish	gheqwesdfgrtyem	owaish@kalpppim1	1341232qwertbnbnmyuil	2020-04-06 01:45:08.574426	164	2020-04-06 01:45:08.574426	owashpp	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzcXdoQHNka2FsaW0iLCJleHAiOjE1ODc0ODgyOTh9.jCwbjRlBTCVg7hRg8cFspDjmVe0S91lwYIH0HcG1HUg	sdfgh	123494112121234fdf1212	\N
sdfghjk	ghj	dgh@jhg	fe	2020-04-10 16:38:54.725265	207	2020-04-10 16:38:54.725266	$2a$08$TvfY3lR.fMLFFfMKg7gp0eyFKAawNXDA.ludCA1n0Vwe5QaR3sWWq	\N	Inactive	sdfgh	Customer Care
dfghj	ddfgh	hjghj@jgj	gv	2020-04-10 17:18:09.019004	208	2020-04-10 17:18:09.019004	$2a$08$iYBnmWWxwGZH9D0PxilNIOudBJbltWCjftYm27aXjF9l8r3LAmGyC	\N	Inactive	vvv	Accountant
John	Wick	np@lllum	hhs	2020-04-11 12:33:05.495498	218	2020-04-11 12:33:05.495498	$2a$08$wVDVSrhmssn4aS99Gl53TOUb4VeX/m/J9JFTu5AN6w3c7ETyeTF7K	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5wQGxsbHVtIiwiZXhwIjoxNTg3NDkwNjM2fQ.L1QIhORHNC0F-GB8qgzrn9GoGlRvU26_Xfk59w8-8Hk	sdfbhm	94121212	Admin
John	Wick	np@llsodum	hhs	2020-04-11 00:42:57.652819	209	2020-04-11 00:42:57.652819	$2a$08$UckSvwtWPOUvwBljYPvSaOlk8u4pwuT6/0Bg.SwMc7k7DQv7353.G	\N	sdfbhm	94121212	Admin
John	Wick	np@llsodlllum	hhs	2020-04-11 00:44:02.307898	211	2020-04-11 00:44:02.307898	$2a$08$LYDXkGLR6P57oLzhazOOfOxXgkR7Risahab1qNbqfHV.toOCueGMy	\N	sdfbhm	94121212	Admin
John	Wick	np@llsobbvdlllum	hhs	2020-04-11 00:45:12.34403	213	2020-04-11 00:45:12.344031	$2a$08$bNJppTIseeI1hN39cr9aqe6iI.S10.My//4INmcqloS7tuMKNrjdW	\N	sdfbhm	94121212	Admin
John	Wick	np@llsobwertyudlllum	hhs	2020-04-11 00:51:01.930495	214	2020-04-11 00:51:01.930495	$2a$08$UUpiFc5UNONg7hD477spuOetLBRqzKrqSw7fgYZJ.JhfFUIRjotdK	\N	sdfbhm	94121212	Admin
John	Wick	np@llsobweqazyudlllum	hhs	2020-04-11 00:55:26.335267	215	2020-04-11 00:55:26.335268	$2a$08$WW8gAIaz2UzADBJuZTF.5.rE.bogqK/iUYDliJRbMCskYSGNkUQxG	\N	sdfbhm	94121212	Admin
John	Wick	np@llsobwedlllum	hhs	2020-04-11 00:59:06.51994	216	2020-04-11 00:59:06.519941	$2a$08$.a.CkZ9YRhtNfaIEENoa7uaDooQw230Cf/H6QO9jY33J4igHfhE92	\N	sdfbhm	94121212	Admin
\.


--
-- Data for Name: slh_transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_transactions ("Customer_Order_Id", "Customer_Id", "Therapist_Gender", "Massage_For", "Slot_Time", "Slot_Date", "Customer_Address", "Pincode", "Transaction_Mode", "Bank_Type", "Merchant_Transaction_Id", "Massage_Duration", "Customer_Name", "CreatedAt", "Total_Order_Amount", "Customer_Souls_Id", "Payment_Gateway_Id", "Latitude", "Longitude", "Number_Of_Therapist_Required", "Payment_Gateway_Mode") FROM stdin;
\N	38	male	cbn	2020-04-07 22:58:54.26731	2020-04-07	ertyj	9854334	dfghj	cvbn	158628053438	trew	kalim	2020-04-07 22:58:54.26731	10000	2020040738	sdfghj	qwerty	asdfg	2	dfgh
18	38	male	cbn	2020-04-07 22:58:54.26731	2020-04-07	ertyj	9854334	dfghj	cvbn	158628053438	trew	kalim	2020-04-07 22:58:54.26731	10000	2020040738	sdfghj	qwerty	asdfg	2	dfgh
19	39	male	cbn	2020-04-08 01:05:35.177064	2020-04-08	ertyj	9854334	dfghj	cvbn	158628813539	trew	kalim	2020-04-08 01:05:35.177064	10000	2020040839	sdfghj	qwerty	asdfg	2	dfgh
19	39	male	cbn	2020-04-08 01:05:35.177064	2020-04-08	ertyj	9854334	dfghj	cvbn	158628813539	trew	kalim	2020-04-08 01:05:35.177064	10000	2020040839	sdfghj	qwerty	asdfg	2	dfgh
20	40	male	cbn	2020-04-08 01:06:47.71179	2020-04-08	ertyj	9854334	dfghj	cvbn	158628820740	trew	kalim	2020-04-08 01:06:47.71179	10000	2020040840	sdfghj	qwerty	asdfg	2	dfgh
20	40	male	cbn	2020-04-08 01:06:47.71179	2020-04-08	ertyj	9854334	dfghj	cvbn	158628820740	trew	kalim	2020-04-08 01:06:47.71179	10000	2020040840	sdfghj	qwerty	asdfg	2	dfgh
20	40	male	cbn	2020-04-08 01:06:47.71179	2020-04-08	ertyj	9854334	dfghj	cvbn	158628820740	trew	kalim	2020-04-08 01:06:47.71179	10000	2020040840	sdfghj	qwerty	asdfg	2	dfgh
20	40	male	cbn	2020-04-08 01:06:47.71179	2020-04-08	ertyj	9854334	dfghj	cvbn	158628820740	trew	kalim	2020-04-08 01:06:47.71179	10000	2020040840	sdfghj	qwerty	asdfg	2	dfgh
21	40	male	cbn	2020-04-08 15:59:05.800142	2020-04-08	ertyj	9854334	debit	sbi	158634174540	trew	kalim	2020-04-08 15:59:05.800141	10000	2020040840	sherlock	qwerty	asdfg	2	cheque
21	40	male	cbn	2020-04-08 15:59:05.800142	2020-04-08	ertyj	9854334	debit	sbi	158634174540	trew	kalim	2020-04-08 15:59:05.800141	10000	2020040840	sherlock	qwerty	asdfg	2	cheque
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, "Name", "Email", "Password", "Token", "CreatedAt", "UpdatedAt", "Role") FROM stdin;
\.


--
-- Name: slh_customers_Customer_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_Customer_Id_seq"', 48, true);


--
-- Name: slh_customers_pending_orders_Customer_Primary_Order_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq"', 22, true);


--
-- Name: slh_partners_Partner_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_partners_Partner_Id_seq"', 5, true);


--
-- Name: slh_roles_RoleId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_roles_RoleId_seq"', 1, false);


--
-- Name: slh_roles_Role_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_roles_Role_Id_seq"', 6, true);


--
-- Name: slh_team_has_role_Team_Has_Role_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_team_has_role_Team_Has_Role_Id_seq"', 1, true);


--
-- Name: slh_teams_TeamId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_teams_TeamId_seq"', 219, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 28, true);


--
-- Name: slh_partners Partner_Email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_partners
    ADD CONSTRAINT "Partner_Email" UNIQUE ("Partner_Email");


--
-- Name: slh_roles RoleName; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_roles
    ADD CONSTRAINT "RoleName" UNIQUE ("Role_Name");


--
-- Name: slh_customers slh_customers_Customer_Mobile_No_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers
    ADD CONSTRAINT "slh_customers_Customer_Mobile_No_key" UNIQUE ("Customer_Mobile_No");


--
-- Name: slh_customers_pending_orders slh_customers_pending_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers_pending_orders
    ADD CONSTRAINT slh_customers_pending_orders_pkey PRIMARY KEY ("Customer_Order_Id");


--
-- Name: slh_customers slh_customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers
    ADD CONSTRAINT slh_customers_pkey PRIMARY KEY ("Customer_Id");


--
-- Name: slh_partners slh_partners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_partners
    ADD CONSTRAINT slh_partners_pkey PRIMARY KEY ("Partner_Id");


--
-- Name: slh_roles slh_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_roles
    ADD CONSTRAINT slh_roles_pkey PRIMARY KEY ("Role_Id");


--
-- Name: slh_teams slh_teams_Email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams
    ADD CONSTRAINT "slh_teams_Email_key" UNIQUE ("Email");


--
-- Name: slh_teams slh_teams_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_teams
    ADD CONSTRAINT slh_teams_pkey PRIMARY KEY ("TeamId");


--
-- Name: users users_Email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "users_Email_key" UNIQUE ("Email");


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: fki_slh_customers_pending_orders_id_fkey; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_slh_customers_pending_orders_id_fkey ON public.slh_customers_pending_orders USING btree ("Customer_Id");


--
-- Name: slh_customers_pending_orders slh_customers_pending_orders_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers_pending_orders
    ADD CONSTRAINT slh_customers_pending_orders_id_fkey FOREIGN KEY ("Customer_Id") REFERENCES public.slh_customers("Customer_Id") NOT VALID;


--
-- Name: slh_team_has_role slh_team_has_role_TeamId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_team_has_role
    ADD CONSTRAINT "slh_team_has_role_TeamId_fkey" FOREIGN KEY ("Team_Id") REFERENCES public.slh_teams("TeamId") NOT VALID;


--
-- PostgreSQL database dump complete
--

