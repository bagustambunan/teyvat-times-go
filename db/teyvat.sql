--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)

-- Started on 2022-09-04 19:18:56 WIB

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
-- TOC entry 213 (class 1259 OID 18609)
-- Name: addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.addresses (
    id bigint NOT NULL,
    street character varying NOT NULL,
    city character varying NOT NULL,
    state character varying NOT NULL,
    country character varying NOT NULL,
    postal_code character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.addresses OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 18608)
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.addresses_id_seq OWNER TO postgres;

--
-- TOC entry 3611 (class 0 OID 0)
-- Dependencies: 212
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- TOC entry 241 (class 1259 OID 18944)
-- Name: gift_claim_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gift_claim_items (
    id bigint NOT NULL,
    gift_id integer NOT NULL,
    gift_claim_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.gift_claim_items OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 18943)
-- Name: gift_claim_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gift_claim_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gift_claim_items_id_seq OWNER TO postgres;

--
-- TOC entry 3612 (class 0 OID 0)
-- Dependencies: 240
-- Name: gift_claim_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gift_claim_items_id_seq OWNED BY public.gift_claim_items.id;


--
-- TOC entry 237 (class 1259 OID 18907)
-- Name: gift_claim_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gift_claim_statuses (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.gift_claim_statuses OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 18919)
-- Name: gift_claims; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gift_claims (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    address_id bigint NOT NULL,
    status_id integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.gift_claims OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 18918)
-- Name: gift_claims_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gift_claims_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gift_claims_id_seq OWNER TO postgres;

--
-- TOC entry 3613 (class 0 OID 0)
-- Dependencies: 238
-- Name: gift_claims_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gift_claims_id_seq OWNED BY public.gift_claims.id;


--
-- TOC entry 236 (class 1259 OID 18889)
-- Name: gifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gifts (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    image_id bigint DEFAULT 3 NOT NULL,
    stock integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.gifts OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 18587)
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    url text NOT NULL,
    alt_text character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.images OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 18586)
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.images_id_seq OWNER TO postgres;

--
-- TOC entry 3614 (class 0 OID 0)
-- Dependencies: 209
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- TOC entry 220 (class 1259 OID 18685)
-- Name: post_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_categories (
    id bigint NOT NULL,
    name character varying NOT NULL,
    color character varying,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.post_categories OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 18684)
-- Name: post_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_categories_id_seq OWNER TO postgres;

--
-- TOC entry 3615 (class 0 OID 0)
-- Dependencies: 219
-- Name: post_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_categories_id_seq OWNED BY public.post_categories.id;


--
-- TOC entry 218 (class 1259 OID 18673)
-- Name: post_tiers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_tiers (
    id integer NOT NULL,
    name character varying NOT NULL,
    mora_required integer NOT NULL,
    color character varying,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.post_tiers OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 18764)
-- Name: post_unlocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_unlocks (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.post_unlocks OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 18763)
-- Name: post_unlocks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_unlocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_unlocks_id_seq OWNER TO postgres;

--
-- TOC entry 3616 (class 0 OID 0)
-- Dependencies: 225
-- Name: post_unlocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_unlocks_id_seq OWNED BY public.post_unlocks.id;


--
-- TOC entry 222 (class 1259 OID 18696)
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    post_tier_id integer DEFAULT 1 NOT NULL,
    post_category_id bigint NOT NULL,
    title character varying NOT NULL,
    content text NOT NULL,
    slug character varying NOT NULL,
    summary character varying NOT NULL,
    img_thumbnail_id bigint DEFAULT 2 NOT NULL,
    img_content_id bigint DEFAULT 3 NOT NULL,
    created_by_id bigint NOT NULL,
    updated_by_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 18695)
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO postgres;

--
-- TOC entry 3617 (class 0 OID 0)
-- Dependencies: 221
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- TOC entry 211 (class 1259 OID 18597)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 18782)
-- Name: subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subscriptions (
    id integer NOT NULL,
    name character varying NOT NULL,
    price integer NOT NULL,
    mora_amount integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.subscriptions OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 18849)
-- Name: transaction_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_statuses (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transaction_statuses OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 18861)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    subscription_id integer NOT NULL,
    status_id integer NOT NULL,
    gross_total integer NOT NULL,
    net_total integer NOT NULL,
    user_voucher_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 18860)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3618 (class 0 OID 0)
-- Dependencies: 234
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 243 (class 1259 OID 18964)
-- Name: user_gifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_gifts (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    gift_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    is_claimed integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.user_gifts OWNER TO postgres;

--
-- TOC entry 242 (class 1259 OID 18963)
-- Name: user_gifts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_gifts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_gifts_id_seq OWNER TO postgres;

--
-- TOC entry 3619 (class 0 OID 0)
-- Dependencies: 242
-- Name: user_gifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_gifts_id_seq OWNED BY public.user_gifts.id;


--
-- TOC entry 224 (class 1259 OID 18742)
-- Name: user_post_activities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_post_activities (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    views_count integer DEFAULT 0 NOT NULL,
    is_liked integer DEFAULT 0 NOT NULL,
    is_shared integer DEFAULT 0 NOT NULL,
    date_liked timestamp without time zone,
    date_shared timestamp without time zone,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_post_activities OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 18741)
-- Name: user_post_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_post_activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_post_activities_id_seq OWNER TO postgres;

--
-- TOC entry 3620 (class 0 OID 0)
-- Dependencies: 223
-- Name: user_post_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_post_activities_id_seq OWNED BY public.user_post_activities.id;


--
-- TOC entry 217 (class 1259 OID 18655)
-- Name: user_referrals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_referrals (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    referrer_user_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_referrals OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 18654)
-- Name: user_referrals_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_referrals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_referrals_id_seq OWNER TO postgres;

--
-- TOC entry 3621 (class 0 OID 0)
-- Dependencies: 216
-- Name: user_referrals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_referrals_id_seq OWNED BY public.user_referrals.id;


--
-- TOC entry 229 (class 1259 OID 18794)
-- Name: user_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_subscriptions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    subscription_id integer NOT NULL,
    date_start date NOT NULL,
    date_ended date NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_subscriptions OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 18793)
-- Name: user_subscriptions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_subscriptions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_subscriptions_id_seq OWNER TO postgres;

--
-- TOC entry 3622 (class 0 OID 0)
-- Dependencies: 228
-- Name: user_subscriptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_subscriptions_id_seq OWNED BY public.user_subscriptions.id;


--
-- TOC entry 232 (class 1259 OID 18830)
-- Name: user_vouchers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_vouchers (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    voucher_id integer NOT NULL,
    date_expired date NOT NULL,
    is_used integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_vouchers OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 18829)
-- Name: user_vouchers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_vouchers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_vouchers_id_seq OWNER TO postgres;

--
-- TOC entry 3623 (class 0 OID 0)
-- Dependencies: 231
-- Name: user_vouchers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_vouchers_id_seq OWNED BY public.user_vouchers.id;


--
-- TOC entry 215 (class 1259 OID 18620)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    role_id integer DEFAULT 2 NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    name character varying NOT NULL,
    phone character varying NOT NULL,
    address_id bigint,
    referral_code character varying NOT NULL,
    profile_pic_id bigint DEFAULT 1 NOT NULL,
    mora integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 18619)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3624 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 230 (class 1259 OID 18812)
