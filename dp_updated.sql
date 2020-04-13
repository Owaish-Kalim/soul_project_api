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
-- Name: slh_communication_templates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_communication_templates (
    "Comm.Template_Id" integer NOT NULL,
    "Comm.Template_Type" text NOT NULL,
    "Trigger_For" text NOT NULL,
    "SMS_Content" text NOT NULL,
    "Email_Content" text NOT NULL,
    "Subject" text NOT NULL,
    "Trigger_Time" text NOT NULL
);


ALTER TABLE public.slh_communication_templates OWNER TO postgres;

--
-- Name: slh_communication_templates_Comm.Template_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_communication_templates_Comm.Template_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_communication_templates_Comm.Template_Id_seq" OWNER TO postgres;

--
-- Name: slh_communication_templates_Comm.Template_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_communication_templates_Comm.Template_Id_seq" OWNED BY public.slh_communication_templates."Comm.Template_Id";


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
    "Pincode" integer NOT NULL,
    "Status" boolean
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
    "Massage_Duration" text NOT NULL,
    "CreatedAt" timestamp without time zone NOT NULL,
    "Customer_Souls_Id" text,
    "Pincode" integer,
    "Is_Order_Confirmed" boolean,
    "Customer_Name" text,
    "Total_Order_Amount" integer,
    "Merchant_Transaction_Id" text,
    "Location" point
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
    "Partner_Email" text,
    "Latitude" text,
    "Longitude" text,
    "Onboard_Date" timestamp without time zone,
    "CreatedAt" timestamp without time zone,
    "UpdatedAt" timestamp without time zone,
    "Last_Updated_By" text,
    "Commission_Type" text,
    "CreatedBy" text,
    "Partner_Mobile_No" text,
    "Pincode" text,
    "Per_Visit_Price_Commission" text
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
-- Name: slh_souls_settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_souls_settings (
    "Email_Settings" text,
    "SMS_Settings" text
);


ALTER TABLE public.slh_souls_settings OWNER TO postgres;

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
    "TeamId" integer NOT NULL,
    "Password" text NOT NULL,
    "Token" text,
    "Status" text NOT NULL,
    "MobileNo" text NOT NULL,
    "Role" text,
    "CreatedAt" text,
    "JoiningDate" text
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
-- Name: slh_communication_templates Comm.Template_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_communication_templates ALTER COLUMN "Comm.Template_Id" SET DEFAULT nextval('public."slh_communication_templates_Comm.Template_Id_seq"'::regclass);


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
-- Data for Name: slh_communication_templates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_communication_templates ("Comm.Template_Id", "Comm.Template_Type", "Trigger_For", "SMS_Content", "Email_Content", "Subject", "Trigger_Time") FROM stdin;
1	abcd	customers	popopo	opopopop	Owaish	ten-min
\.


--
-- Data for Name: slh_customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers ("Customer_Id", "Customer_Email", "CreatedAt", "Customer_Gender", "Customer_Name", "Customer_Address", "Registered_Source", "Last_Access_Time", "Customer_Souls_Id", "Customer_Mobile_No", "Pincode", "Status") FROM stdin;
39	kalim@1234	2020-04-08 01:05:21.630388	poiuytrdes	kalim	ertyj	website	2020-04-08 01:05:21.630388	2020040839	172345615	9854334	\N
40	kalim@1234	2020-04-08 01:06:39.704405	poiuytrdes	kalim	ertyj	website	2020-04-08 01:06:39.704405	2020040840	1725615	9854334	\N
34	owaish@1234	2020-04-07 16:44:13.338694	male	ashish	ertyj	sdfghj	2020-04-07 16:44:24.024427	2020040734	9876543	123	\N
41	owaish@1234	2020-04-08 10:51:30.046757	male	ayush	ertyj	sdfghj	2020-04-08 10:51:30.046757	2020040841	987656543	123	\N
36	owaish@1234	2020-04-07 18:43:46.476663	male	laptop	ertyj	sdfghj	2020-04-07 19:37:19.496711	2020040736	956543	123	\N
42	kalim@1234	2020-04-08 14:26:05.92923	poiuytrdes	kumar	ertyj	website	2020-04-08 14:26:05.92923	2020040842	121	9854334	\N
38	owaish@1234	2020-04-07 22:57:58.287051	male	laptop	ertyj	sdfghj	2020-04-07 22:57:58.287052	2020040738	5335	123	\N
43	kalim@1234	2020-04-08 21:31:13.75261	poiuytrdes	kumar	ertyj	website	2020-04-08 21:31:13.75261	20200408-43	12561	9854334	\N
44	kalim@1234	2020-04-08 21:31:34.472219	poiuytrdes	kumar	ertyj	website	2020-04-08 21:31:34.472219	20200408-44	1256155555	9854334	\N
45	kalim@1234	2020-04-08 21:43:13.069007	poiuytrdes	kumar	ertyj	website	2020-04-08 21:43:13.069007	20200408-45	125615545	9854334	\N
46	kalim@1234	2020-04-08 22:07:26.431891	poiuytrdes	kumar	ertyj	website	2020-04-08 22:07:26.431891	20200408-46	125545	9854334	\N
47	Johnw@1234	2020-04-09 23:06:05.839646	Male	John-Wick	Ny	website	2020-04-12 21:40:17.713669	2020040947	999	135	\N
49	Johnw@1234	2020-04-12 21:54:50.821133	Male	John-Wick	Ny	website	2020-04-12 21:54:50.821133	2020041249	12126749	135	t
50	Johnw@1234	2020-04-12 21:58:38.603594	Male	John-Wick	Ny	website	2020-04-12 21:58:38.603594	2020041250	129	135	t
\.


