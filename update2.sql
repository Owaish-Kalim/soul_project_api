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
    "Slot_Time" text,
    "Slot_Date" text,
    "Customer_Souls_Id" text,
    "Customer_Name" text,
    "Customer_Gender" text,
    "Customer_Address" text,
    "Customer_Pincode" text,
    "Merchant_Transaction_Id" text,
    "Partner_Souls_Id" text,
    "Partner_Name" text,
    "Partner_Mobile_No" text,
    "Commision_Type" text,
    "Commision_Amount" text,
    "CreatedAt" timestamp without time zone,
    "Created_By" text,
    "Updated_By" text,
    "Status" text,
    "Id" integer NOT NULL
);


ALTER TABLE public.slh_assign_customer_with_partner OWNER TO postgres;

--
-- Name: slh_assign_customer_with_partner_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_assign_customer_with_partner_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_assign_customer_with_partner_Id_seq" OWNER TO postgres;

--
-- Name: slh_assign_customer_with_partner_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_assign_customer_with_partner_Id_seq" OWNED BY public.slh_assign_customer_with_partner."Id";


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
    "Trigger_Time" text NOT NULL,
    "Status" text
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
    "Customer_Gender" character varying(50) NOT NULL,
    "Customer_Name" character varying(50) NOT NULL,
    "Customer_Address" character varying(200) NOT NULL,
    "Registered_Source" character varying(100) NOT NULL,
    "Customer_Souls_Id" text,
    "Status" text NOT NULL,
    "Customer_Mobile_No" text NOT NULL,
    "Pincode" text NOT NULL,
    "CreatedAt" text NOT NULL,
    "Last_Access_Time" text NOT NULL
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
    "Therapist_Gender" character varying(50) NOT NULL,
    "Massage_For" character varying(50) NOT NULL,
    "Customer_Address" character varying(200) NOT NULL,
    "Massage_Duration" text NOT NULL,
    "Customer_Souls_Id" text,
    "Customer_Name" text,
    "Merchant_Transaction_Id" text,
    "Number_Of_Therapists_Required" text,
    "Pincode" text,
    "Is_Order_Confirmed" text,
    "Total_Order_Amount" text,
    "Latitude" text,
    "Longitude" text,
    "Slot_Date" text,
    "Slot_Time" text,
    "CreatedAt" text
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
    "Last_Updated_By" text,
    "Commission_Type" text,
    "CreatedBy" text,
    "Partner_Mobile_No" text,
    "Pincode" text,
    "Per_Visit_Price_Commission" text,
    "Latitude" text,
    "Longitude" text,
    "Partner_Souls_Id" text,
    "Onboard_Date" text,
    "CreatedAt" text,
    "UpdatedAt" text
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
    "Souls_Setting_Id" integer NOT NULL,
    "Type" text NOT NULL,
    "URL" text NOT NULL,
    "Description" text NOT NULL,
    "HostName" text NOT NULL,
    "UserName" text NOT NULL,
    "Password" text NOT NULL
);


ALTER TABLE public.slh_souls_settings OWNER TO postgres;

--
-- Name: slh_souls_settings_Souls_Setting_Id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."slh_souls_settings_Souls_Setting_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."slh_souls_settings_Souls_Setting_Id_seq" OWNER TO postgres;

--
-- Name: slh_souls_settings_Souls_Setting_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."slh_souls_settings_Souls_Setting_Id_seq" OWNED BY public.slh_souls_settings."Souls_Setting_Id";