-- Name: vouchers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vouchers (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text NOT NULL,
    image_id bigint DEFAULT 3 NOT NULL,
    amount integer NOT NULL,
    code character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.vouchers OWNER TO postgres;

--
-- TOC entry 3270 (class 2604 OID 18612)
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- TOC entry 3329 (class 2604 OID 18947)
-- Name: gift_claim_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_items ALTER COLUMN id SET DEFAULT nextval('public.gift_claim_items_id_seq'::regclass);


--
-- TOC entry 3325 (class 2604 OID 18922)
-- Name: gift_claims id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claims ALTER COLUMN id SET DEFAULT nextval('public.gift_claims_id_seq'::regclass);


--
-- TOC entry 3265 (class 2604 OID 18590)
-- Name: images id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- TOC entry 3284 (class 2604 OID 18688)
-- Name: post_categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_categories ALTER COLUMN id SET DEFAULT nextval('public.post_categories_id_seq'::regclass);


--
-- TOC entry 3299 (class 2604 OID 18767)
-- Name: post_unlocks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_unlocks ALTER COLUMN id SET DEFAULT nextval('public.post_unlocks_id_seq'::regclass);


--
-- TOC entry 3287 (class 2604 OID 18699)
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- TOC entry 3316 (class 2604 OID 18864)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3332 (class 2604 OID 18967)
-- Name: user_gifts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts ALTER COLUMN id SET DEFAULT nextval('public.user_gifts_id_seq'::regclass);


--
-- TOC entry 3293 (class 2604 OID 18745)
-- Name: user_post_activities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_post_activities ALTER COLUMN id SET DEFAULT nextval('public.user_post_activities_id_seq'::regclass);


--
-- TOC entry 3279 (class 2604 OID 18658)
-- Name: user_referrals id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_referrals ALTER COLUMN id SET DEFAULT nextval('public.user_referrals_id_seq'::regclass);


--
-- TOC entry 3304 (class 2604 OID 18797)
-- Name: user_subscriptions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_subscriptions ALTER COLUMN id SET DEFAULT nextval('public.user_subscriptions_id_seq'::regclass);


--
-- TOC entry 3310 (class 2604 OID 18833)
-- Name: user_vouchers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers ALTER COLUMN id SET DEFAULT nextval('public.user_vouchers_id_seq'::regclass);


--
-- TOC entry 3273 (class 2604 OID 18623)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3575 (class 0 OID 18609)
-- Dependencies: 213
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (1, 'Dandelion St.127', 'Windrise', 'Mondstadt', 'Teyvat', 'MOND38', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (2, 'Mt. Tianheng', 'Minlin', 'Liyue', 'Teyvat', 'LYE77012', '2022-08-30 03:14:30.363034', '2022-08-30 03:14:30.363034', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (3, 'Inazuma City', 'Narukami Islan', 'Inazuma', 'Teyvat', 'INMA645117', '2022-08-30 03:18:21.430858', '2022-08-30 03:18:21.430858', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (4, 'Favonious St. 8', 'Dodoko', 'Mondstadt', 'Teyvat', 'DDO001', '2022-08-30 13:04:00.090645', '2022-08-30 13:04:00.090645', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (5, 'Soil St. 777', 'Alchemy of Favonious', 'Mondstadt', 'Teyvat', 'ALCH812', '2022-08-30 13:27:35.088577', '2022-08-30 13:27:35.088577', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (6, 'Wolfhooks St. 001', 'Wolvendom', 'Mondstadt', 'Teyvat', 'WOLV381', '2022-09-01 23:12:08.237104', '2022-09-01 23:12:08.237104', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (7, 'Good Alchemy', 'Anemo Hypostasis', 'Mondstadt', 'Teyvat', 'ANM659', '2022-09-04 18:10:17.530955', '2022-09-04 18:10:17.530955', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (8, 'Megistus St. 78', 'Astrologist', 'Mondstadt', 'Teyvat', 'ASTR009', '2022-09-04 18:12:29.143362', '2022-09-04 18:12:29.143362', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (9, 'Tempus Fugit', 'Knights of Favonius', 'Mondstadt', 'Teyvat', 'LIBR81111', '2022-09-04 18:14:03.750597', '2022-09-04 18:14:03.750597', NULL);
INSERT INTO public.addresses (id, street, city, state, country, postal_code, created_at, updated_at, deleted_at) VALUES (10, 'Leo Minor', 'Dandelions City', 'Mondstadt', 'Teyvat', 'DANDE3', '2022-09-04 18:15:40.450552', '2022-09-04 18:15:40.450552', NULL);


--
-- TOC entry 3603 (class 0 OID 18944)
-- Dependencies: 241
-- Data for Name: gift_claim_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gift_claim_items (id, gift_id, gift_claim_id, created_at, updated_at, deleted_at) VALUES (21, 1, 13, '2022-09-04 19:13:43.300167', '2022-09-04 19:13:43.300167', NULL);
INSERT INTO public.gift_claim_items (id, gift_id, gift_claim_id, created_at, updated_at, deleted_at) VALUES (22, 2, 13, '2022-09-04 19:13:43.300167', '2022-09-04 19:13:43.300167', NULL);


--
-- TOC entry 3599 (class 0 OID 18907)
-- Dependencies: 237
-- Data for Name: gift_claim_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gift_claim_statuses (id, name, created_at, updated_at, deleted_at) VALUES (1, 'Draft', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.gift_claim_statuses (id, name, created_at, updated_at, deleted_at) VALUES (2, 'Processing', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.gift_claim_statuses (id, name, created_at, updated_at, deleted_at) VALUES (3, 'On Delivery', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.gift_claim_statuses (id, name, created_at, updated_at, deleted_at) VALUES (5, 'Delivered', '2022-09-04 15:59:08.541177', '2022-09-04 15:59:08.541177', NULL);
INSERT INTO public.gift_claim_statuses (id, name, created_at, updated_at, deleted_at) VALUES (4, 'Rejected', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3601 (class 0 OID 18919)
-- Dependencies: 239
-- Data for Name: gift_claims; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gift_claims (id, user_id, address_id, status_id, created_at, updated_at, deleted_at) VALUES (13, 5, 5, 3, '2022-09-04 19:13:43.298666', '2022-09-04 19:13:43.298666', NULL);


--
-- TOC entry 3598 (class 0 OID 18889)
-- Dependencies: 236
-- Data for Name: gifts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gifts (id, name, description, image_id, stock, created_at, updated_at, deleted_at) VALUES (3, 'Keyboard', '-', 3, 10, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.gifts (id, name, description, image_id, stock, created_at, updated_at, deleted_at) VALUES (1, 'Windsong Lyre', '-', 4, 47, '2022-08-30 02:31:39.662158', '2022-09-04 19:14:08.855956', NULL);
INSERT INTO public.gifts (id, name, description, image_id, stock, created_at, updated_at, deleted_at) VALUES (2, 'Wings of Descension', '-', 5, 22, '2022-08-30 02:31:39.662158', '2022-09-04 19:14:08.858358', NULL);


--
-- TOC entry 3572 (class 0 OID 18587)
-- Dependencies: 210
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.images (id, url, alt_text, created_at, updated_at, deleted_at) VALUES (3, 'https://static.zerochan.net/Lumine.%28Genshin.Impact%29.full.3126229.jpg', 'Default image', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.images (id, url, alt_text, created_at, updated_at, deleted_at) VALUES (1, 'https://static.zerochan.net/Lumine.%28Genshin.Impact%29.full.3126229.jpg', 'Default profile pic', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.images (id, url, alt_text, created_at, updated_at, deleted_at) VALUES (2, 'https://asset.kompas.com/crops/EGyrhYPJi_psXlT29Sq3KyHMR0E=/40x0:904x576/750x500/data/photo/2021/09/29/6153c38932e38.jpeg', 'Default thumbnail', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.images (id, url, alt_text, created_at, updated_at, deleted_at) VALUES (4, 'https://static.wikia.nocookie.net/gensin-impact/images/4/4a/Item_Windsong_Lyre.png', 'Windsong Lyre', '2022-09-02 10:06:30.32303', '2022-09-02 10:06:30.32303', NULL);
INSERT INTO public.images (id, url, alt_text, created_at, updated_at, deleted_at) VALUES (5, 'https://static.wikia.nocookie.net/gensin-impact/images/6/6d/Item_Wings_of_Descension.png', 'Wings of Descension', '2022-09-02 10:08:27.693755', '2022-09-02 10:08:27.693755', NULL);


--
-- TOC entry 3582 (class 0 OID 18685)
-- Dependencies: 220
-- Data for Name: post_categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_categories (id, name, color, created_at, updated_at, deleted_at) VALUES (1, 'Politic', '#d16262', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.post_categories (id, name, color, created_at, updated_at, deleted_at) VALUES (2, 'Economy', '#fcc142', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.post_categories (id, name, color, created_at, updated_at, deleted_at) VALUES (3, 'Sport', '#6b46e3', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.post_categories (id, name, color, created_at, updated_at, deleted_at) VALUES (4, 'Entertainment', '#32afbf', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3580 (class 0 OID 18673)
-- Dependencies: 218
-- Data for Name: post_tiers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_tiers (id, name, mora_required, color, created_at, updated_at, deleted_at) VALUES (1, 'Free', 0, '#adadad', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.post_tiers (id, name, mora_required, color, created_at, updated_at, deleted_at) VALUES (2, 'Premium', 1, '#d49455', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.post_tiers (id, name, mora_required, color, created_at, updated_at, deleted_at) VALUES (3, 'VIP', 2, '#3e92e6', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3588 (class 0 OID 18764)
-- Dependencies: 226
-- Data for Name: post_unlocks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (1, 4, 6, '2022-08-30 13:29:10.687922', '2022-08-30 13:29:10.687922', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (2, 4, 7, '2022-08-30 22:16:20.151548', '2022-08-30 22:16:20.151548', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (3, 4, 7, '2022-08-31 14:59:50.290108', '2022-08-31 14:59:50.290108', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (4, 4, 7, '2022-08-31 14:59:50.297044', '2022-08-31 14:59:50.297044', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (5, 4, 7, '2022-08-31 15:00:10.115508', '2022-08-31 15:00:10.115508', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (6, 4, 7, '2022-08-31 15:00:10.127696', '2022-08-31 15:00:10.127696', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (7, 4, 7, '2022-08-31 15:00:15.985014', '2022-08-31 15:00:15.985014', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (8, 2, 7, '2022-08-31 16:26:44.928629', '2022-08-31 16:26:44.928629', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (9, 2, 4, '2022-08-31 16:36:31.968781', '2022-08-31 16:36:31.968781', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (10, 1, 7, '2022-09-01 16:41:33.428743', '2022-09-01 16:41:33.428743', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (11, 4, 7, '2022-09-02 12:54:57.121514', '2022-09-02 12:54:57.121514', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (12, 4, 7, '2022-09-02 12:55:26.607861', '2022-09-02 12:55:26.607861', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (13, 4, 7, '2022-09-02 12:56:02.271722', '2022-09-02 12:56:02.271722', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (14, 3, 9, '2022-09-04 18:59:51.362501', '2022-09-04 18:59:51.362501', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (15, 3, 6, '2022-09-04 19:00:00.199547', '2022-09-04 19:00:00.199547', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (16, 9, 12, '2022-09-04 19:02:20.009642', '2022-09-04 19:02:20.009642', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (17, 9, 6, '2022-09-04 19:02:27.57293', '2022-09-04 19:02:27.57293', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (18, 9, 3, '2022-09-04 19:02:46.483921', '2022-09-04 19:02:46.483921', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (19, 10, 13, '2022-09-04 19:05:28.396489', '2022-09-04 19:05:28.396489', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (20, 10, 10, '2022-09-04 19:05:48.873129', '2022-09-04 19:05:48.873129', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (21, 10, 9, '2022-09-04 19:05:53.837607', '2022-09-04 19:05:53.837607', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (22, 10, 4, '2022-09-04 19:06:30.301646', '2022-09-04 19:06:30.301646', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (23, 7, 9, '2022-09-04 19:09:05.590967', '2022-09-04 19:09:05.590967', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (24, 7, 10, '2022-09-04 19:09:15.43382', '2022-09-04 19:09:15.43382', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (25, 7, 6, '2022-09-04 19:09:25.602325', '2022-09-04 19:09:25.602325', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (26, 5, 10, '2022-09-04 19:13:08.018086', '2022-09-04 19:13:08.018086', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (27, 5, 7, '2022-09-04 19:13:12.941792', '2022-09-04 19:13:12.941792', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (28, 5, 6, '2022-09-04 19:13:17.578513', '2022-09-04 19:13:17.578513', NULL);
INSERT INTO public.post_unlocks (id, user_id, post_id, created_at, updated_at, deleted_at) VALUES (29, 5, 3, '2022-09-04 19:13:22.876571', '2022-09-04 19:13:22.876571', NULL);


--
-- TOC entry 3584 (class 0 OID 18696)
-- Dependencies: 222
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (1, 1, 1, 'Mondstadt signs a deal with Fatui', 'Test content', 'mondstadt-signs-a-deal-with-fatui', 'What do you think, Jean?', 2, 3, 1, 1, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (5, 1, 3, 'Knights of Favonius held a hilichurl hunting competition', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'knights-of-favonius-held-a-hilichurl-hunting-competition', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 1, 1, '2022-08-30 03:32:09.665955', '2022-08-30 03:32:09.665955', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (6, 3, 1, 'Who is Barbatos and where is he now?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'who-is-barbatos-and-where-is-he-now', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 1, 1, '2022-08-30 03:33:24.658018', '2022-08-30 03:33:24.658018', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (7, 3, 4, 'Osmanthus wine taste the same as I remember', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'osmanthus-wine-taste-the-same-as-i-remember', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 2, 2, '2022-08-30 03:34:33.375571', '2022-08-30 03:34:33.375571', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (2, 1, 1, 'Inazuma Closing It''s Gate From the Outside World Again?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'inazuma-closing-it-s-gate-from-the-outside-world-again', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-08-30 03:23:43.019342', '2022-08-30 03:23:43.019342', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (3, 3, 2, 'Ningguang has just announced an important matter regarding financial policy', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'ningguang-has-just-announced-an-important-matter-regarding-financial-policy', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-08-30 03:27:13.758213', '2022-08-30 03:27:13.758213', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (8, 1, 2, 'Liben finally retired', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'liben-finally-retired', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 2, 2, '2022-08-30 03:45:27.676783', '2022-08-30 03:45:27.676783', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (4, 2, 4, 'This is why Diluc''s tavern is the best tavern in town', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'why-diluc-s-tavern-is-the-best-tavern-in-town', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 1, 1, '2022-08-30 03:29:34.004397', '2022-08-30 03:29:34.004397', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (9, 2, 3, 'Lesser Lord Kusanali is the current Dendro Archon', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'lesser-lord-kusanali-is-the-current-dendro-archon', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:19:25.029148', '2022-09-04 18:19:25.029148', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (10, 2, 2, 'People who encounter misfortunes in the Avidya Forest are lucky to meet a Forest Watcher named Tighnari', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'people-who-encounter-misfortunes-in-the-avidya-forest-are-lucky-to-meet-a-forest-watcher-named-tighnari', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:22:54.334267', '2022-09-04 18:22:54.334267', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (14, 1, 2, 'Time for.... Retribution', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'time-for-retribution', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:58:24.883304', '2022-09-04 18:58:24.883304', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (12, 2, 1, 'The civilization on Tsurumi Island comes to an end through unknown means', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'he-civilization-on-tsurumi-island-comes-to-an-end-through-unknown-means', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:26:22.450365', '2022-09-04 18:26:22.450365', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (13, 3, 3, 'The Archon War comes to an end as the final divine seat in Celestia is claimed', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'the-archon-war-comes-to-an-end-as-the-final-divine-seat-in-celestia-is-claimed', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:28:18.139363', '2022-09-04 18:28:18.139363', NULL);
INSERT INTO public.posts (id, post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id, created_at, updated_at, deleted_at) VALUES (11, 1, 4, 'Early civilizations in the Lisha region of Liyue are established at Dunyu Ruins, Lingju Pass, Qingxu Pool', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'early-civilizations-in-the-lisha-region-of-liyue-are-established-at-dunyu-ruins-lingju-pass-qingxu-pool', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 2, 3, 3, 3, '2022-09-04 18:25:05.80598', '2022-09-04 18:25:05.80598', NULL);


--
-- TOC entry 3573 (class 0 OID 18597)
-- Dependencies: 211
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (1, 'admin', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (2, 'user', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3589 (class 0 OID 18782)
-- Dependencies: 227
-- Data for Name: subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.subscriptions (id, name, price, mora_amount, created_at, updated_at, deleted_at) VALUES (1, 'Standard', 30000, 5, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.subscriptions (id, name, price, mora_amount, created_at, updated_at, deleted_at) VALUES (2, 'Premium', 50000, 10, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.subscriptions (id, name, price, mora_amount, created_at, updated_at, deleted_at) VALUES (3, 'Gold', 90000, 20, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3595 (class 0 OID 18849)
-- Dependencies: 233
-- Data for Name: transaction_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transaction_statuses (id, name, created_at, updated_at, deleted_at) VALUES (1, 'Waiting for payment', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.transaction_statuses (id, name, created_at, updated_at, deleted_at) VALUES (2, 'Processing', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.transaction_statuses (id, name, created_at, updated_at, deleted_at) VALUES (3, 'Completed', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.transaction_statuses (id, name, created_at, updated_at, deleted_at) VALUES (4, 'Canceled', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3597 (class 0 OID 18861)
-- Dependencies: 235
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (1, 1, 1, 2, 30000, 5000, 1, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (2, 4, 2, 3, 50000, 50000, NULL, '2022-08-30 13:04:21.452523', '2022-08-30 13:04:21.452583', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (3, 4, 1, 3, 30000, 30000, NULL, '2022-08-30 19:13:30.68225', '2022-08-30 19:13:30.682303', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (4, 2, 2, 1, 50000, 50000, NULL, '2022-08-31 16:25:00.875033', '2022-08-31 16:25:00.875228', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (5, 2, 2, 3, 50000, 50000, NULL, '2022-08-31 16:25:29.302905', '2022-08-31 16:25:29.30302', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (6, 5, 3, 3, 90000, 90000, NULL, '2022-08-31 23:02:28.092173', '2022-08-31 23:02:28.092282', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (7, 5, 2, 3, 50000, 50000, NULL, '2022-08-31 23:03:04.442445', '2022-08-31 23:03:04.442515', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (8, 1, 1, 3, 30000, 5000, 1, '2022-09-01 00:28:21.454653', '2022-09-01 00:28:21.454736', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (9, 1, 2, 3, 50000, 50000, NULL, '2022-09-01 16:15:37.59822', '2022-09-01 16:15:37.598276', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (10, 1, 3, 3, 90000, 40000, 2, '2022-09-01 16:18:22.813932', '2022-09-01 16:18:22.814028', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (11, 1, 2, 3, 50000, 50000, NULL, '2022-09-01 16:42:02.288415', '2022-09-01 16:42:02.288504', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (12, 1, 2, 1, 50000, 50000, NULL, '2022-09-01 17:17:20.403547', '2022-09-01 17:17:20.4037', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (13, 4, 2, 3, 50000, 50000, NULL, '2022-09-02 00:23:57.508151', '2022-09-02 00:23:57.508252', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (14, 4, 1, 3, 30000, 30000, NULL, '2022-09-02 00:25:14.06352', '2022-09-02 00:25:14.063611', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (15, 4, 3, 3, 90000, 90000, NULL, '2022-09-02 12:56:59.074738', '2022-09-02 12:56:59.074797', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (16, 3, 2, 3, 50000, 50000, NULL, '2022-09-04 18:59:08.720719', '2022-09-04 18:59:08.720905', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (17, 9, 1, 1, 30000, 30000, NULL, '2022-09-04 19:00:29.779744', '2022-09-04 19:00:29.779945', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (18, 9, 3, 3, 90000, 90000, NULL, '2022-09-04 19:00:33.037074', '2022-09-04 19:00:33.037191', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (19, 10, 1, 2, 30000, 30000, NULL, '2022-09-04 19:04:12.37446', '2022-09-04 19:04:12.374599', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (21, 10, 3, 3, 90000, 90000, NULL, '2022-09-04 19:04:17.120167', '2022-09-04 19:04:17.120271', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (20, 10, 2, 4, 50000, 50000, NULL, '2022-09-04 19:04:14.769356', '2022-09-04 19:04:14.769502', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (23, 7, 1, 3, 30000, 30000, NULL, '2022-09-04 19:07:51.892772', '2022-09-04 19:07:51.892971', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (22, 7, 2, 3, 50000, 50000, NULL, '2022-09-04 19:07:35.15522', '2022-09-04 19:07:35.155321', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (26, 5, 2, 3, 50000, 50000, NULL, '2022-09-04 19:11:44.082466', '2022-09-04 19:11:44.082519', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (25, 5, 3, 3, 90000, 90000, NULL, '2022-09-04 19:11:42.134977', '2022-09-04 19:11:42.135129', NULL);
INSERT INTO public.transactions (id, user_id, subscription_id, status_id, gross_total, net_total, user_voucher_id, created_at, updated_at, deleted_at) VALUES (24, 5, 1, 3, 30000, 30000, NULL, '2022-09-04 19:11:39.65282', '2022-09-04 19:11:39.652881', NULL);


--
-- TOC entry 3605 (class 0 OID 18964)
-- Dependencies: 243
-- Data for Name: user_gifts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (1, 4, 1, '2022-09-02 00:25:36.643323', '2022-09-04 00:35:24.944414', NULL, 1);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (2, 4, 2, '2022-09-02 12:58:08.187244', '2022-09-04 00:35:24.957653', NULL, 1);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (3, 9, 1, '2022-09-04 19:01:53.70992', '2022-09-04 19:01:53.70992', NULL, 0);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (4, 10, 1, '2022-09-04 19:05:10.675384', '2022-09-04 19:05:10.675384', NULL, 0);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (5, 7, 1, '2022-09-04 19:08:12.542197', '2022-09-04 19:08:12.542197', NULL, 0);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (6, 5, 1, '2022-09-04 19:12:36.349855', '2022-09-04 19:13:43.286293', NULL, 1);
INSERT INTO public.user_gifts (id, user_id, gift_id, created_at, updated_at, deleted_at, is_claimed) VALUES (7, 5, 2, '2022-09-04 19:12:38.263912', '2022-09-04 19:13:43.298052', NULL, 1);


--
-- TOC entry 3586 (class 0 OID 18742)
-- Dependencies: 224
-- Data for Name: user_post_activities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (5, 1, 5, 3, 0, 0, NULL, NULL, '2022-08-30 13:38:55.561602', '2022-08-30 13:38:55.561602', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (40, 5, 3, 1, 1, 1, NULL, NULL, '2022-09-04 19:13:22.878486', '2022-09-04 19:13:24.003158', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (1, 3, 1, 2, 0, 0, NULL, NULL, '2022-08-30 03:27:48.027048', '2022-08-30 03:27:48.027048', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (15, 3, 5, 1, 1, 1, NULL, NULL, '2022-09-04 18:58:55.028116', '2022-09-04 18:58:57.13879', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (34, 7, 10, 1, 1, 1, NULL, NULL, '2022-09-04 19:09:15.448066', '2022-09-04 19:09:17.006382', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (2, 4, 8, 23, 0, 1, NULL, NULL, '2022-08-30 13:17:27.58799', '2022-09-01 23:03:23.523719', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (13, 1, 7, 1, 0, 0, NULL, NULL, '2022-09-01 16:41:33.435699', '2022-09-01 16:41:33.435699', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (23, 10, 14, 2, 1, 1, NULL, NULL, '2022-09-04 19:05:31.889169', '2022-09-04 19:05:36.153207', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (11, 1, 8, 6, 1, 1, NULL, NULL, '2022-08-31 16:32:52.312089', '2022-08-31 16:32:52.312089', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (27, 10, 8, 1, 1, 1, NULL, NULL, '2022-09-04 19:05:58.075521', '2022-09-04 19:05:59.734829', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (20, 9, 6, 1, 1, 1, NULL, NULL, '2022-09-04 19:02:27.585729', '2022-09-04 19:02:28.963916', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (16, 3, 9, 1, 1, 1, NULL, NULL, '2022-09-04 18:59:51.3803', '2022-09-04 18:59:53.445407', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (3, 4, 6, 10, 1, 0, NULL, NULL, '2022-08-30 13:29:10.731845', '2022-08-30 13:29:10.731845', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (31, 7, 14, 1, 1, 1, NULL, NULL, '2022-09-04 19:09:00.33865', '2022-09-04 19:09:01.858582', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (38, 5, 7, 1, 1, 1, NULL, NULL, '2022-09-04 19:13:12.943655', '2022-09-04 19:13:14.039105', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (35, 7, 6, 1, 1, 1, NULL, NULL, '2022-09-04 19:09:25.614906', '2022-09-04 19:09:26.903175', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (8, 4, 7, 36, 1, 1, NULL, NULL, '2022-08-30 22:16:20.190774', '2022-09-02 12:56:02.340538', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (21, 9, 3, 1, 1, 1, NULL, NULL, '2022-09-04 19:02:46.498056', '2022-09-04 19:02:47.601761', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (17, 3, 6, 1, 1, 1, NULL, NULL, '2022-09-04 19:00:00.203013', '2022-09-04 19:00:02.934532', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (10, 2, 7, 1, 0, 1, NULL, NULL, '2022-08-31 16:26:44.943661', '2022-08-31 16:26:44.943661', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (7, 1, 1, 4, 1, 1, NULL, NULL, '2022-08-30 13:50:23.042623', '2022-08-30 13:50:23.042623', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (24, 10, 11, 2, 1, 0, NULL, NULL, '2022-09-04 19:05:39.798434', '2022-09-04 19:05:44.18177', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (28, 10, 1, 1, 1, 1, NULL, NULL, '2022-09-04 19:06:13.452924', '2022-09-04 19:06:14.752174', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (4, 4, 5, 4, 0, 0, NULL, NULL, '2022-08-30 13:38:46.518869', '2022-09-02 12:58:43.600445', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (32, 7, 9, 1, 1, 0, NULL, NULL, '2022-09-04 19:09:05.596284', '2022-09-04 19:09:06.470149', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (9, 2, 8, 36, 1, 1, NULL, NULL, '2022-08-31 15:34:50.924801', '2022-08-31 15:34:50.924801', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (6, 4, 1, 11, 1, 1, NULL, NULL, '2022-08-30 13:48:07.835717', '2022-09-01 22:51:41.065917', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (18, 9, 5, 1, 1, 1, NULL, NULL, '2022-09-04 19:02:12.650951', '2022-09-04 19:02:14.285571', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (22, 10, 13, 1, 1, 1, NULL, NULL, '2022-09-04 19:05:28.41229', '2022-09-04 19:05:29.741588', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (14, 3, 11, 1, 1, 1, NULL, NULL, '2022-09-04 18:58:44.458566', '2022-09-04 18:58:46.083381', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (25, 10, 10, 1, 0, 1, NULL, NULL, '2022-09-04 19:05:48.887425', '2022-09-04 19:05:49.884971', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (29, 10, 2, 1, 1, 0, NULL, NULL, '2022-09-04 19:06:22.522172', '2022-09-04 19:06:23.494829', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (36, 5, 14, 1, 0, 1, NULL, NULL, '2022-09-04 19:13:00.997842', '2022-09-04 19:13:02.317914', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (12, 2, 4, 1, 1, 1, NULL, NULL, '2022-08-31 16:36:31.982885', '2022-08-31 16:36:31.982885', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (33, 7, 1, 1, 1, 0, NULL, NULL, '2022-09-04 19:09:10.284779', '2022-09-04 19:09:11.17771', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (19, 9, 12, 1, 1, 1, NULL, NULL, '2022-09-04 19:02:20.023729', '2022-09-04 19:02:21.453723', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (26, 10, 9, 1, 1, 0, NULL, NULL, '2022-09-04 19:05:53.852911', '2022-09-04 19:05:54.718755', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (39, 5, 6, 1, 1, 1, NULL, NULL, '2022-09-04 19:13:17.593154', '2022-09-04 19:13:18.752196', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (30, 10, 4, 1, 1, 0, NULL, NULL, '2022-09-04 19:06:30.30451', '2022-09-04 19:06:31.342686', NULL);
INSERT INTO public.user_post_activities (id, user_id, post_id, views_count, is_liked, is_shared, date_liked, date_shared, created_at, updated_at, deleted_at) VALUES (37, 5, 10, 1, 1, 0, NULL, NULL, '2022-09-04 19:13:08.030886', '2022-09-04 19:13:08.866117', NULL);


--
-- TOC entry 3579 (class 0 OID 18655)
-- Dependencies: 217
-- Data for Name: user_referrals; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (1, 5, 4, '2022-08-30 13:27:35.090781', '2022-08-30 13:27:35.090781', NULL);
INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (2, 6, 4, '2022-09-01 23:12:08.251957', '2022-09-01 23:12:08.251957', NULL);
INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (3, 7, 5, '2022-09-04 18:10:17.533533', '2022-09-04 18:10:17.533533', NULL);
INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (4, 8, 5, '2022-09-04 18:12:29.154565', '2022-09-04 18:12:29.154565', NULL);
INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (5, 9, 4, '2022-09-04 18:14:03.751538', '2022-09-04 18:14:03.751538', NULL);
INSERT INTO public.user_referrals (id, user_id, referrer_user_id, created_at, updated_at, deleted_at) VALUES (6, 10, 1, '2022-09-04 18:15:40.462184', '2022-09-04 18:15:40.462184', NULL);


--
-- TOC entry 3591 (class 0 OID 18794)
-- Dependencies: 229
-- Data for Name: user_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (1, 1, 1, '2022-08-01', '2022-08-31', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (2, 4, 2, '2022-08-30', '2022-09-30', '2022-08-30 13:28:45.255957', '2022-08-30 13:28:45.255957', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (3, 4, 1, '2022-10-01', '2022-11-01', '2022-08-30 19:14:04.161414', '2022-08-30 19:14:04.161414', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (4, 2, 2, '2022-08-31', '2022-10-01', '2022-08-31 16:26:01.481084', '2022-08-31 16:26:01.481084', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (5, 5, 3, '2022-08-31', '2022-10-01', '2022-08-31 23:03:40.044652', '2022-08-31 23:03:40.044652', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (6, 5, 2, '2022-10-02', '2022-11-02', '2022-08-31 23:03:43.805998', '2022-08-31 23:03:43.805998', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (7, 1, 1, '2022-09-01', '2022-10-01', '2022-09-01 00:28:50.788815', '2022-09-01 00:28:50.788815', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (8, 1, 2, '2022-10-02', '2022-11-02', '2022-09-01 16:17:14.423228', '2022-09-01 16:17:14.423228', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (9, 1, 3, '2022-11-03', '2022-12-03', '2022-09-01 16:18:57.576758', '2022-09-01 16:18:57.576758', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (10, 1, 2, '2022-12-04', '2023-01-04', '2022-09-01 16:43:54.973776', '2022-09-01 16:43:54.973776', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (11, 4, 2, '2022-11-02', '2022-12-02', '2022-09-02 00:24:25.878871', '2022-09-02 00:24:25.878871', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (12, 4, 1, '2022-12-03', '2023-01-03', '2022-09-02 00:25:36.640609', '2022-09-02 00:25:36.640609', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (13, 4, 3, '2023-01-04', '2023-02-04', '2022-09-02 12:58:08.184142', '2022-09-02 12:58:08.184142', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (14, 3, 2, '2022-09-04', '2022-10-04', '2022-09-04 18:59:31.353029', '2022-09-04 18:59:31.353029', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (15, 9, 3, '2022-09-04', '2022-10-04', '2022-09-04 19:01:53.707774', '2022-09-04 19:01:53.707774', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (16, 10, 3, '2022-09-04', '2022-10-04', '2022-09-04 19:05:10.672699', '2022-09-04 19:05:10.672699', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (17, 7, 1, '2022-09-04', '2022-10-04', '2022-09-04 19:08:08.903658', '2022-09-04 19:08:08.903658', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (18, 7, 2, '2022-10-05', '2022-11-05', '2022-09-04 19:08:12.540009', '2022-09-04 19:08:12.540009', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (19, 5, 2, '2022-11-03', '2022-12-03', '2022-09-04 19:12:33.981418', '2022-09-04 19:12:33.981418', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (20, 5, 3, '2022-12-04', '2023-01-04', '2022-09-04 19:12:36.348292', '2022-09-04 19:12:36.348292', NULL);
INSERT INTO public.user_subscriptions (id, user_id, subscription_id, date_start, date_ended, created_at, updated_at, deleted_at) VALUES (21, 5, 1, '2023-01-05', '2023-02-05', '2022-09-04 19:12:38.261156', '2022-09-04 19:12:38.261156', NULL);


--
-- TOC entry 3594 (class 0 OID 18830)
-- Dependencies: 232
-- Data for Name: user_vouchers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_vouchers (id, user_id, voucher_id, date_expired, is_used, created_at, updated_at, deleted_at) VALUES (3, 4, 1, '2022-10-01', 0, '2022-08-31 23:03:43.808344', '2022-08-31 23:03:43.808344', NULL);
INSERT INTO public.user_vouchers (id, user_id, voucher_id, date_expired, is_used, created_at, updated_at, deleted_at) VALUES (1, 1, 1, '2022-11-12', 1, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.user_vouchers (id, user_id, voucher_id, date_expired, is_used, created_at, updated_at, deleted_at) VALUES (2, 1, 2, '2022-12-12', 1, '2022-08-31 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.user_vouchers (id, user_id, voucher_id, date_expired, is_used, created_at, updated_at, deleted_at) VALUES (4, 4, 2, '2022-10-04', 0, '2022-09-04 19:12:36.350789', '2022-09-04 19:12:36.350789', NULL);


--
-- TOC entry 3577 (class 0 OID 18620)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (3, 1, 'raiden', 'raiden@mail.com', '$2a$10$IkK0RXxo68ntEAuuMLCRXOiEEucmvPtwdehHHm6eVMQ4ATKagD.kO', 'Raiden Shogun', '0199903', 3, 'XNDTBK', 1, 7, '2022-08-30 03:18:21.431388', '2022-08-30 03:18:21.431388', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (9, 2, 'lisa', 'lisa@mail.com', '$2a$10$4vE2vd/QEhl/g7gN7UfYbONmcTj.p/iWfnP30.Cbu1v1Y2hvVuegi', 'Lisa', '232386718', 9, 'RZXZUG', 1, 15, '2022-09-04 18:14:03.750834', '2022-09-04 18:14:03.750834', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (2, 1, 'zhongli', 'zhongli@mail.com', '$2a$10$Q1mNoIGtJefJzssxbUWzR.pqwcbABPMR04cIbs0rogsxjNfbwPbyq', 'Zhongli', '0199902', 2, 'ECVSXJ', 1, 7, '2022-08-30 03:14:30.36381', '2022-08-30 03:14:30.36381', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (10, 2, 'jean', 'jean@mail.com', '$2a$10$oLzgK3MwiZTf.6pPDPNmZuzCCu8xK4pgpLsbjDk6tvLITO5CjTFOm', 'Jean Gunnhildr', '4121827', 10, 'KQPXPC', 1, 15, '2022-09-04 18:15:40.450985', '2022-09-04 18:15:40.450985', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (1, 1, 'venti', 'venti@mail.com', '$2a$10$ZqZtA6rp4orAtgZ8tj3jsuHTtQ6fEMDLI8No046F/i2.mF4bycv5q', 'Venti', '0199901', 1, 'VENTI', 1, 43, '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (6, 2, 'razor', 'razor@mail.com', '$2a$10$mzwSEdTfMGlvdhzfMQSUGezYHdu30.VS58suhErb17QEwxASpAdNC', 'Razor', '0392038', 6, 'CMKDJC', 1, 0, '2022-09-01 23:12:08.249028', '2022-09-01 23:12:08.249028', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (7, 2, 'sucrose', 'sucrose@mail.com', '$2a$10$paQeDmxqQY5tscnCGFBtFOHZNXRGl9qy3adfA06i811XjUiANpT..', 'Sucrose', '01209237', 7, 'AWSVPN', 1, 11, '2022-09-04 18:10:17.532042', '2022-09-04 18:10:17.532042', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (4, 2, 'klee', 'klee@mail.com', '$2a$10$2PK01yupDfMtrct.X7z9x.v3EbPZEdJPRnNZPotymqi3PVmpnhFoe', 'Klee', '0188880192', 4, 'FEDYZH', 1, 30, '2022-08-30 13:04:00.092372', '2022-08-30 13:04:00.092372', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (8, 2, 'mona', 'mona@mail.com', '$2a$10$0Qp5yTS5AsGIJ4QB1fJCiuepnrcvzEhUtJ8nVvrQN9J2PQXxxM.Ui', 'Mona', '92328180', 8, 'QKPFAR', 1, 0, '2022-09-04 18:12:29.143648', '2022-09-04 18:12:29.143648', NULL);
INSERT INTO public.users (id, role_id, username, email, password, name, phone, address_id, referral_code, profile_pic_id, mora, created_at, updated_at, deleted_at) VALUES (5, 2, 'albedo', 'albedo@mail.com', '$2a$10$MzMA/rCojQlY0PYDiogL5eZRhckSFMi48xk8sAl2kb2viPC632VDS', 'Albedo', '012903883', 5, 'HYVFLR', 1, 58, '2022-08-30 13:27:35.08905', '2022-08-30 13:27:35.08905', NULL);


--
-- TOC entry 3592 (class 0 OID 18812)
-- Dependencies: 230
-- Data for Name: vouchers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.vouchers (id, name, description, image_id, amount, code, created_at, updated_at, deleted_at) VALUES (1, 'VOUCHER 25K', 'Voucher 25K untuk pembelian subscription.', 3, 25000, 'GLAZELILY', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.vouchers (id, name, description, image_id, amount, code, created_at, updated_at, deleted_at) VALUES (2, 'VOUCHER 50K', 'Voucher 50K untuk pembelian subscription.', 3, 50000, 'WINDWHEELASTER', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);
INSERT INTO public.vouchers (id, name, description, image_id, amount, code, created_at, updated_at, deleted_at) VALUES (3, 'VOUCHER 75K', 'Voucher 75K untuk pembelian subscription.', 3, 75000, 'NILOTPALALOTUS', '2022-08-30 02:31:39.662158', '2022-08-30 02:31:39.662158', NULL);


--
-- TOC entry 3625 (class 0 OID 0)
-- Dependencies: 212
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.addresses_id_seq', 10, true);


--
-- TOC entry 3626 (class 0 OID 0)
-- Dependencies: 240
-- Name: gift_claim_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gift_claim_items_id_seq', 22, true);


--
-- TOC entry 3627 (class 0 OID 0)
-- Dependencies: 238
-- Name: gift_claims_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gift_claims_id_seq', 13, true);


--
-- TOC entry 3628 (class 0 OID 0)
-- Dependencies: 209
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.images_id_seq', 5, true);


--
-- TOC entry 3629 (class 0 OID 0)
-- Dependencies: 219
-- Name: post_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_categories_id_seq', 4, true);


--
-- TOC entry 3630 (class 0 OID 0)
-- Dependencies: 225
-- Name: post_unlocks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_unlocks_id_seq', 29, true);


--
-- TOC entry 3631 (class 0 OID 0)
-- Dependencies: 221
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.posts_id_seq', 14, true);


--
-- TOC entry 3632 (class 0 OID 0)
-- Dependencies: 234
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 26, true);


--
-- TOC entry 3633 (class 0 OID 0)
-- Dependencies: 242
-- Name: user_gifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_gifts_id_seq', 7, true);


--
-- TOC entry 3634 (class 0 OID 0)
-- Dependencies: 223
-- Name: user_post_activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_post_activities_id_seq', 40, true);


--
-- TOC entry 3635 (class 0 OID 0)
-- Dependencies: 216
-- Name: user_referrals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_referrals_id_seq', 6, true);


--
-- TOC entry 3636 (class 0 OID 0)
-- Dependencies: 228
-- Name: user_subscriptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_subscriptions_id_seq', 21, true);


--
-- TOC entry 3637 (class 0 OID 0)
-- Dependencies: 231
-- Name: user_vouchers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_vouchers_id_seq', 4, true);


--
-- TOC entry 3638 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- TOC entry 3343 (class 2606 OID 18618)
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- TOC entry 3397 (class 2606 OID 18951)
-- Name: gift_claim_items gift_claim_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_items
    ADD CONSTRAINT gift_claim_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3391 (class 2606 OID 18917)
-- Name: gift_claim_statuses gift_claim_statuses_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_statuses
    ADD CONSTRAINT gift_claim_statuses_name_key UNIQUE (name);


--
-- TOC entry 3393 (class 2606 OID 18915)
-- Name: gift_claim_statuses gift_claim_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_statuses
    ADD CONSTRAINT gift_claim_statuses_pkey PRIMARY KEY (id);


--
-- TOC entry 3395 (class 2606 OID 18927)
-- Name: gift_claims gift_claims_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claims
    ADD CONSTRAINT gift_claims_pkey PRIMARY KEY (id);


--
-- TOC entry 3387 (class 2606 OID 18901)
-- Name: gifts gifts_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_name_key UNIQUE (name);


--
-- TOC entry 3389 (class 2606 OID 18899)
-- Name: gifts gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_pkey PRIMARY KEY (id);


--
-- TOC entry 3337 (class 2606 OID 18596)
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- TOC entry 3359 (class 2606 OID 18694)
-- Name: post_categories post_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_categories
    ADD CONSTRAINT post_categories_pkey PRIMARY KEY (id);


--
-- TOC entry 3355 (class 2606 OID 18683)
-- Name: post_tiers post_tiers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_tiers
    ADD CONSTRAINT post_tiers_name_key UNIQUE (name);


--
-- TOC entry 3357 (class 2606 OID 18681)
-- Name: post_tiers post_tiers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_tiers
    ADD CONSTRAINT post_tiers_pkey PRIMARY KEY (id);


--
-- TOC entry 3367 (class 2606 OID 18771)
-- Name: post_unlocks post_unlocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_unlocks
    ADD CONSTRAINT post_unlocks_pkey PRIMARY KEY (id);


--
-- TOC entry 3361 (class 2606 OID 18708)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- TOC entry 3363 (class 2606 OID 18710)
-- Name: posts posts_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_slug_key UNIQUE (slug);


--
-- TOC entry 3339 (class 2606 OID 18607)
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- TOC entry 3341 (class 2606 OID 18605)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 3369 (class 2606 OID 18792)
-- Name: subscriptions subscriptions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_name_key UNIQUE (name);


--
-- TOC entry 3371 (class 2606 OID 18790)
-- Name: subscriptions subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);


--
-- TOC entry 3381 (class 2606 OID 18859)
-- Name: transaction_statuses transaction_statuses_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_statuses
    ADD CONSTRAINT transaction_statuses_name_key UNIQUE (name);


--
-- TOC entry 3383 (class 2606 OID 18857)
-- Name: transaction_statuses transaction_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_statuses
    ADD CONSTRAINT transaction_statuses_pkey PRIMARY KEY (id);


--
-- TOC entry 3385 (class 2606 OID 18868)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3399 (class 2606 OID 18971)
-- Name: user_gifts user_gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT user_gifts_pkey PRIMARY KEY (id);


--
-- TOC entry 3365 (class 2606 OID 18752)
-- Name: user_post_activities user_post_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_post_activities
    ADD CONSTRAINT user_post_activities_pkey PRIMARY KEY (id);


--
-- TOC entry 3353 (class 2606 OID 18662)
-- Name: user_referrals user_referrals_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_referrals
    ADD CONSTRAINT user_referrals_pkey PRIMARY KEY (id);


--
-- TOC entry 3373 (class 2606 OID 18801)
-- Name: user_subscriptions user_subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_subscriptions
    ADD CONSTRAINT user_subscriptions_pkey PRIMARY KEY (id);


--
-- TOC entry 3379 (class 2606 OID 18838)
-- Name: user_vouchers user_vouchers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT user_vouchers_pkey PRIMARY KEY (id);


--
-- TOC entry 3345 (class 2606 OID 18636)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3347 (class 2606 OID 18632)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3349 (class 2606 OID 18638)
-- Name: users users_referral_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_referral_code_key UNIQUE (referral_code);


--
-- TOC entry 3351 (class 2606 OID 18634)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 3375 (class 2606 OID 18823)
-- Name: vouchers vouchers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT vouchers_name_key UNIQUE (name);


--
-- TOC entry 3377 (class 2606 OID 18821)
-- Name: vouchers vouchers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT vouchers_pkey PRIMARY KEY (id);


--
-- TOC entry 3401 (class 2606 OID 18644)
-- Name: users fk_address; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES public.addresses(id);


--
-- TOC entry 3426 (class 2606 OID 18933)
-- Name: gift_claims fk_address; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claims
    ADD CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES public.addresses(id);


--
-- TOC entry 3409 (class 2606 OID 18731)
-- Name: posts fk_created_by; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id);


--
-- TOC entry 3428 (class 2606 OID 18952)
-- Name: gift_claim_items fk_gift; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_items
    ADD CONSTRAINT fk_gift FOREIGN KEY (gift_id) REFERENCES public.gifts(id);


--
-- TOC entry 3431 (class 2606 OID 18977)
-- Name: user_gifts fk_gift; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT fk_gift FOREIGN KEY (gift_id) REFERENCES public.gifts(id);


--
-- TOC entry 3429 (class 2606 OID 18957)
-- Name: gift_claim_items fk_gift_claim; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claim_items
    ADD CONSTRAINT fk_gift_claim FOREIGN KEY (gift_claim_id) REFERENCES public.gift_claims(id);


--
-- TOC entry 3427 (class 2606 OID 18938)
-- Name: gift_claims fk_gift_claim_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claims
    ADD CONSTRAINT fk_gift_claim_status FOREIGN KEY (status_id) REFERENCES public.gift_claim_statuses(id);


--
-- TOC entry 3417 (class 2606 OID 18824)
-- Name: vouchers fk_image; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT fk_image FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- TOC entry 3424 (class 2606 OID 18902)
-- Name: gifts fk_image; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT fk_image FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- TOC entry 3408 (class 2606 OID 18726)
-- Name: posts fk_img_content; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_img_content FOREIGN KEY (img_content_id) REFERENCES public.images(id);


--
-- TOC entry 3407 (class 2606 OID 18721)
-- Name: posts fk_img_thumbnail; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_img_thumbnail FOREIGN KEY (img_thumbnail_id) REFERENCES public.images(id);


--
-- TOC entry 3412 (class 2606 OID 18758)
-- Name: user_post_activities fk_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_post_activities
    ADD CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- TOC entry 3414 (class 2606 OID 18777)
-- Name: post_unlocks fk_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_unlocks
    ADD CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- TOC entry 3406 (class 2606 OID 18716)
-- Name: posts fk_post_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_post_category FOREIGN KEY (post_category_id) REFERENCES public.post_categories(id);


--
-- TOC entry 3405 (class 2606 OID 18711)
-- Name: posts fk_post_tier; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_post_tier FOREIGN KEY (post_tier_id) REFERENCES public.post_tiers(id);


--
-- TOC entry 3402 (class 2606 OID 18649)
-- Name: users fk_profile_pic; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_profile_pic FOREIGN KEY (profile_pic_id) REFERENCES public.images(id);


--
-- TOC entry 3404 (class 2606 OID 18668)
-- Name: user_referrals fk_referrer_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_referrals
    ADD CONSTRAINT fk_referrer_user FOREIGN KEY (referrer_user_id) REFERENCES public.users(id);


--
-- TOC entry 3400 (class 2606 OID 18639)
-- Name: users fk_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 3416 (class 2606 OID 18807)
-- Name: user_subscriptions fk_subscription; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_subscriptions
    ADD CONSTRAINT fk_subscription FOREIGN KEY (subscription_id) REFERENCES public.subscriptions(id);


--
-- TOC entry 3421 (class 2606 OID 18874)
-- Name: transactions fk_subscription; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_subscription FOREIGN KEY (subscription_id) REFERENCES public.subscriptions(id);


--
-- TOC entry 3422 (class 2606 OID 18879)
-- Name: transactions fk_transaction_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transaction_status FOREIGN KEY (status_id) REFERENCES public.transaction_statuses(id);


--
-- TOC entry 3410 (class 2606 OID 18736)
-- Name: posts fk_updated_by; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_updated_by FOREIGN KEY (updated_by_id) REFERENCES public.users(id);


--
-- TOC entry 3403 (class 2606 OID 18663)
-- Name: user_referrals fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_referrals
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3411 (class 2606 OID 18753)
-- Name: user_post_activities fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_post_activities
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3413 (class 2606 OID 18772)
-- Name: post_unlocks fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_unlocks
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3415 (class 2606 OID 18802)
-- Name: user_subscriptions fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_subscriptions
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3418 (class 2606 OID 18839)
-- Name: user_vouchers fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3420 (class 2606 OID 18869)
-- Name: transactions fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3425 (class 2606 OID 18928)
-- Name: gift_claims fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gift_claims
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3430 (class 2606 OID 18972)
-- Name: user_gifts fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3423 (class 2606 OID 18884)
-- Name: transactions fk_user_voucher; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_user_voucher FOREIGN KEY (user_voucher_id) REFERENCES public.user_vouchers(id);


--
-- TOC entry 3419 (class 2606 OID 18844)
-- Name: user_vouchers fk_voucher; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT fk_voucher FOREIGN KEY (voucher_id) REFERENCES public.vouchers(id);


-- Completed on 2022-09-04 19:18:57 WIB

--
-- PostgreSQL database dump complete
--