--
-- Data for Name: slh_customers_pending_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers_pending_orders ("Customer_Order_Id", "Customer_Id", "Number_Of_Therapists_Required", "Therapist_Gender", "Massage_For", "Slot_Time", "Slot_Date", "Customer_Address", "Massage_Duration", "CreatedAt", "Customer_Souls_Id", "Pincode", "Is_Order_Confirmed", "Customer_Name", "Total_Order_Amount", "Merchant_Transaction_Id", "Location") FROM stdin;
9	36	5	zxcvb	zxcvbn	2020-04-07 19:37:32.242377	2020-04-07 19:37:32.242377	wqewrtydfghjk	asdfghj	2020-04-07 19:37:32.242376	2020040736	987654	f	owasih	10000	\N	\N
10	36	5	zxcvb	zxcvbn	2020-04-07 21:33:13.790638	2020-04-07 21:33:13.790637	wqewrtydfghjk	asdfghj	2020-04-07 21:33:13.790637	2020040736	987654	f	owasih	10000	158627539336	\N
11	36	5	zxcvb	zxcvbn	2020-04-07 21:34:24.432404	2020-04-07 21:34:24.432404	wqewrtydfghjk	asdfghj	2020-04-07 21:34:24.432403	2020040736	987654	f	owasih	10000	158627546436	\N
12	36	5	zxcvb	zxcvbn	2020-04-07 21:34:36.322148	2020-04-07 21:34:36.322148	wqewrtydfghjk	asdfghj	2020-04-07 21:34:36.322148	2020040736	987654	f	owasih	10000	158627547636	\N
13	36	5	zxcvb	zxcvbn	2020-04-07 22:27:27.526831	2020-04-07 22:27:27.526831	wqewrtydfghjk	asdfghj	2020-04-07 22:27:27.526831	2020040736	987654	f	owasih	10000	158627864736	\N
14	38	2	male	cbn	2020-04-07 22:58:35.691487	2020-04-07 22:58:35.691486	ertyj	trew	2020-04-07 22:58:35.691486	2020040738	9854334	f	kalim	10000	158628051538	\N
15	38	2	male	cbn	2020-04-07 22:58:41.269672	2020-04-07 22:58:41.269672	ertyj	trew	2020-04-07 22:58:41.269672	2020040738	9854334	f	kalim	10000	158628052138	\N
16	38	2	male	cbn	2020-04-07 22:58:44.053138	2020-04-07 22:58:44.053138	ertyj	trew	2020-04-07 22:58:44.053138	2020040738	9854334	f	kalim	10000	158628052438	\N
17	38	2	male	cbn	2020-04-07 22:58:48.846522	2020-04-07 22:58:48.846522	ertyj	trew	2020-04-07 22:58:48.846522	2020040738	9854334	f	kalim	10000	158628052838	\N
18	38	2	male	cbn	2020-04-07 22:58:54.26731	2020-04-07 22:58:54.26731	ertyj	trew	2020-04-07 22:58:54.26731	2020040738	9854334	f	kalim	10000	158628053438	\N
19	39	2	male	cbn	2020-04-08 01:05:35.177064	2020-04-08 01:05:35.177064	ertyj	trew	2020-04-08 01:05:35.177064	2020040839	9854334	f	kalim	10000	158628813539	\N
20	40	2	male	cbn	2020-04-08 01:06:47.71179	2020-04-08 01:06:47.71179	ertyj	trew	2020-04-08 01:06:47.71179	2020040840	9854334	f	kalim	10000	158628820740	\N
21	40	2	male	cbn	2020-04-08 15:59:05.800142	2020-04-08 15:59:05.800142	ertyj	trew	2020-04-08 15:59:05.800141	2020040840	9854334	f	kalim	10000	158634174540	\N
22	40	2	male	cbn	2020-04-10 15:05:25.65068	2020-04-10 15:05:25.65068	ertyj	trew	2020-04-10 15:05:25.65068	2020040840	9854334	f	kalim	10000	1586511325-40	\N
\.