--
-- Name: slh_team_has_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.slh_team_has_role (
    "Team_Has_Role_Id" integer NOT NULL,
    "Team_Id" integer NOT NULL,
    "Status" text NOT NULL,
    "FirstName" text,
    "LastName" text,
    "CreatedAt" text,
    "UpdatedAt" text
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
    "JoiningDate" text,
    "Gender" text,
    "Member_Image" text
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
    "Customer_Address" character varying,
    "Transaction_Mode" character varying,
    "Bank_Type" character varying,
    "Merchant_Transaction_Id" text,
    "Massage_Duration" text,
    "Customer_Name" text,
    "Customer_Souls_Id" text,
    "Payment_Gateway_Id" text,
    "Latitude" text,
    "Longitude" text,
    "Payment_Gateway_Mode" text,
    "Total_Order_Amount" text,
    "Number_Of_Therapist_Required" text,
    "Pincode" text,
    "Slot_Time" text,
    "Slot_Date" text,
    "CreatedAt" text
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
-- Name: slh_assign_customer_with_partner Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_assign_customer_with_partner ALTER COLUMN "Id" SET DEFAULT nextval('public."slh_assign_customer_with_partner_Id_seq"'::regclass);


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
-- Name: slh_souls_settings Souls_Setting_Id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_souls_settings ALTER COLUMN "Souls_Setting_Id" SET DEFAULT nextval('public."slh_souls_settings_Souls_Setting_Id_seq"'::regclass);


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

COPY public.slh_assign_customer_with_partner ("Slot_Time", "Slot_Date", "Customer_Souls_Id", "Customer_Name", "Customer_Gender", "Customer_Address", "Customer_Pincode", "Merchant_Transaction_Id", "Partner_Souls_Id", "Partner_Name", "Partner_Mobile_No", "Commision_Type", "Commision_Amount", "CreatedAt", "Created_By", "Updated_By", "Status", "Id") FROM stdin;
\.


--
-- Data for Name: slh_communication_templates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_communication_templates ("Comm.Template_Id", "Comm.Template_Type", "Trigger_For", "SMS_Content", "Email_Content", "Subject", "Trigger_Time", "Status") FROM stdin;
1	abcd	customers	popopo	opopopop	Owaish	ten-min	\N
2	email	customer	adfa	asfag	maas	10	\N
3	erty	customer	dfgnm	poiuyt	dfgh	dfghj	\N
5	VINAY	partner	sdadasdsa	<p>dasdfasfax</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p>	sdasdasd	10	Active
6	VINAY	partner	Check sms	<p>This is the only content</p>	Check	30	Inactive
7	owaish	customer	wwwr	<p>fghjkl</p>	dww	10	Active
4	erty	customer	dfgnm	poiuyt	owaish	dfghj	Active
8	vinay kumar	customer	jhjh	<p>here</p>	hghgk	30	Active
\.


--
-- Data for Name: slh_customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers ("Customer_Id", "Customer_Email", "Customer_Gender", "Customer_Name", "Customer_Address", "Registered_Source", "Customer_Souls_Id", "Status", "Customer_Mobile_No", "Pincode", "CreatedAt", "Last_Access_Time") FROM stdin;
64	vinay@L	FeMale	Vinay	DL	website	2020041664	Active	9632581670	13596	16-04-2020 6:57:27 PM	16-04-2020 6:57:37 PM
58	ahirnotia@ch.iitr.ac.in	female	ayushi	Ny		2020041558	active	1234522420	2345678	15-04-2020 03:50:04 AM	16-04-2020 10:38:22 PM
67	vinay@L	FeMale	Vinay	DL	website	2020041667	Active	9632341670	13596	16-04-2020 10:27:25 PM	16-04-2020 10:27:25 PM
\.


--
-- Data for Name: slh_customers_pending_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_customers_pending_orders ("Customer_Order_Id", "Customer_Id", "Therapist_Gender", "Massage_For", "Customer_Address", "Massage_Duration", "Customer_Souls_Id", "Customer_Name", "Merchant_Transaction_Id", "Number_Of_Therapists_Required", "Pincode", "Is_Order_Confirmed", "Total_Order_Amount", "Latitude", "Longitude", "Slot_Date", "Slot_Time", "CreatedAt") FROM stdin;
9	36	zxcvb	zxcvbn	wqewrtydfghjk	asdfghj	2020040736	owasih	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
10	36	zxcvb	zxcvbn	wqewrtydfghjk	asdfghj	2020040736	owasih	158627539336	\N	\N	\N	\N	\N	\N	\N	\N	\N
11	36	zxcvb	zxcvbn	wqewrtydfghjk	asdfghj	2020040736	owasih	158627546436	\N	\N	\N	\N	\N	\N	\N	\N	\N
12	36	zxcvb	zxcvbn	wqewrtydfghjk	asdfghj	2020040736	owasih	158627547636	\N	\N	\N	\N	\N	\N	\N	\N	\N
13	36	zxcvb	zxcvbn	wqewrtydfghjk	asdfghj	2020040736	owasih	158627864736	\N	\N	\N	\N	\N	\N	\N	\N	\N
14	38	male	cbn	ertyj	trew	2020040738	kalim	158628051538	\N	\N	\N	\N	\N	\N	\N	\N	\N
15	38	male	cbn	ertyj	trew	2020040738	kalim	158628052138	\N	\N	\N	\N	\N	\N	\N	\N	\N
16	38	male	cbn	ertyj	trew	2020040738	kalim	158628052438	\N	\N	\N	\N	\N	\N	\N	\N	\N
17	38	male	cbn	ertyj	trew	2020040738	kalim	158628052838	\N	\N	\N	\N	\N	\N	\N	\N	\N
18	38	male	cbn	ertyj	trew	2020040738	kalim	158628053438	\N	\N	\N	\N	\N	\N	\N	\N	\N
19	39	male	cbn	ertyj	trew	2020040839	kalim	158628813539	\N	\N	\N	\N	\N	\N	\N	\N	\N
20	40	male	cbn	ertyj	trew	2020040840	kalim	158628820740	\N	\N	\N	\N	\N	\N	\N	\N	\N
21	40	male	cbn	ertyj	trew	2020040840	kalim	158634174540	\N	\N	\N	\N	\N	\N	\N	\N	\N
22	40	male	cbn	ertyj	trew	2020040840	kalim	1586511325-40	\N	\N	\N	\N	\N	\N	\N	\N	\N
23	40	male	cbn	ertyj	trew	2020040840	kalim	1586801692-40	\N	\N	\N	\N	\N	\N	\N	\N	\N
24	40	male	cbn	ertyj	trew	2020040840	kalim	1586802136-40	\N	\N	\N	\N	\N	\N	\N	\N	\N
25	40	male	cbn	ertyj	trew	2020040840	kalim	1586802280-40	\N	\N	\N	\N	\N	\N	\N	\N	\N
26	40	male	cbn	ertyj	trew	2020040840	kalim	1586802657-40	\N	\N	\N	\N	\N	\N	\N	\N	\N
27	60	male	cbn	DL	trew	2020041560	Vinay	1586900095-60	2	13596	Confirmmed	10000	qwerty	asdfg	\N	\N	\N
28	58	male	cbn	Ny	trew	2020041558	Kashyap	1586900113-58	2	65449	Confirmmed	10000	qwerty	asdfg	\N	\N	\N
29	58	male	cbn	Ny	trew	2020041558	Kashyap	1586933802-58	2	65449	Confirmmed	10000	qwerty	asdfg	\N	\N	\N
30	60	male	cbn	DL	trew	2020041560	Vinay	1586934034-60	2	13596	Confirmmed	10000	qwerty	asdfg	\N	\N	\N
31	64	male	cbn	DL	trew	2020041664	Vinay	1587057071-64	2	13596	Confirmed	10000	5.239	1.256	05-05-2020	5:5	16-04-2020 10:41:11 PM
\.


--
-- Data for Name: slh_partners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_partners ("Partner_Id", "Partner_Name", "Partner_Gender", "Partner_Address", "Partner_Email", "Last_Updated_By", "Commission_Type", "CreatedBy", "Partner_Mobile_No", "Pincode", "Per_Visit_Price_Commission", "Latitude", "Longitude", "Partner_Souls_Id", "Onboard_Date", "CreatedAt", "UpdatedAt") FROM stdin;
20	Jack Reacher	male	poiu	jack@123	ll	oiioi	uu	9087654	234569	3456922	1.236	2.36	2020041620	20-02-2020	16-04-2020 9:51:29 PM	16-04-2020 9:51:29 PM
21	Jack Sparrow	female	poiu	jack@1234	lal	oiiocci	cuu	908763254	23456339	345693322	1.2336	2.436	2020041621	23-02-2020	16-04-2020 9:52:15 PM	16-04-2020 9:52:15 PM
22	Jason Bourne	female	poiu	Jason@1234	lal	oiiocci	cuu	90863254	23456339	345693322	1.2336	2.436	2020041622	23-02-2020	16-04-2020 9:55:55 PM	16-04-2020 9:55:55 PM
23	James Bond	female	poiu	James@1234	lal	oiiocci	cuu	920863254	23456339	345693322	1.2336	2.436	2020041623	23-02-2020	16-04-2020 9:58:16 PM	16-04-2020 9:58:16 PM
24	jacky	Male	A-125, Rajiv Bhawan, Roorkee	vg@ch.iitr.ac.in	nkdbj	Percentage(%)	vhjx	07078264246	247664	23	ed	das	2020041624		16-04-2020 9:58:20 PM	16-04-2020 9:58:20 PM
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

COPY public.slh_souls_settings ("Souls_Setting_Id", "Type", "URL", "Description", "HostName", "UserName", "Password") FROM stdin;
1	SMS	qwertyui	1234567	hughjg	owaish	jbbj
2	EMAIL	wertyui	qwer	wqedsc	kumar	1234567
\.


--
-- Data for Name: slh_team_has_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_team_has_role ("Team_Has_Role_Id", "Team_Id", "Status", "FirstName", "LastName", "CreatedAt", "UpdatedAt") FROM stdin;
1	185	sdfbhm	\N	\N	\N	\N
1	188	sdfbhm	\N	\N	\N	\N
1	193	Active	\N	\N	\N	\N
1	194	Active	\N	\N	\N	\N
2	195	Inactive	\N	\N	\N	\N
1	202	sdfbhm	\N	\N	\N	\N
2	203	Deleted	\N	\N	\N	\N
2	204	Inactive	\N	\N	\N	\N
1	206	sdfbhm	\N	\N	\N	\N
2	208	Inactive	\N	\N	\N	\N
2	205	sdfbhm	\N	\N	\N	\N
1	213	sdfbhm	John	Wick	\N	\N
1	214	sdfbhm	John	Wick	\N	\N
1	215	sdfbhm	John	Wick	\N	\N
1	216	sdfbhm	John	Wick	\N	\N
1	218	sdfbhm	John	Wick	\N	\N
1	220	sdfbhm	John	Wick	\N	\N
1	221	sdfbhm	John	Wick	\N	\N
2	225	Inactive	Dhanprakash	Hirnotia	\N	\N
1	226	sdfbhm	John	Wick	\N	\N
1	230	inactive	Ow	Kal	\N	\N
3	229	Active	Ashish	Hirnotia	\N	\N
3	228	inactive	Ow	Kal	\N	\N
1	227	Inactive	aakash	sa	\N	\N
1	231	inactive	ABD	KHAN	\N	\N
2	232	Inactive	ashutosh	hirnotia	\N	\N
2	233	Deleted	aakash	asdfghjkl	\N	\N
1	235	inactive	ABD	KHAN	\N	\N
1	237	inactive	ABD	KHAN	\N	\N
1	238	sdfbhm	John	Wick	\N	\N
1	239	sdfbhm	John	Wick	\N	\N
1	240	sdfbhm	John	Wick	\N	\N
1	241	sdfbhm	Owaish	Kalim	15-04-2020 8:41:41 PM	15-04-2020 8:41:41 PM
2	242	sdfbhm	jason	bourne	15-04-2020 8:44:46 PM	15-04-2020 8:44:46 PM
2	243	sdfbhm	Jack	Reacher	15-04-2020 11:2:58 PM	15-04-2020 11:2:58 PM
2	244	sdfbhm	Jack	Reacher	15-04-2020 11:5:22 PM	15-04-2020 11:5:22 PM
2	246	sdfbhm	Jack	Reacher	15-04-2020 11:5:45 PM	15-04-2020 11:5:45 PM
2	248	sdfbhm	Jack	Reacher	15-04-2020 11:6:2 PM	15-04-2020 11:6:2 PM
2	249	sdfbhm	Jack	Reacher	15-04-2020 11:6:44 PM	15-04-2020 11:6:44 PM
2	250	sdfbhm	Jack	Reacher	15-04-2020 11:12:11 PM	15-04-2020 11:12:11 PM
2	251	sdfbhm	Jack	Reacher	15-04-2020 11:13:25 PM	15-04-2020 11:13:25 PM
2	252	sdfbhm	Jack	Reacher	15-04-2020 11:13:43 PM	15-04-2020 11:13:43 PM
2	253	sdfbhm	Jack	Reacher	16-04-2020 2:32:26 AM	16-04-2020 2:32:26 AM
2	256	sdfbhm	Jack	Reacher	16-04-2020 3:23:33 AM	16-04-2020 3:23:33 AM
2	258	sdfbhm	Jack	Reacher	16-04-2020 3:24:6 AM	16-04-2020 3:24:6 AM
2	260	sdfbhm	Jack	Reacher	16-04-2020 3:33:40 AM	16-04-2020 3:33:40 AM
2	261	sdfbhm	Jack	Reacher	16-04-2020 3:13:9 PM	16-04-2020 3:13:9 PM
1	263	sdfbhm	Jack	Reacher	16-04-2020 9:50:1 PM	16-04-2020 9:50:1 PM
\.


--
-- Data for Name: slh_teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_teams ("FirstName", "LastName", "Email", "Address", "TeamId", "Password", "Token", "Status", "MobileNo", "Role", "CreatedAt", "JoiningDate", "Gender", "Member_Image") FROM stdin;
AShish	kumar	ak@asf.as	Qwerty	141	$2a$08$tsYFF7bP/noieJBR14.rWeXMR5Gg9bUvkYQFHWz7xXLdQSW4kgE8m	\N	Inactive	234567898	\N	\N	\N	\N	\N
AShish	kumar	sds@af.af	asdgfhgjk	116	$2a$08$EfOW8CYMoHX3rQIgcWEsROUVCndwe2yHElpTzJ07ZAD1OkVcolCRW	\N	inactive	1234567789	admin	\N	\N	\N	\N
	cvbn	sh@gghjsdertfvghmail.com	v	125	$2a$08$grUDESzu7oc4kPvGn3zFmuPh6RLoIVq1Orwu19RtKCEP5K8VIu54C	\N		94112121212		\N	\N	\N	\N
saSdfa	test	sh@gghjmail.com	dfghjk	117	$2a$08$KsiRMhJslypxxptG3H9AM.9dvEewCFZE7SrS/Wk3GuJkpFpKrPiti	\N	active	94112121212	admin	\N	\N	\N	\N
bn	test	sh@gghjvghmail.com	v	119	$2a$08$1kG4naPdgil9YQW3WhrdfuxS64HJz.wtEqtG4Ow6zJ7YepfIYiE06	\N	fgh	94112121212	admin	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjsdertfvail.com	v	127	$2a$08$XVV/PmNuQNFyKLK0bFdu9uw2nL1E6EqLphlzRRrnjitT.vjp1uYmO	\N	sdfbgnhm	94112121212	zxcv 	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjsdertfvartyil.com	v	128	$2a$08$PHaTWixqJLnXxufOGSsEIeYPc2iGSe7IGKqdc7vuonfi99YVIK1ZW	\N	sdfbgnhm	94112121212	zxcv 	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjsdefrgthsdertfvartyil.com	v	130	$2a$08$CqJKFGCslUEL4v0Vt.z.rOLtIbk7mOHOVtiU8uF2GWSh3SN11RJBS	\N	sdfbgnhm	94112121212	zxcv	\N	\N	\N	\N
asfa	test	owaish@gmail.com	dfghjk	113	$2a$08$xQ7ctow8rGKAUrcRWw.LBetRFGK.GwHOU2M7GM.jcKfpchDDNhXx2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzaEBnbWFpbC5jb20iLCJleHAiOjE1ODY5MTI5MDV9.zyfDAqPo4AdYNo-V2kWB6gmEWXLR7t3--H9atmOczkM	accountant	94112121212	admin	\N	\N	\N	\N
bn	cvbn	sh@gghjsdfvghmail.com	v	123	$2a$08$XIZ./xeRmlr55ZDqHbfXF.fPCuJRLcHW3uCjJycv0FPYB0PB6Hk8.	\N	fgh	94112121212	admin	\N	\N	\N	\N
AShish	kumar	a@a.cc	asdfghj	131	$2a$08$bgExze5662U3E8v3pPtlE.AmoJ2WDIQkur7135857Z5M1..8.qdxS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYS5jYyIsImV4cCI6MTU4NjkwMjU0MX0.ruT9B7wYQXksugTi6FCSggdSqVvXk_I5H9bbiI0WKo4	Active	1234567894	Admin	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfvdefrgthsdertfvartyil.com	v	134	$2a$08$sc1O84ampL5zG/A2yLggq.6UNosQlxINgZPS3C4KOYrBi7aduJuPi	\N	sdfbgnhm	94112121212	zxcv	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfasdfgvdefrgthsdertfvartyil.com	v	135	$2a$08$.q5yPJhyOto60AZdWjNrjemK2VxkgxN0NrAJGI8cFbDt1GxqYs9sG	\N	sdfbgnhm	94112121212	zxcv	\N	\N	\N	\N
sdvf	sdfcvbn	sh@gghjssdcfwertyasdfgvdefrgthsdertfvartyil.com	v	137	$2a$08$rv11WiK/voyPw7aqH588NeELy/DW6psUv.V7TYDR.DDaISJMfZ0Zi	\N	sdfbgnhm	94112121212	zxcv	\N	\N	\N	\N
AShish	kumar	ashutoshhhirnotia23@gmail.com	Gali no 16 B,Lajjapuri ,chamri road ,Hapur	140	$2a$08$XDRmXktAKpz2UFT1QzgXjeKvzvtRGW1iAQS/WjNYVZkCgkf3SnuCe	\N	Inactive	09084768046	\N	\N	\N	\N	\N
sdvf	sdfcvbn	as@gmvbnail.com	v	145	$2a$08$JfZvVtQjv4asM1vBYhU3AewUtysWC75fgCsMGasJPDviR7XDAHVBi	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbnail.com	v	148	$2a$08$6GwLCpmKrIjIpT1ok27Aw.5VzV.WcqMwYWpVkox7fjqWk5OfITjD2	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbbnail.com	v	152	$2a$08$u6u0N.HdLZBDRY3IE2zIRuuuYufeiGrqcglXeBHxz9fkX/G5zGWJq	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
sdvf	sdfcvbn	as@gmvbvbbxcnail.com	jv	155	$2a$08$isp0ldIIBQD5MeiZNgkkgelFa88Lch/PVBpGYhUpFa6.6/DXI.3eG	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
sdvf	ghjk	as@gmvbvbbbnmxcnail.com	jv	158	$2a$08$pSK3bO90lI4qQg.VV1bV6.F694BjTelTDyjYCMZnFwcrKGLwUBaju	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
owaish	kalim	np@kliakashlodum	s	186	$2a$08$BeWp.U7AhqE0hQnvJpQ1EODfU8c25Ehw6JLEmsw6Mr0ijbf/mN3ca	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
ayush	bhai	ayush@123456	rtdt	196	$2a$08$xHj/41rSidIlKdiFD8hymOSMbqDKGsuDeEwWbfx14bePeyrTwq/Gm	\N	Inactive	7894561235	Accountant	\N	\N	\N	\N
sdvf	ghjk	owaish@kalim	jv	161	$2a$08$4gIPpUW6mHOPJtXqfyg3.ecAKUidtoaMtOTn0.LCts.g3XuL8Rc2q	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
owaish	gheem	a@a.aa	12qwertbnbnmyuil	157	$2a$08$l.jy1pudS6lWL/Q.4/ctjOj0sGqYz.OoigzgvRanMb1/QiNast0Ye	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtdmJ2YmJibm14Y25haWwuY29tIiwiZXhwIjoxNTg3MDAyOTQ0fQ.LVug_MIsMthyLgRp5VzPI606bGGMOWuY6g3Wbi79S5U	actddive	9411212fdf1212	\N	\N	\N	\N	\N
AShish	kumar	a@ab.aa	rtfghjk	166	$2a$08$bOLJhUs59z/6YRI7MuEPq.N1Pm6kmsQU1DRqnV1kJSVlkzSwXuqAG	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYWIuYWEiLCJleHAiOjE1ODcwNzg2MDZ9.3nRVUyyOE6bbrbBkazdmWuwmaSSipU6fS9fmgmmZm_0	Active	1234567890	\N	\N	\N	\N	\N
aakash123456	sdfgh	bnm@ayush.com	vbn	169	$2a$08$.LhaCSpeamAif75sGZrSfeeNYoJReNEN9jW/KFZsQDgRDRsPddDqC	\N	Inactive	234567890	\N	\N	\N	\N	\N
vinay	tyu	a@ayush.com	sdfghjk	168	$2a$08$vQAW3vcZ6nE.cSZm7LDDGeob7CrTRCBqQDU2ekbcsDAosB5Iyie0C	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYXl1c2guY29tIiwiZXhwIjoxNTg3MDc0MjQ0fQ.OaEib0e9TgjrutcisGJq5U0syem2ccFJJdRcaZNOPDY	Active	1234567890	\N	\N	\N	\N	\N
owa1212123fghj456ish	gheqwesdfgrtyem	owaish@kalim1	1342qwertbnbnmyuil	160	fgh		sdfgh	123494112121234fdf1212	\N	\N	\N	\N	\N
sdvf	ghjk	owaish@sdkalim	jv	163	$2a$08$A8ACFvKdrJBStZ9SzM5q1OG5cxN18hZStb98iKcz1y8QNrkfsw/Ku	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
z	ghjk	owaisqwrhcvbn@sdkalim	s	173	$2a$08$vTbwmnIIswLiShzX7452j.A5jU.ntV1bMezQ0FL3iZyeIERtNypqG	\N	Active	94112121212	\N	\N	\N	\N	\N
sdvf	ghjk	owaisqwhcvbn@sdkalim	jv	165	$2a$08$K/rpaPrRRAE4srC68hUzDusxgnb45nKJle30BvIldlDDjijlJEfCK	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
owaish	kalim	np@kliakashmahalodum	s	188	$2a$08$GLGbSWyM0/9vKzhynEessee2We0Fenh8Gi6It.x2lLvY66FrM5RCG	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
asdfghj	ssdfghj	hfgdjj@ggvjg	fgjgh	197	$2a$08$QVJEgq8zhxPYrHG6bYOB9.P.i5AsWvY1QaT/JRz23IhInBYZK72kS	\N	Active	gcfgvhj	Accountant	\N	\N	\N	\N
vinay	test	sh@gmail.com	fefefe	114	$2a$08$XWC2NdF8/5a4fv9uiBvD2ujxHNY9laSh89ERnbJ.W.cxPOdfjc0JK	\N	active	94112121212	admin	\N	\N	\N	\N
vinay	Aryan	k@kk.kk	qwertyui	142	$2a$08$r1N2fi5YGNgGKUgelFBA7u.os5AkqHe.YaJRG5VCKmC3fG1TF0PmC	\N	Inactive	23456789	\N	\N	\N	\N	\N
Vinay	Goyal	as@gmail.com	A-125, Rajiv Bhawan, Roorkee	139	$2a$08$oNfiL1o0bskZJZQJFuHaZu3Kev9kBGS3E.5VVFpSXIASut7vkT2KK	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzQGdtYWlsLmNvbSIsImV4cCI6MTU4NjkxMzgxOH0.0XI66C4bCqeC7YMUvxe4akls1ECjkwo875kZM6nalNU	Active	07078264246	\N	\N	\N	\N	\N
Ashish kumar	kumar	sd@af.af	asdgfhgjk	115	$2a$08$mHcOXlXlOaAdv5rsZ9O8B.8LyBm/zwwbQXzr11Gt5kICor/SoCRGW	\N	Active	1234567789	admin	\N	\N	\N	\N
z	ghjk	n@sdkalim	s	174	$2a$08$7us5YCW2QC3W4WdPWrZ2nu3AeNUe8fMsI8q.F6WFgI2feAJChDwFq	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
z	ghjk	np@sdkalim	s	176	$2a$08$r77le6tMFzcNUOiQALLg5uC5.yCbxTHur/SWyGueK6Ii7l/TG8hIq	\N	sdfbgnhm	94112121212	\N	\N	\N	\N	\N
vkgoyal	goyal	vgoyal@j	asbhscvj	171	$2a$08$GijssIzO5YKEf7gEyLlliuQF2qwWtHY0qcBK404n30wX3gYD3mDG.	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZnb3lhbEBqIiwiZXhwIjoxNTg3MDkwMzE4fQ.__ADOMKyHofZt4TL1ng4aSFPjBFWcChp6OlN3bizzuE	Deleted	456789255	\N	\N	\N	\N	\N
vinay	kumar	a@erty	fjf	172	$2a$08$2AWPjq/d9S5bCH5NfVVdEe7bQUBdjXV6aCJwrUzGwst7sV5dspDm6	\N	Inactive	12345678	\N	\N	\N	\N	\N
vinay	ashish	v@goyaljgj	sdfghj	189	$2a$08$9mxXssayLfjFGd4S3aXyYO.DT3zfCiRUA6eVmtUeyd2Ao3KmLtrci	\N	Inactive	7894561235	Customer Care	\N	\N	\N	\N
z	ghjk	np@sdcfddddfckalim	s	177	$2a$08$amHkKxFTmTCMIKiD.2AMpuNjuYaWdO9EILDnNoboAlZPTIUcRPEWG	\N	sdfbgnhm	94112121212	accountant	\N	\N	\N	\N
owaish	kalim	np@sfddfckalim	s	178	$2a$08$Ho1VZrpDxP5ajYotr/3VQOsj4KBr8AKxeqrt/LozxUl.D/CaTzbMu	\N	sdfbhm	94121212	admin	\N	\N	\N	\N
aghvgvjg	gvghvgh	hfc@hvg	ffc	190	$2a$08$K5NbJMf73uABWasUQ0tb5efYA4ae.N/V7nEBi/.W4Gc65b9xMqtMy	\N	Inactive	hgghv	Customer Care	\N	\N	\N	\N
owaish	kalim	np@sfddfjhjhjckalim	s	180	$2a$08$dH.Ezopc0x0Q6tM0O6koK.Fuarvo70ZqUF/birDo2EzvRcr3DUXeO	\N	sdfbhm	94121212	admin	\N	\N	\N	\N
owaish	kalim	np@sfddfjhjhjcddkalim	s	181	$2a$08$8szFCR9vCFIc3Aqq6HRKnOWY5/88snYHjdhza43YZL53NTiHypjB.	\N	sdfbhm	94121212	admin	\N	\N	\N	\N
owaish	kalim	np@kllalim	s	182	$2a$08$Zz6SVw5PctywgPCZ6.rmSuk8QNjaZOgTlyWQD/dQUbcFDFPgjfkLO	\N	sdfbhm	94121212	admin	\N	\N	\N	\N
owaish	kalim	np@klim	s	185	$2a$08$.FWDqvQJEwMEoTEpggmyBOd3Q8j88bWCd3Q1bFh2kFRsMmfImhYwq	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
aajka	gftff	a@hgfghh	sdfgyu	193	$2a$08$z9xz53HTkj0ZZAgqWb4IrObh5KabyhMbV5x7a1pqWx8I0tyBk4h7C	\N	Active	2345678	Admin	\N	\N	\N	\N
12345678	esdyuj	s@fh	gvgvg	194	$2a$08$QVhIY3hoIMixi48elmUsGum596rYH0sb25GBSJdwYEFXTEk3aJ93.	\N	Active	697653	Admin	\N	\N	\N	\N
svbn	fgbhnjm	bhgb@jhhj	rfvev	195	$2a$08$UlHZOOrn54efiw.BG7EZxerxiSrLWzKvAiXlbM5iY2KrpsKhHIkfG	\N	Inactive	gbhn	Accountant	\N	\N	\N	\N
owaish	kkalimllll	a@owaish	ghvgh	198	$2a$08$k0Xm/Ej19k.B.CewVv98CeVzZ9kHdzeTowSswesSUnKc73n.wuDn2	\N	Active	88945562	Accountant	\N	\N	\N	\N
ABD	KHAN	np@mqqqahalodum	s	199	$2a$08$x2oMufuPOjCFRtstSn6rFuab7LvDAB8CmxsozpA9NoVpyXDO64joy	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
ABD	KHAN	np@mqqqahadum	hhs	200	$2a$08$eMW.4SyZJYnixdkCm1Ct..oK8Rpin8lgqCaVLiDVUPVYgq1MMbt1G	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
ABD	KHAN	np@mqqqahaoosodum	hhs	202	$2a$08$JF1o1L/uuc8tYrORdnFoN.4GQ1BHfndNjXy1mKEBzNSqKpBW8VHCu	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
ayush	sdfghj	efewf@jhbjh	sdgsgvfd	203	$2a$08$TzaalUBBDWVO8WSoABQOZeFiKyO/.5UzwICQncaIJ8aWn8sQzjQgC	\N	Inactive	sdfghefewf	Accountant	\N	\N	\N	\N
John	Wick	np@johnaoosodum	hhs	205	$2a$08$9pt7v29QzIBXBTq4KrhPDuPFe.nL2USvLLSISstUCmslZp5um0u16	\N	sdfbhm	94121212	Accountant	\N	\N	\N	\N
abhishek	kumar	dfds@nvv	greg	204	$2a$08$hVG9jvTGG7kO.i88tllOEOgOesqJHYeNMf4Tk478I5X1S7cJAvNby	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRmZHNAbnZ2IiwiZXhwIjoxNTg3NTc3MDYwfQ.pTr8Kv8i7QfhWvfBqN8rWpRO17bVOS3YeCt4iQ69h_g	Inactive	sdfghj	Accountant	\N	\N	\N	\N
John	Wick	np@johnaohnosodum	hhs	206	$2a$08$Z/oiiVa4HAeB6umKQxvJD.I1SG6R1TUI94ccqgJJsMOd6gV0Kcst.	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
owaish	gheqwesdfgrtyem	owaish@kalpppim1	1341232qwertbnbnmyuil	164	owashpp	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzcXdoQHNka2FsaW0iLCJleHAiOjE1ODc0ODgyOTh9.jCwbjRlBTCVg7hRg8cFspDjmVe0S91lwYIH0HcG1HUg	sdfgh	123494112121234fdf1212	\N	\N	\N	\N	\N
sdfghjk	ghj	dgh@jhg	fe	207	$2a$08$TvfY3lR.fMLFFfMKg7gp0eyFKAawNXDA.ludCA1n0Vwe5QaR3sWWq	\N	Inactive	sdfgh	Customer Care	\N	\N	\N	\N
dfghj	ddfgh	hjghj@jgj	gv	208	$2a$08$iYBnmWWxwGZH9D0PxilNIOudBJbltWCjftYm27aXjF9l8r3LAmGyC	\N	Inactive	vvv	Accountant	\N	\N	\N	\N
John	Wick	np@lllqqum	hhs	220	$2a$08$fOUHpXouvmPW3J31eMwsSushac/A7pnv0VmQ6xcgfYkJZ5hx8THQa	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@lllqaaqum	hhs	221	$2a$08$TT.ivWr2cJnDroB3hRKTzOQ6iLFboklAgzszPbzu6zVkMBVoM4.PK	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@lllaaqum	hhs	222	$2a$08$CR4664b82vdA8bHYVHT0aOs/RHOLPZfl5V4FwEy5eI/43CMB6yAMC	\N	sdfbhm	94121212	Adminpp	\N	\N	\N	\N
John	Wick	np@lllqum	hhs	223	$2a$08$NiPqBJgg1DKH7bUFo.19/eL53aloXAm.pHrtkBp9Dtx5XU2LAcz/O	\N	sdfbhm	94121212	Adminpp	\N	\N	\N	\N
John	Wick	np@lllzzqum	hhs	224	$2a$08$n6pDTV7jC8yH5X4Bhj7P7Ov00MpoSHz5BRhdFxzogVpopUTuhbsfW	\N	sdfbhm	94121212	Adminpp	\N	\N	\N	\N
John	Wick	np@llsodum	hhs	209	$2a$08$UckSvwtWPOUvwBljYPvSaOlk8u4pwuT6/0Bg.SwMc7k7DQv7353.G	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@llsodlllum	hhs	211	$2a$08$LYDXkGLR6P57oLzhazOOfOxXgkR7Risahab1qNbqfHV.toOCueGMy	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@llsobbvdlllum	hhs	213	$2a$08$bNJppTIseeI1hN39cr9aqe6iI.S10.My//4INmcqloS7tuMKNrjdW	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@llsobwertyudlllum	hhs	214	$2a$08$UUpiFc5UNONg7hD477spuOetLBRqzKrqSw7fgYZJ.JhfFUIRjotdK	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@llsobweqazyudlllum	hhs	215	$2a$08$WW8gAIaz2UzADBJuZTF.5.rE.bogqK/iUYDliJRbMCskYSGNkUQxG	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	np@llsobwedlllum	hhs	216	$2a$08$.a.CkZ9YRhtNfaIEENoa7uaDooQw230Cf/H6QO9jY33J4igHfhE92	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
abhi	Hirnotia	ashutoshhhirnotia23@gmail.co	Gali no 16 B,Lajjapuri ,chamri road ,Hapur	225	$2a$08$KNUx7RV/N.26TuUIbpyTM.VOEQgLtsV0SdTJU0vBdwE2bONrzKQKu	\N	Inactive	09084768046	Accountant	\N	\N	\N	\N
John	Wick	np@lllzxxzqum	hhs	226	$2a$08$yDtLA0X9VKjWwMcWyBKV8.vqbPEivkUwqO1k26vFmTEtdHApuywAK	\N	sdfbhm	94121212	Admin	\N	\N	\N	\N
aakash	sa	ass@asdfghjk	werg	227	$2a$08$cre.cfqP6WA7EiIzowj2yeTbekkA7dkvnnqPE69QksV1JC.0//q9i	\N	Inactive	sas	Admin	\N	\N	\N	\N
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa	kumar	ahirnotia@ch.iitr.ac.in	dfghkj	132	$2a$08$QsKrZ4Vsgsj.fRw4eLo4jOZ0D4PDQF7BgC9sCoeZ/OnHxHriyjDka	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFoaXJub3RpYUBjaC5paXRyLmFjLmluIiwiZXhwIjoxNTg3NjE4NTU0fQ.29ThHn7P9bKgklwaduvHVAG_KcaFrRrJAig67oZ6a74	Active	1234567892	Admin	\N	\N	\N	\N
John	Wick	np@lllum	hhs	218	$2a$08$wVDVSrhmssn4aS99Gl53TOUb4VeX/m/J9JFTu5AN6w3c7ETyeTF7K	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5wQGxsbHVtIiwiZXhwIjoxNTg3NTg1NDE3fQ._geB3KGG8qgzpwiRf_KYHqKJj7eB4Hdlt61hNIXxQC4	sdfbhm	94121212	Admin	\N	\N	\N	\N
John	Wick	okalim@ch.iitr.ac.in	hhs	240	$2a$08$O/GxbRzbtaY5Kh/zqRzXseB3YiHvu594MeJOAYItG4FLxaHJ545BW	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im9rYWxpbUBjaC5paXRyLmFjLmluIiwiZXhwIjoxNTg3ODU3NDI0fQ.75aa5uBk3ixDN3NJkuEOQVMLNrjXqJxWsXHrCEq2zZk	sdfbhm	94121212	Admin	2020-04-14 21:56:46.385476089+05:30		\N	\N
ABD	KHAN	abd@4321	RB	231	$2a$08$n1CCExpnfC7yBJiQ2g75su4iTyUA.NQ6LzRg848OQ4rn387NU1GgS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiZEA0MzIxIiwiZXhwIjoxNTg3NjEyNzQ0fQ.nh9KiNxMnC64sGHoRa3fFfoRgzl_8DLAF1buymURMiU	inactive	1234	Admin	\N	\N	\N	\N
ABD-OW	KHANnnnnnnnnnnn	kalim@4321	RKB	230	$2a$08$QM6T2WXL9LsGzw1zM.vjW.VjY3IjJo30xYcClFF4vZNxKP4LLa0jW	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltQDQzMjEiLCJleHAiOjE1ODc1OTA4NTV9.cofGchjkEox3CqkAkOcVcvsbYxISAECo5ZfMfN3D8dU	sdfgh	4321	Admin	\N	\N	\N	\N
Ashish123	Hirnotia	a@aakash	sdfgh	229	$2a$08$OCG1Aj4hzf1Q.fYopQC85u0SSdeK4d8zpMxCSUZRa0nhe25Vy//QS	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAYWFrYXNoIiwiZXhwIjoxNTg3NTg3MDk5fQ.DjSuu5q3MxMT_1Ak1oM-kzmzZd4pIF4DhCtxgrMdOEY	Inactive	1234567	Customer Service	\N	\N	\N	\N
Akash	Hirnotia	kalim@321	RKB	228	owas	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltQDMyMSIsImV4cCI6MTU4NzU4NjQwNX0.bPeb0_1sCyIxMDyKI3nKg89WRAH-2idfNmXD6UqomSA	sdfgh	4321	Customer Service	\N	\N	\N	\N
ABD	KHAN	kalimowaish@gmail.com	RB	237	$2a$08$7ptZr.PM1JqwbdGDfdFy5OTt5ZEWCU.DAkj4LFclakKmt3jIrVT0S	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbGltb3dhaXNoQGdtYWlsLmNvbSIsImV4cCI6MTU4NzY5MjUzNn0.YUn0BecJf4irP4WENzngqcMJVX_UcM3nm3I38A0Z5N4	inactive	1234	Admin	2020-04-13 21:11:38.752651831+05:30	24-April	\N	\N
John	Wick	np@llsoqbwedlllum	hhs	238	$2a$08$ZI.SHWn.ecxi6s6tjEQRuuKI8cl/GrIWLDzZAuJU9/hC/S/QpBjpq	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5wQGxsc29xYndlZGxsbHVtIiwiZXhwIjoxNTg3NzAzMDM5fQ.ihpEo6SydhUzKLR7y34f_NxGIrETWjRNI8zPURNNKKg	sdfbhm	94121212	Admin	2020-04-14 00:07:06.223161261+05:30		\N	\N
John	Wick	kumarayush639@gmail.com	hhs	239	$2a$08$urq17hgiHus7t9zAy8kN0uF4dAGjxHeaS7tajrkgkFd0T2JS37Do2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imt1bWFyYXl1c2g2MzlAZ21haWwuY29tIiwiZXhwIjoxNTg3NzA3MTQ2fQ.9nkYg23Q585X6Y3GALUKtrSHxEwOAAzQjsdb2FIJEQw	sdfbhm	94121212	Admin	2020-04-14 00:50:14.636829509+05:30		\N	\N
jason	bourne	d@w	hhs	242	$2a$08$3OSdZvUDuU7z2KIkbjcJCuOU7D..34QlIFyE0IVgreGPJ9B7pHKcK	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRAdyIsImV4cCI6MTU4Nzg2MzczN30.ekiGmtaP19lITciOpZ6sx-wc1y3CB48T1oMYzn0VwAM	sdfbhm	94121212	Accountant	15-04-2020 8:44:46 PM		\N	\N
aakash	hirnotia	a@f	etrdtfyguh	170	$2a$08$p.p22Arzrb05lkHieShEUO6NqD16lWny50/IC8WzwSJTPogZkTCXm	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFAZiIsImV4cCI6MTU4Nzg2NDk4NX0.OsjW8EfDEwqsGNTP5YxNq6kwgANWIRkmDvEPQsbPKuM	Active	9258656829	\N	\N	\N	\N	\N
Jack	Reacher	d@aqqaaaaaaz	hhs	250	$2a$08$Wj2LCikcCTxTzvCpYOt/kewtb7rWOuMhuGxKzf/WpncTdLCzBIOpy	\N	sdfbhm	34567890	Accountant	15-04-2020 11:12:11 PM		\N	\N
Ashish kumar	Singh	asingh2@ch.iitr.ac.in	Rajiv Bhwan, IIT Roorkee	120	$2a$08$Odz87MnbLSN4evI12sVqPecm.YRBRrPhi1PSUooU0onWmBB3BJAHa	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFzaW5naDJAY2guaWl0ci5hYy5pbiIsImV4cCI6MTU4Njg4NzI0OH0.b1S7c_-yt2QIXYwRfSsQe0JaJ0F_HCby5ObiTcQQIcY	Inactive	+918769304216	Accountant	\N	\N	\N	\N
ashutosh	hirnotia	s@sdfghjgvhjvhvbvn	ghfh	232	$2a$08$qvDTY7NpIptml0m2DBpCie8eh.zv/KT6HkkdcIAuaa6/beeIEpyc6	\N	Inactive	5464	Accountant	2020-04-12 21:45:48.808732474+05:30	2020-04-12 21:45:48.808732327+05:30	\N	\N
aakash	asdfghjkl	hj@hjj	ff	233	$2a$08$X819PuQ.YnL2.ExR5oy6ke1Ffw.OZWU2X.O5htVqUsP/dqtKCx6ui	\N	Deleted	00000000000000000000000000000000000000	Accountant	2020-04-12 21:48:38.879969665+05:30	2020-04-12 21:48:38.879969137+05:30	\N	\N
ABD	KHAN	abd@4233321	RB	235	$2a$08$WHT6x7sxkNPZeJp4YCYMceWZeZTokMh1EqHuGiqnvmx1czjIvM09q	\N	inactive	1234	Admin	2020-04-12 22:06:06.449717438+05:30	24-April	\N	\N
Jack	Reacher	d@e	hhs	243	$2a$08$7v2N8N0sPO2Rk6fn.EhcBeY9oTBUAyT9DFFd23JxTY4kpR/gz.pY.	\N	sdfbhm	94121212w	Accountant	15-04-2020 11:2:58 PM		\N	\N
Jack	Reacher	d@z	hhs	244	$2a$08$dZZUnpRgqhyHX9G2KWcJWuaEbd.15rrFBEWQ4vMxJl4dRDKFFbm2K	\N	sdfbhm	94121	Accountant	15-04-2020 11:5:22 PM		\N	\N
Jack	Reacher	d@az	hhs	246	$2a$08$paxviig5.UQF.OjIK.7Rje/1TpTzlvS/qLSN63x1Redi8WwtCuGmS	\N	sdfbhm	9414567890-678901	Accountant	15-04-2020 11:5:45 PM		\N	\N
Jack	Reacher	d@aqqaaqwerhaaaaz	hhs	251	$2a$08$Yeq2bOOyqFpjrA79JPVLwuA4/fS0BvThWYETa7e1Mz5vhEAEeT4ri	\N	sdfbhm	34567890	Accountant	15-04-2020 11:13:25 PM		\N	\N
Jack	Reacher	d@aqqz	hhs	248	$2a$08$ElMLL7aN61J6POd24uoOPOZhkNK5e2OBdHi.UAy6pVrlD/hpsizhO	\N	sdfbhm	9414567890-67890134567890	Accountant	15-04-2020 11:6:2 PM		\N	\N
Jack	Reacher	d@aqqaaz	hhs	249	$2a$08$iv8E4ipiQyyUgSssTIQ/K.lS8mvlMDf8otN.Z3p1mKedLLEQ5b8Li	\N	sdfbhm	9	Accountant	15-04-2020 11:6:44 PM		\N	\N
Jack	Reacher	d@erq	hhs	256	$2a$08$z0AtDUg2jQ116h9Fdm5n8.wnivYkPZf4IASD8TdfJJFCabVAolHV6	\N	sdfbhm	7894567894	Accountant	16-04-2020 3:23:33 AM		\N	\N
Jack	Reacher	d@aqqttt	hhs	252	$2a$08$U10OG2PvXIW7m2QhGL7e3OvNoCBfqyyOor8V4JcAyFgg8VhAU3SNi	\N	sdfbhm	34567890	Accountant	15-04-2020 11:13:43 PM		\N	\N
Jack	Reacher	d@er	hhs	253	$2a$08$6yF2Xu63SKmxGR2g/Cuc3eA3OJ8LbFWssO9iJnjXJaAC.srVg9862	\N	sdfbhm	7894567894	Accountant	16-04-2020 2:32:26 AM		\N	\N
Jack	Reacher	d@erqq	hhs	258	$2a$08$cDUP8bp0QLOa75AVnbHh2.UzGod01J8iSQMX4TPH2g/jEtfT7YskW	\N	sdfbhm	7894567894	Accountant	16-04-2020 3:24:6 AM	20-02-2020	\N	\N
Jack	Reacher	d@eqrqq	hhs	260	$2a$08$xp1QBG.8fWa0/Km8oQWw.OYFzuK6/5e5p3x0pZLQtXf27eFDDkfey	\N	sdfbhm	7894567894	Accountant	16-04-2020 3:33:40 AM	20-02-2020	male	\N
Jack	Reacher	d@eqrqwwq	hhs	261	$2a$08$ahjhN5jLS6gVJZpteHL6/Ou3JbkaeUfgGsJqQQTihVj76epFZpUnm	\N	sdfbhm	0908877	Accountant	16-04-2020 3:13:9 PM	20-02-2020	male	\N
Jack	Reacher	ow@kal	hhs	263	$2a$08$5XplXRSZOxisAENfK7uS1ewWGmFPghxb0Mnc5RD9Z0uZEap.Md5dO	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93QGthbCIsImV4cCI6MTU4Nzk2MDAzN30.O43ZtCPo4fh2N57Li1t5xH5K8a2bp8ukqopcn_P5Rrw	sdfbhm	0908877	Admin	16-04-2020 9:50:1 PM	20-02-2020	male	\N
Owaish 	Kalim	owaishkalim@gmail.com	hhs	241	$2a$08$dtS.dJW2SdTaDv57fpRR0Oya35LqwlpeiQgsycUC1MEzk7vojgcd2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im93YWlzaGthbGltQGdtYWlsLmNvbSIsImV4cCI6MTU4Nzk2MTg2Mn0.4AS5rWgoiaXztJpYCAcCrq0u9OVtxV_CMn9vXDoS1JM	Active	94121212	Admin	15-04-2020 8:41:41 PM		Male	73324651_124404222298899_4376515792804511744_n.jpg
\.


--
-- Data for Name: slh_transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.slh_transactions ("Customer_Order_Id", "Customer_Id", "Therapist_Gender", "Massage_For", "Customer_Address", "Transaction_Mode", "Bank_Type", "Merchant_Transaction_Id", "Massage_Duration", "Customer_Name", "Customer_Souls_Id", "Payment_Gateway_Id", "Latitude", "Longitude", "Payment_Gateway_Mode", "Total_Order_Amount", "Number_Of_Therapist_Required", "Pincode", "Slot_Time", "Slot_Date", "CreatedAt") FROM stdin;
\N	38	male	cbn	ertyj	dfghj	cvbn	158628053438	trew	kalim	2020040738	sdfghj	qwerty	asdfg	dfgh	\N	\N	\N	\N	\N	\N
18	38	male	cbn	ertyj	dfghj	cvbn	158628053438	trew	kalim	2020040738	sdfghj	qwerty	asdfg	dfgh	\N	\N	\N	\N	\N	\N
19	39	male	cbn	ertyj	dfghj	cvbn	158628813539	trew	kalim	2020040839	sdfghj	qwerty	asdfg	dfgh	\N	\N	\N	\N	\N	\N
19	39	male	cbn	ertyj	dfghj	cvbn	158628813539	trew	kalim	2020040839	sdfghj	qwerty	asdfg	dfgh	\N	\N	\N	\N	\N	\N
21	40	male	myself	ertyj	debit	sbi	158634174540	trew	kalim	2020040840	sherlock	qwerty	asdfg	cheque	\N	\N	\N	\N	\N	\N
21	40	male	myself	ertyj	debit	sbi	158634174540	trew	kalim	2020040840	sherlock	qwerty	asdfg	cheque	\N	\N	\N	\N	\N	\N
20	40	male	myself	ertyjf,m,lgk	dfghjwwre	cvbn3rr	158628820740	trewfklkf	kalim	2020040840	sdfghjfsg	qwertyerwtt	asdfgeet	dfghefsfs	\N	\N	\N	\N	\N	\N
20	40	male	myself	ertyjf,m,lgk	dfghjwwre	cvbn3rr	158628820740	trewfklkf	kalim	2020040840	sdfghjfsg	qwertyerwtt	asdfgeet	dfghefsfs	\N	\N	\N	\N	\N	\N
20	40	male	myself	ertyjf,m,lgk	dfghjwwre	cvbn3rr	158628820740	trewfklkf	kalim	2020040840	sdfghjfsg	qwertyerwtt	asdfgeet	dfghefsfs	\N	\N	\N	\N	\N	\N
20	40	male	myself	ertyjf,m,lgk	dfghjwwre	cvbn3rr	158628820740	trewfklkf	kalim	2020040840	sdfghjfsg	qwertyerwtt	asdfgeet	dfghefsfs	\N	\N	\N	\N	\N	\N
29	58	male	cbn	Ny	dfghj	cvbn	1586933802-58	trew	Kashyap	2020041558	sdfghj	qwerty	asdfg	dfgh	10000	2	65449	\N	\N	\N
30	60	male	cbn	DL	dfghj	cvbn	1586934034-60	trew	Vinay	2020041560	sdfghj	qwerty	asdfg	dfgh	10000	2	13596	\N	\N	\N
31	64	male	cbn	DL	dfghj	cvbn	1587057071-64	trew	Vinay	2020041664	sdfghj	5.239	1.256	dfgh	10000	2	13596	5:5	05-05-2020	16-04-2020 10:41:11 PM
31	64	male	cbn	DL	dfghj	cvbn	1587057071-64	trew	Vinay	2020041664	sdfghj	5.239	1.256	dfgh	10000	2	13596	5:5	05-05-2020	16-04-2020 10:41:11 PM
31	64	male	cbn	DL	dfghj	cvbn	1587057071-64	trew	Vinay	2020041664	sdfghj	5.239	1.256	dfgh	10000	2	13596	5:5	05-05-2020	16-04-2020 10:41:11 PM
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, "Name", "Email", "Password", "Token", "CreatedAt", "UpdatedAt", "Role") FROM stdin;
\.