--
-- Data for Name: slh_partners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_partners ("Partner_Id", "Partner_Name", "Partner_Gender", "Partner_Address", "Partner_Email", "Latitude", "Longitude", "Onboard_Date", "CreatedAt", "UpdatedAt", "Last_Updated_By", "Commission_Type", "CreatedBy", "Partner_Mobile_No", "Pincode", "Per_Visit_Price_Commission") FROM stdin;
4	owaishkalim	male	poiu	owa@12	LL	pp	2020-04-10 21:15:57.539975	2020-04-10 21:15:57.539975	2020-04-10 21:15:57.539975	ll	oiioi	uu	\N	\N	\N
5	owaishkalimlll	male	poiu	owa@1hhh2	LL	pp	2020-04-10 23:27:25.090668	2020-04-10 23:27:25.090668	2020-04-10 23:27:25.090668	ll	oiioi	uu	\N	\N	\N
6	owaishkalimlll	male	poiu	owa@1hhqqqh2	LL	pp	2020-04-12 22:20:52.8123	2020-04-12 22:20:52.8123	2020-04-12 22:20:52.8123	ll	oiioi	uu	\N	\N	\N
7	owaishkalimlll	male	poiu	owa@1hmmhqqqx	LL	pp	2020-04-12 23:55:17.507001	2020-04-12 23:55:17.507001	2020-04-12 23:55:17.507001	ll	oiioi	uu	12467345	23456789	3456789
8	owaishkalimlll	male	poiu	owa@1hmqmhqqqx	LL	pp	2020-04-13 00:18:35.936433	2020-04-13 00:18:35.936433	2020-04-13 00:18:35.936433	ll	oiioi	uu	12467345	23456789	3456789
9	Dhanprakash Hirnotia	Female	09084768046	ashutoshhhirnotia23@gmail.com	adad	dad	2020-04-13 00:22:00.624653	2020-04-13 00:22:00.624653	2020-04-13 00:22:00.624654	adad	Percentage(%)	dad	dad	245101	dad
3	buzurgbuddhaashish	male	poiu	ow@12	LL	pp	2020-04-10 18:55:21.808417	2020-04-10 18:55:21.808417	2020-04-10 18:55:21.808417	ll	oiioi	uu	1095	10909	190
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
-- Data for Name: slh_souls_settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_souls_settings ("Email_Settings", "SMS_Settings") FROM stdin;
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
1	2020-04-12 12:38:54.396016	2020-04-12 12:38:54.44396	220	sdfbhm	John	Wick
1	2020-04-12 12:47:44.732189	2020-04-12 12:47:44.764121	221	sdfbhm	John	Wick
2	2020-04-12 13:06:30.347297	2020-04-12 13:06:30.411562	225	Inactive	Dhanprakash	Hirnotia
1	2020-04-12 13:11:42.997092	2020-04-12 13:11:43.04015	226	sdfbhm	John	Wick
1	2020-04-12 16:06:40.822677	2020-04-12 16:06:41.007063	230	inactive	Ow	Kal
3	2020-04-12 15:54:34.751474	2020-04-12 15:54:35.305042	229	Active	Ashish	Hirnotia
3	2020-04-12 15:43:01.40102	2020-04-12 15:43:01.675147	228	inactive	Ow	Kal
1	2020-04-12 13:12:03.234442	2020-04-12 13:12:03.675693	227	Inactive	aakash	sa
1	2020-04-12 16:58:09.798263	2020-04-12 16:58:10.387982	231	inactive	ABD	KHAN
2	2020-04-12 21:45:48.808732	2020-04-12 21:45:48.930373	232	Inactive	ashutosh	hirnotia
2	2020-04-12 21:48:38.87997	2020-04-12 21:48:39.458355	233	Deleted	aakash	asdfghjkl
1	2020-04-12 22:06:06.449717	2020-04-12 22:06:06.668484	235	inactive	ABD	KHAN
1	2020-04-13 21:11:38.752652	2020-04-13 21:11:38.921357	237	inactive	ABD	KHAN
\.


--
-- Data for Name: slh_teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_teams ("FirstName", "LastName", "Email", "Address", "TeamId", "Password", "Token", "Status", "MobileNo", "Role", "CreatedAt", "JoiningDate") FROM stdin;
AShish	kumar	ak@asf.as	Qwerty	141	$2a$08$tsYFF7bP/noieJBR14.rWeXMR5Gg9bUvkYQFHWz7xXLdQSW4kgE8m	\N	Inactive	234567898	\N	\N	\N
Ashish	Singh	asingh2@ch.iitr.ac.in	Rajiv Bhwan, IIT Roorkee	120	$2a$08$Odz87MnbLSN4evI12sVqPecm.YRBRrPhi1PSUooU0onWmBB3BJAHa	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzaW5naDJAY2guaWl0ci5hYy5pbiIsImV4cCI6MTU4Njg4NzI0OH0.b1S7c_-yt2QIXYwRfSsQe0JaJ0F_HCby5ObiTcQQIcY	Inactive	+918769304216	Accountant	\N	\N
saSdfa	test	sh@gmail.com	dfghjk	114	$2a$08$XWC2NdF8/5a4fv9uiBvD2ujxHNY9laSh89ERnbJ.W.cxPOdfjc0JK	\N	active	94112121212	admin	\N	\N
Kartik 	Aryan	k@kk.kk	qwertyui	142	$2a$08$r1N2fi5YGNgGKUgelFBA7u.os5AkqHe.YaJRG5VCKmC3fG1TF0PmC	\N	Inactive	23456789	\N	\N	\N
sdvf	sdfcvbn	as@gmail.com	v	139	$2a$08$oNfiL1o0bskZJZQJFuHaZu3Kev9kBGS3E.5VVFpSXIASut7vkT2KK	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtYWlsLmNvbSIsImV4cCI6MTU4NjkxMzgxOH0.0XI66C4bCqeC7YMUvxe4akls1ECjkwo875kZM6nalNU	accountant	94112121212	\N	\N	\N
AShish	kumar	sd@af.af	asdgfhgjk	115	$2a$08$mHcOXlXlOaAdv5rsZ9O8B.8LyBm/zwwbQXzr11Gt5kICor/SoCRGW	\N	inactive	1234567789	admin	\N	\N
AShish	kumar	sds@af.af	asdgfhgjk	116	$2a$08$EfOW8CYMoHX3rQIgcWEsROUVCndwe2yHElpTzJ07ZAD1OkVcolCRW	\N	inactive	1234567789	admin	\N	\N
	cvbn	sh@gghjsdertfvghmail.com	v	125	$2a$08$grUDESzu7oc4kPvGn3zFmuPh6RLoIVq1Orwu19RtKCEP5K8VIu54C	\N		94112121212		\N	\N
saSdfa	test	sh@gghjmail.com	dfghjk	117	$2a$08$KsiRMhJslypxxptG3H9AM.9dvEewCFZE7SrS/Wk3GuJkpFpKrPiti	\N	active	94112121212	admin	\N	\N
bn	test	sh@gghjvghmail.com	v	119	$2a$08$1kG4naPdgil9YQW3WhrdfuxS64HJz.wtEqtG4Ow6zJ7YepfIYiE06	\N	fgh	94112121212	admin	\N	\N
sdvf	sdfcvbn	sh@gghjsdertfvail.com	v	127	$2a$08$XVV/PmNuQNFyKLK0bFdu9uw2nL1E6EqLphlzRRrnjitT.vjp1uYmO	\N	sdfbgnhm	94112121212	zxcv 	\N	\N
sdvf	sdfcvbn	sh@gghjsdertfvartyil.com	v	128	$2a$08$PHaTWixqJLnXxufOGSsEIeYPc2iGSe7IGKqdc7vuonfi99YVIK1ZW	\N	sdfbgnhm	94112121212	zxcv 	\N	\N
sdvf	sdfcvbn	sh@gghjsdefrgthsdertfvartyil.com	v	130	$2a$08$CqJKFGCslUEL4v0Vt.z.rOLtIbk7mOHOVtiU8uF2GWSh3SN11RJBS	\N	sdfbgnhm	94112121212	zxcv	\N	\N
asfa	test	owaish@gmail.com	dfghjk	113	$2a$08$xQ7ctow8rGKAUrcRWw.LBetRFGK.GwHOU2M7GM.jcKfpchDDNhXx2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzaEBnbWFpbC5jb20iLCJleHAiOjE1ODY5MTI5MDV9.zyfDAqPo4AdYNo-V2kWB6gmEWXLR7t3--H9atmOczkM	accountant	94112121212	admin	\N	\N
bn	cvbn	sh@gghjsdfvghmail.com	v	123	$2a$08$XIZ./xeRmlr55ZDqHbfXF.fPCuJRLcHW3uCjJycv0FPYB0PB6Hk8.	\N	fgh	94112121212	admin	\N	\N
AShish	kumar	a@a.cc	asdfghj	131	$2a$08$bgExze5662U3E8v3pPtlE.AmoJ2WDIQkur7135857Z5M1..8.qdxS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYS5jYyIsImV4cCI6MTU4NjkwMjU0MX0.ruT9B7wYQXksugTi6FCSggdSqVvXk_I5H9bbiI0WKo4	Active	1234567894	Admin	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfvdefrgthsdertfvartyil.com	v	134	$2a$08$sc1O84ampL5zG/A2yLggq.6UNosQlxINgZPS3C4KOYrBi7aduJuPi	\N	sdfbgnhm	94112121212	zxcv	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfasdfgvdefrgthsdertfvartyil.com	v	135	$2a$08$.q5yPJhyOto60AZdWjNrjemK2VxkgxN0NrAJGI8cFbDt1GxqYs9sG	\N	sdfbgnhm	94112121212	zxcv	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfwertyasdfgvdefrgthsdertfvartyil.com	v	137	$2a$08$rv11WiK/voyPw7aqH588NeELy/DW6psUv.V7TYDR.DDaISJMfZ0Zi	\N	sdfbgnhm	94112121212	zxcv	\N	\N
AShish	kumar	ashutoshhhirnotia23@gmail.com	Gali no 16 B,Lajjapuri ,chamri road ,Hapur	140	$2a$08$XDRmXktAKpz2UFT1QzgXjeKvzvtRGW1iAQS/WjNYVZkCgkf3SnuCe	\N	Inactive	09084768046	\N	\N	\N
sdvf	sdfcvbn	as@gmvbnail.com	v	145	$2a$08$JfZvVtQjv4asM1vBYhU3AewUtysWC75fgCsMGasJPDviR7XDAHVBi	\N	sdfbgnhm	94112121212	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbnail.com	v	148	$2a$08$6GwLCpmKrIjIpT1ok27Aw.5VzV.WcqMwYWpVkox7fjqWk5OfITjD2	\N	sdfbgnhm	94112121212	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbbnail.com	v	152	$2a$08$u6u0N.HdLZBDRY3IE2zIRuuuYufeiGrqcglXeBHxz9fkX/G5zGWJq	\N	sdfbgnhm	94112121212	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbbxcnail.com	jv	155	$2a$08$isp0ldIIBQD5MeiZNgkkgelFa88Lch/PVBpGYhUpFa6.6/DXI.3eG	\N	sdfbgnhm	94112121212	\N	\N	\N
sdvf	ghjk	as@gmvbvbbbnmxcnail.com	jv	158	$2a$08$pSK3bO90lI4qQg.VV1bV6.F694BjTelTDyjYCMZnFwcrKGLwUBaju	\N	sdfbgnhm	94112121212	\N	\N	\N
owaish	kalim	np@kliakashlodum	s	186	$2a$08$BeWp.U7AhqE0hQnvJpQ1EODfU8c25Ehw6JLEmsw6Mr0ijbf/mN3ca	\N	sdfbhm	94121212	Admin	\N	\N
ayush	bhai	ayush@123456	rtdt	196	$2a$08$xHj/41rSidIlKdiFD8hymOSMbqDKGsuDeEwWbfx14bePeyrTwq/Gm	\N	Inactive	7894561235	Accountant	\N	\N
sdvf	ghjk	owaish@kalim	jv	161	$2a$08$4gIPpUW6mHOPJtXqfyg3.ecAKUidtoaMtOTn0.LCts.g3XuL8Rc2q	\N	sdfbgnhm	94112121212	\N	\N	\N
owaish	gheem	a@a.aa	12qwertbnbnmyuil	157	$2a$08$l.jy1pudS6lWL/Q.4/ctjOj0sGqYz.OoigzgvRanMb1/QiNast0Ye	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtdmJ2YmJibm14Y25haWwuY29tIiwiZXhwIjoxNTg3MDAyOTQ0fQ.LVug_MIsMthyLgRp5VzPI606bGGMOWuY6g3Wbi79S5U	actddive	9411212fdf1212	\N	\N	\N
AShish	kumar	a@ab.aa	rtfghjk	166	$2a$08$bOLJhUs59z/6YRI7MuEPq.N1Pm6kmsQU1DRqnV1kJSVlkzSwXuqAG	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYWIuYWEiLCJleHAiOjE1ODcwNzg2MDZ9.3nRVUyyOE6bbrbBkazdmWuwmaSSipU6fS9fmgmmZm_0	Active	1234567890	\N	\N	\N
aakash123456	sdfgh	bnm@ayush.com	vbn	169	$2a$08$.LhaCSpeamAif75sGZrSfeeNYoJReNEN9jW/KFZsQDgRDRsPddDqC	\N	Inactive	234567890	\N	\N	\N
vinay	tyu	a@ayush.com	sdfghjk	168	$2a$08$vQAW3vcZ6nE.cSZm7LDDGeob7CrTRCBqQDU2ekbcsDAosB5Iyie0C	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYXl1c2guY29tIiwiZXhwIjoxNTg3MDc0MjQ0fQ.OaEib0e9TgjrutcisGJq5U0syem2ccFJJdRcaZNOPDY	Active	1234567890	\N	\N	\N
owa1212123fghj456ish	gheqwesdfgrtyem	owaish@kalim1	1342qwertbnbnmyuil	160	fgh		sdfgh	123494112121234fdf1212	\N	\N	\N
sdvf	ghjk	owaish@sdkalim	jv	163	$2a$08$A8ACFvKdrJBStZ9SzM5q1OG5cxN18hZStb98iKcz1y8QNrkfsw/Ku	\N	sdfbgnhm	94112121212	\N	\N	\N
z	ghjk	owaisqwrhcvbn@sdkalim	s	173	$2a$08$vTbwmnIIswLiShzX7452j.A5jU.ntV1bMezQ0FL3iZyeIERtNypqG	\N	Active	94112121212	\N	\N	\N
sdvf	ghjk	owaisqwhcvbn@sdkalim	jv	165	$2a$08$K/rpaPrRRAE4srC68hUzDusxgnb45nKJle30BvIldlDDjijlJEfCK	\N	sdfbgnhm	94112121212	\N	\N	\N
owaish	kalim	np@kliakashmahalodum	s	188	$2a$08$GLGbSWyM0/9vKzhynEessee2We0Fenh8Gi6It.x2lLvY66FrM5RCG	\N	sdfbhm	94121212	Admin	\N	\N
asdfghj	ssdfghj	hfgdjj@ggvjg	fgjgh	197	$2a$08$QVJEgq8zhxPYrHG6bYOB9.P.i5AsWvY1QaT/JRz23IhInBYZK72kS	\N	Active	gcfgvhj	Accountant	\N	\N
z	ghjk	n@sdkalim	s	174	$2a$08$7us5YCW2QC3W4WdPWrZ2nu3AeNUe8fMsI8q.F6WFgI2feAJChDwFq	\N	sdfbgnhm	94112121212	\N	\N	\N
z	ghjk	np@sdkalim	s	176	$2a$08$r77le6tMFzcNUOiQALLg5uC5.yCbxTHur/SWyGueK6Ii7l/TG8hIq	\N	sdfbgnhm	94112121212	\N	\N	\N
vkgoyal	goyal	vgoyal@j	asbhscvj	171	$2a$08$GijssIzO5YKEf7gEyLlliuQF2qwWtHY0qcBK404n30wX3gYD3mDG.	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZnb3lhbEBqIiwiZXhwIjoxNTg3MDkwMzE4fQ.__ADOMKyHofZt4TL1ng4aSFPjBFWcChp6OlN3bizzuE	Deleted	456789255	\N	\N	\N
vinay	kumar	a@erty	fjf	172	$2a$08$2AWPjq/d9S5bCH5NfVVdEe7bQUBdjXV6aCJwrUzGwst7sV5dspDm6	\N	Inactive	12345678	\N	\N	\N
vinay	ashish	v@goyaljgj	sdfghj	189	$2a$08$9mxXssayLfjFGd4S3aXyYO.DT3zfCiRUA6eVmtUeyd2Ao3KmLtrci	\N	Inactive	7894561235	Customer Care	\N	\N
z	ghjk	np@sdcfddddfckalim	s	177	$2a$08$amHkKxFTmTCMIKiD.2AMpuNjuYaWdO9EILDnNoboAlZPTIUcRPEWG	\N	sdfbgnhm	94112121212	accountant	\N	\N
owaish	kalim	np@sfddfckalim	s	178	$2a$08$Ho1VZrpDxP5ajYotr/3VQOsj4KBr8AKxeqrt/LozxUl.D/CaTzbMu	\N	sdfbhm	94121212	admin	\N	\N
aghvgvjg	gvghvgh	hfc@hvg	ffc	190	$2a$08$K5NbJMf73uABWasUQ0tb5efYA4ae.N/V7nEBi/.W4Gc65b9xMqtMy	\N	Inactive	hgghv	Customer Care	\N	\N
owaish	kalim	np@sfddfjhjhjckalim	s	180	$2a$08$dH.Ezopc0x0Q6tM0O6koK.Fuarvo70ZqUF/birDo2EzvRcr3DUXeO	\N	sdfbhm	94121212	admin	\N	\N
owaish	kalim	np@sfddfjhjhjcddkalim	s	181	$2a$08$8szFCR9vCFIc3Aqq6HRKnOWY5/88snYHjdhza43YZL53NTiHypjB.	\N	sdfbhm	94121212	admin	\N	\N
owaish	kalim	np@kllalim	s	182	$2a$08$Zz6SVw5PctywgPCZ6.rmSuk8QNjaZOgTlyWQD/dQUbcFDFPgjfkLO	\N	sdfbhm	94121212	admin	\N	\N
owaish	kalim	np@klim	s	185	$2a$08$.FWDqvQJEwMEoTEpggmyBOd3Q8j88bWCd3Q1bFh2kFRsMmfImhYwq	\N	sdfbhm	94121212	Admin	\N	\N
aajka	gftff	a@hgfghh	sdfgyu	193	$2a$08$z9xz53HTkj0ZZAgqWb4IrObh5KabyhMbV5x7a1pqWx8I0tyBk4h7C	\N	Active	2345678	Admin	\N	\N
12345678	esdyuj	s@fh	gvgvg	194	$2a$08$QVhIY3hoIMixi48elmUsGum596rYH0sb25GBSJdwYEFXTEk3aJ93.	\N	Active	697653	Admin	\N	\N
svbn	fgbhnjm	bhgb@jhhj	rfvev	195	$2a$08$UlHZOOrn54efiw.BG7EZxerxiSrLWzKvAiXlbM5iY2KrpsKhHIkfG	\N	Inactive	gbhn	Accountant	\N	\N
owaish	kkalimllll	a@owaish	ghvgh	198	$2a$08$k0Xm/Ej19k.B.CewVv98CeVzZ9kHdzeTowSswesSUnKc73n.wuDn2	\N	Active	88945562	Accountant	\N	\N
ABD	KHAN	np@mqqqahalodum	s	199	$2a$08$x2oMufuPOjCFRtstSn6rFuab7LvDAB8CmxsozpA9NoVpyXDO64joy	\N	sdfbhm	94121212	Admin	\N	\N
ABD	KHAN	np@mqqqahadum	hhs	200	$2a$08$eMW.4SyZJYnixdkCm1Ct..oK8Rpin8lgqCaVLiDVUPVYgq1MMbt1G	\N	sdfbhm	94121212	Admin	\N	\N
ABD	KHAN	np@mqqqahaoosodum	hhs	202	$2a$08$JF1o1L/uuc8tYrORdnFoN.4GQ1BHfndNjXy1mKEBzNSqKpBW8VHCu	\N	sdfbhm	94121212	Admin	\N	\N
ayush	sdfghj	efewf@jhbjh	sdgsgvfd	203	$2a$08$TzaalUBBDWVO8WSoABQOZeFiKyO/.5UzwICQncaIJ8aWn8sQzjQgC	\N	Inactive	sdfghefewf	Accountant	\N	\N
John	Wick	np@johnaoosodum	hhs	205	$2a$08$9pt7v29QzIBXBTq4KrhPDuPFe.nL2USvLLSISstUCmslZp5um0u16	\N	sdfbhm	94121212	Accountant	\N	\N
abhishek	kumar	dfds@nvv	greg	204	$2a$08$hVG9jvTGG7kO.i88tllOEOgOesqJHYeNMf4Tk478I5X1S7cJAvNby	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRmZHNAbnZ2IiwiZXhwIjoxNTg3NTc3MDYwfQ.pTr8Kv8i7QfhWvfBqN8rWpRO17bVOS3YeCt4iQ69h_g	Inactive	sdfghj	Accountant	\N	\N
John	Wick	np@johnaohnosodum	hhs	206	$2a$08$Z/oiiVa4HAeB6umKQxvJD.I1SG6R1TUI94ccqgJJsMOd6gV0Kcst.	\N	sdfbhm	94121212	Admin	\N	\N
owaish	gheqwesdfgrtyem	owaish@kalpppim1	1341232qwertbnbnmyuil	164	owashpp	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzcXdoQHNka2FsaW0iLCJleHAiOjE1ODc0ODgyOTh9.jCwbjRlBTCVg7hRg8cFspDjmVe0S91lwYIH0HcG1HUg	sdfgh	123494112121234fdf1212	\N	\N	\N
sdfghjk	ghj	dgh@jhg	fe	207	$2a$08$TvfY3lR.fMLFFfMKg7gp0eyFKAawNXDA.ludCA1n0Vwe5QaR3sWWq	\N	Inactive	sdfgh	Customer Care	\N	\N
dfghj	ddfgh	hjghj@jgj	gv	208	$2a$08$iYBnmWWxwGZH9D0PxilNIOudBJbltWCjftYm27aXjF9l8r3LAmGyC	\N	Inactive	vvv	Accountant	\N	\N
John	Wick	np@lllqqum	hhs	220	$2a$08$fOUHpXouvmPW3J31eMwsSushac/A7pnv0VmQ6xcgfYkJZ5hx8THQa	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@lllqaaqum	hhs	221	$2a$08$TT.ivWr2cJnDroB3hRKTzOQ6iLFboklAgzszPbzu6zVkMBVoM4.PK	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@lllaaqum	hhs	222	$2a$08$CR4664b82vdA8bHYVHT0aOs/RHOLPZfl5V4FwEy5eI/43CMB6yAMC	\N	sdfbhm	94121212	Adminpp	\N	\N
John	Wick	np@lllqum	hhs	223	$2a$08$NiPqBJgg1DKH7bUFo.19/eL53aloXAm.pHrtkBp9Dtx5XU2LAcz/O	\N	sdfbhm	94121212	Adminpp	\N	\N
John	Wick	np@lllzzqum	hhs	224	$2a$08$n6pDTV7jC8yH5X4Bhj7P7Ov00MpoSHz5BRhdFxzogVpopUTuhbsfW	\N	sdfbhm	94121212	Adminpp	\N	\N
John	Wick	np@llsodum	hhs	209	$2a$08$UckSvwtWPOUvwBljYPvSaOlk8u4pwuT6/0Bg.SwMc7k7DQv7353.G	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@llsodlllum	hhs	211	$2a$08$LYDXkGLR6P57oLzhazOOfOxXgkR7Risahab1qNbqfHV.toOCueGMy	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@llsobbvdlllum	hhs	213	$2a$08$bNJppTIseeI1hN39cr9aqe6iI.S10.My//4INmcqloS7tuMKNrjdW	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@llsobwertyudlllum	hhs	214	$2a$08$UUpiFc5UNONg7hD477spuOetLBRqzKrqSw7fgYZJ.JhfFUIRjotdK	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@llsobweqazyudlllum	hhs	215	$2a$08$WW8gAIaz2UzADBJuZTF.5.rE.bogqK/iUYDliJRbMCskYSGNkUQxG	\N	sdfbhm	94121212	Admin	\N	\N
John	Wick	np@llsobwedlllum	hhs	216	$2a$08$.a.CkZ9YRhtNfaIEENoa7uaDooQw230Cf/H6QO9jY33J4igHfhE92	\N	sdfbhm	94121212	Admin	\N	\N
abhi	Hirnotia	ashutoshhhirnotia23@gmail.co	Gali no 16 B,Lajjapuri ,chamri road ,Hapur	225	$2a$08$KNUx7RV/N.26TuUIbpyTM.VOEQgLtsV0SdTJU0vBdwE2bONrzKQKu	\N	Inactive	09084768046	Accountant	\N	\N
John	Wick	np@lllzxxzqum	hhs	226	$2a$08$yDtLA0X9VKjWwMcWyBKV8.vqbPEivkUwqO1k26vFmTEtdHApuywAK	\N	sdfbhm	94121212	Admin	\N	\N
aakash	sa	ass@asdfghjk	werg	227	$2a$08$cre.cfqP6WA7EiIzowj2yeTbekkA7dkvnnqPE69QksV1JC.0//q9i	\N	Inactive	sas	Admin	\N	\N
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa	kumar	ahirnotia@ch.iitr.ac.in	dfghkj	132	$2a$08$QsKrZ4Vsgsj.fRw4eLo4jOZ0D4PDQF7BgC9sCoeZ/OnHxHriyjDka	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFoaXJub3RpYUBjaC5paXRyLmFjLmluIiwiZXhwIjoxNTg3NjE4NTU0fQ.29ThHn7P9bKgklwaduvHVAG_KcaFrRrJAig67oZ6a74	Active	1234567892	Admin	\N	\N
John	Wick	np@lllum	hhs	218	$2a$08$wVDVSrhmssn4aS99Gl53TOUb4VeX/m/J9JFTu5AN6w3c7ETyeTF7K	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5wQGxsbHVtIiwiZXhwIjoxNTg3NTg1NDE3fQ._geB3KGG8qgzpwiRf_KYHqKJj7eB4Hdlt61hNIXxQC4	sdfbhm	94121212	Admin	\N	\N
ABD	KHAN	abd@4321	RB	231	$2a$08$n1CCExpnfC7yBJiQ2g75su4iTyUA.NQ6LzRg848OQ4rn387NU1GgS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiZEA0MzIxIiwiZXhwIjoxNTg3NjEyNzQ0fQ.nh9KiNxMnC64sGHoRa3fFfoRgzl_8DLAF1buymURMiU	inactive	1234	Admin	\N	\N
aakash	hirnotia	a@f	etrdtfyguh	170	$2a$08$p.p22Arzrb05lkHieShEUO6NqD16lWny50/IC8WzwSJTPogZkTCXm	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAZiIsImV4cCI6MTU4NzYxODE1NH0.UIVLYG03-_2ugsEl4ZtvwjcHqbwmc9RUXJJ8YvqCzjo	Active	9258656829	\N	\N	\N
ABD-OW	KHANnnnnnnnnnnn	kalim@4321	RKB	230	$2a$08$QM6T2WXL9LsGzw1zM.vjW.VjY3IjJo30xYcClFF4vZNxKP4LLa0jW	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltQDQzMjEiLCJleHAiOjE1ODc1OTA4NTV9.cofGchjkEox3CqkAkOcVcvsbYxISAECo5ZfMfN3D8dU	sdfgh	4321	Admin	\N	\N
Ashish123	Hirnotia	a@aakash	sdfgh	229	$2a$08$OCG1Aj4hzf1Q.fYopQC85u0SSdeK4d8zpMxCSUZRa0nhe25Vy//QS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYWFrYXNoIiwiZXhwIjoxNTg3NTg3MDk5fQ.DjSuu5q3MxMT_1Ak1oM-kzmzZd4pIF4DhCtxgrMdOEY	Inactive	1234567	Customer Service	\N	\N
Akash	Hirnotia	kalim@321	RKB	228	owas	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltQDMyMSIsImV4cCI6MTU4NzU4NjQwNX0.bPeb0_1sCyIxMDyKI3nKg89WRAH-2idfNmXD6UqomSA	sdfgh	4321	Customer Service	\N	\N
ABD	KHAN	kalimowaish@gmail.com	RB	237	$2a$08$7ptZr.PM1JqwbdGDfdFy5OTt5ZEWCU.DAkj4LFclakKmt3jIrVT0S	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltb3dhaXNoQGdtYWlsLmNvbSIsImV4cCI6MTU4NzY5MjUzNn0.YUn0BecJf4irP4WENzngqcMJVX_UcM3nm3I38A0Z5N4	inactive	1234	Admin	2020-04-13 21:11:38.752651831+05:30	24-April
ashutosh	hirnotia	s@sdfghjgvhjvhvbvn	ghfh	232	$2a$08$qvDTY7NpIptml0m2DBpCie8eh.zv/KT6HkkdcIAuaa6/beeIEpyc6	\N	Inactive	5464	Accountant	2020-04-12 21:45:48.808732474+05:30	2020-04-12 21:45:48.808732327+05:30
aakash	asdfghjkl	hj@hjj	ff	233	$2a$08$X819PuQ.YnL2.ExR5oy6ke1Ffw.OZWU2X.O5htVqUsP/dqtKCx6ui	\N	Deleted	00000000000000000000000000000000000000	Accountant	2020-04-12 21:48:38.879969665+05:30	2020-04-12 21:48:38.879969137+05:30
ABD	KHAN	abd@4233321	RB	235	$2a$08$WHT6x7sxkNPZeJp4YCYMceWZeZTokMh1EqHuGiqnvmx1czjIvM09q	\N	inactive	1234	Admin	2020-04-12 22:06:06.449717438+05:30	24-April
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
-- Name: slh_communication_templates_Comm.Template_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_communication_templates_Comm.Template_Id_seq"', 1, true);


--
-- Name: slh_customers_Customer_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_Customer_Id_seq"', 50, true);


--
-- Name: slh_customers_pending_orders_Customer_Primary_Order_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq"', 22, true);


--
-- Name: slh_partners_Partner_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_partners_Partner_Id_seq"', 9, true);


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

SELECT pg_catalog.setval('public."slh_teams_TeamId_seq"', 237, true);


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
-- Name: slh_communication_templates slh_communication_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_communication_templates
    ADD CONSTRAINT slh_communication_templates_pkey PRIMARY KEY ("Comm.Template_Id");


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