--
-- Name: slh_assign_customer_with_partner_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_assign_customer_with_partner_Id_seq"', 1, false);


--
-- Name: slh_communication_templates_Comm.Template_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_communication_templates_Comm.Template_Id_seq"', 8, true);


--
-- Name: slh_customers_Customer_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_Customer_Id_seq"', 67, true);


--
-- Name: slh_customers_pending_orders_Customer_Primary_Order_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_customers_pending_orders_Customer_Primary_Order_Id_seq"', 31, true);


--
-- Name: slh_partners_Partner_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_partners_Partner_Id_seq"', 24, true);


--
-- Name: slh_roles_RoleId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_roles_RoleId_seq"', 1, false);


--
-- Name: slh_roles_Role_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_roles_Role_Id_seq"', 6, true);


--
-- Name: slh_souls_settings_Souls_Setting_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_souls_settings_Souls_Setting_Id_seq"', 2, true);


--
-- Name: slh_team_has_role_Team_Has_Role_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_team_has_role_Team_Has_Role_Id_seq"', 1, true);


--
-- Name: slh_teams_TeamId_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."slh_teams_TeamId_seq"', 263, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 28, true);


--
-- Name: slh_customers Customer_Mobile_No; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_customers
    ADD CONSTRAINT "Customer_Mobile_No" UNIQUE ("Customer_Mobile_No");


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
-- Name: slh_assign_customer_with_partner slh_assign_customer_with_partner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_assign_customer_with_partner
    ADD CONSTRAINT slh_assign_customer_with_partner_pkey PRIMARY KEY ("Id");


--
-- Name: slh_communication_templates slh_communication_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_communication_templates
    ADD CONSTRAINT slh_communication_templates_pkey PRIMARY KEY ("Comm.Template_Id");


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
-- Name: slh_souls_settings slh_souls_settings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_souls_settings
    ADD CONSTRAINT slh_souls_settings_pkey PRIMARY KEY ("Souls_Setting_Id");


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
-- Name: slh_team_has_role slh_team_has_role_TeamId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.slh_team_has_role
    ADD CONSTRAINT "slh_team_has_role_TeamId_fkey" FOREIGN KEY ("Team_Id") REFERENCES public.slh_teams("TeamId") NOT VALID;


--
-- PostgreSQL database dump complete
--

