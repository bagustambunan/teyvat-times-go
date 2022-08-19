-- table images
CREATE TABLE public.images (
    id bigserial NOT NULL,
    url text NOT NULL,
    alt_text character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.images (url, alt_text) VALUES ('default-pic.png','Default profile pic');
INSERT INTO public.images (url, alt_text) VALUES ('default-thumbnail.png','Default thumbnail');
INSERT INTO public.images (url, alt_text) VALUES ('default-img.png','Default image');

-- table roles
CREATE TABLE public.roles (
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.roles (id, name) VALUES (1,'admin');
INSERT INTO public.roles (id, name) VALUES (2,'user');

-- table addresses
CREATE TABLE public.addresses (
    id bigserial NOT NULL,
    street character varying NOT NULL,
    city character varying NOT NULL,
    state character varying NOT NULL,
    country character varying NOT NULL,
    postal_code character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.addresses (street, city, state, country, postal_code) VALUES ('Dandelion St.127','Windrise','Mondstadt','Teyvat','MOND38');

-- table users
CREATE TABLE public.users (
    id bigserial NOT NULL,
    role_id int DEFAULT 2 NOT NULL,
    username character varying,
    email character varying NOT NULL UNIQUE,
    password character varying NOT NULL,
    name character varying NOT NULL,
    phone character varying NOT NULL,
    address_id bigint,
    referral_code character varying,
    profile_pic_id bigint DEFAULT 1 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_role
        FOREIGN KEY(role_id)
            REFERENCES public.roles(id),
    CONSTRAINT fk_address
        FOREIGN KEY(address_id)
            REFERENCES public.addresses(id),
    CONSTRAINT fk_profile_pic
        FOREIGN KEY(profile_pic_id)
            REFERENCES public.images(id)
);
INSERT INTO public.users (role_id, username, email, password, name, phone, address_id) VALUES (1,'jean','jean@mail.com','$2a$10$nbuYAZnJXsMjFIP4WkdqQ.BazePvRTc4705hq9CP1T1oRwcCQa2wm','Jean Gunnhildr','0199901',1);

-- table user_referrals
CREATE TABLE public.user_referrals (
    id bigserial NOT NULL,
    user_id bigint NOT NULL,
    referrer_user_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES public.users(id),
    CONSTRAINT fk_referrer_user
        FOREIGN KEY(referrer_user_id)
            REFERENCES public.users(id)
);

-- table tr_tokens
CREATE TABLE public.tr_tokens (
      id bigserial NOT NULL,
      user_id bigint NOT NULL,
      token character varying NOT NULL,
      is_expired int DEFAULT 0 NOT NULL,
      created_at timestamp without time zone DEFAULT now() NOT NULL,
      updated_at timestamp without time zone DEFAULT now() NOT NULL,
      deleted_at timestamp without time zone,
      PRIMARY KEY(id),
      CONSTRAINT fk_user
          FOREIGN KEY(user_id)
              REFERENCES public.users(id)
);

-- table post_tiers
CREATE TABLE public.post_tiers (
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    coins_required int NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.post_tiers (id, name, coins_required) VALUES (1,'Free',0);
INSERT INTO public.post_tiers (id, name, coins_required) VALUES (2,'Premium',1);
INSERT INTO public.post_tiers (id, name, coins_required) VALUES (3,'VIP',2);

-- table post_categories
CREATE TABLE public.post_categories(
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.post_categories (id, name) VALUES (1,'Politic');
INSERT INTO public.post_categories (id, name) VALUES (2,'Economy');
INSERT INTO public.post_categories (id, name) VALUES (3,'Sport');
INSERT INTO public.post_categories (id, name) VALUES (4,'Entertainment');

-- table posts
CREATE TABLE public.posts (
    id bigserial NOT NULL,
    post_tier_id int DEFAULT 1 NOT NULL,
    post_category_id int NOT NULL,
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
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_post_tier
        FOREIGN KEY(post_tier_id)
            REFERENCES public.post_tiers(id),
    CONSTRAINT fk_post_category
        FOREIGN KEY(post_category_id)
            REFERENCES public.post_categories(id),
    CONSTRAINT fk_img_thumbnail
        FOREIGN KEY(img_thumbnail_id)
            REFERENCES public.images(id),
    CONSTRAINT fk_img_content
        FOREIGN KEY(img_content_id)
            REFERENCES public.images(id),
    CONSTRAINT fk_created_by
        FOREIGN KEY(created_by_id)
            REFERENCES public.users(id),
    CONSTRAINT fk_updated_by
        FOREIGN KEY(updated_by_id)
            REFERENCES public.users(id)
);
INSERT INTO public.posts (post_tier_id, post_category_id, title, content, slug, summary, img_thumbnail_id, img_content_id, created_by_id, updated_by_id) VALUES (1,1,'Mondstadt signs a deal with Fatui','Test content','mondstadt-signs-a-deal-with-fatui','What do you think, Jean?',2,3,1,1);



------

-- table subscriptions
CREATE TABLE public.subscriptions(
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    price int NOT NULL,
    coins_amount int NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.subscriptions (id, name, price, coins_amount) VALUES (1,'Standard',30000,5);
INSERT INTO public.subscriptions (id, name, price, coins_amount) VALUES (2,'Premium',50000,10);
INSERT INTO public.subscriptions (id, name, price, coins_amount) VALUES (3,'Gold',90000,20);

-- table user_subscriptions
CREATE TABLE public.user_subscriptions (
    id bigserial NOT NULL,
    user_id bigint NOT NULL,
    subscription_id int NOT NULL,
    date_start date NOT NULL,
    date_ended date NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
       FOREIGN KEY(user_id)
           REFERENCES public.users(id),
    CONSTRAINT fk_subscription
       FOREIGN KEY(subscription_id)
           REFERENCES public.subscriptions(id)
);
INSERT INTO public.user_subscriptions (user_id, subscription_id, date_start, date_ended) VALUES (1, 1, '2022-08-01', '2022-08-31');

-- table vouchers
CREATE TABLE public.vouchers (
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    description text NOT NULL,
    image_id bigint DEFAULT 3 NOT NULL,
    amount int NOT NULL,
    code varchar NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_image
        FOREIGN KEY(image_id)
            REFERENCES public.images(id)
);
INSERT INTO public.vouchers (id, name, description, image_id, amount, code) VALUES (1,'VOUCHER 25K','Voucher 25K untuk pembelian subscription.',3,25000,'NEMESIS25');
INSERT INTO public.vouchers (id, name, description, image_id, amount, code) VALUES (2,'VOUCHER 50K','Voucher 50K untuk pembelian subscription.',3,50000,'NEMESIS50');
INSERT INTO public.vouchers (id, name, description, image_id, amount, code) VALUES (3,'VOUCHER 75K','Voucher 75K untuk pembelian subscription.',3,75000,'NEMESIS75');

-- table user_vouchers
CREATE TABLE public.user_vouchers (
     id bigserial NOT NULL,
     user_id bigint NOT NULL,
     voucher_id int NOT NULL,
     date_expired date NOT NULL,
     is_used int DEFAULT 0 NOT NULL,
     created_at timestamp without time zone DEFAULT now() NOT NULL,
     updated_at timestamp without time zone DEFAULT now() NOT NULL,
     deleted_at timestamp without time zone,
     PRIMARY KEY(id),
     CONSTRAINT fk_user
         FOREIGN KEY (user_id)
             REFERENCES public.users(id),
     CONSTRAINT fk_voucher
         FOREIGN KEY (voucher_id)
             REFERENCES public.vouchers(id)
);
INSERT INTO public.user_vouchers (user_id, voucher_id, date_expired) VALUES (1,1,'2022-11-12');
INSERT INTO public.user_vouchers (user_id, voucher_id, date_expired) VALUES (1,2,'2022-12-12');

-- table transaction_statuses
CREATE TABLE public.transaction_statuses (
    id int NOT NULL,
    name character varying NOT NULL UNIQUE,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id)
);
INSERT INTO public.transaction_statuses (id, name) VALUES (1,'Draft');
INSERT INTO public.transaction_statuses (id, name) VALUES (2,'Waiting for payment');
INSERT INTO public.transaction_statuses (id, name) VALUES (3,'Completed');
INSERT INTO public.transaction_statuses (id, name) VALUES (4,'Canceled');

-- table transactions
CREATE TABLE public.transactions (
    id bigserial NOT NULL,
    user_id bigint NOT NULL,
    status_id int NOT NULL,
    grossTotal int NOT NULL,
    netTotal int NOT NULL,
    user_voucher_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
           REFERENCES public.users(id),
    CONSTRAINT fk_transaction_status
        FOREIGN KEY (status_id)
           REFERENCES public.transaction_statuses(id),
    CONSTRAINT fk_user_voucher
        FOREIGN KEY (user_voucher_id)
            REFERENCES public.user_vouchers(id)
);
INSERT INTO public.transactions (user_id, status_id, grosstotal, netTotal, user_voucher_id) VALUES (1,2,30000,5000,1);

-- table transaction_items
CREATE TABLE public.transaction_items (
     id bigserial NOT NULL,
     subscription_id int NOT NULL,
     transaction_id bigint NOT NULL,
     created_at timestamp without time zone DEFAULT now() NOT NULL,
     updated_at timestamp without time zone DEFAULT now() NOT NULL,
     deleted_at timestamp without time zone,
     PRIMARY KEY(id),
     CONSTRAINT fk_subscription
         FOREIGN KEY (subscription_id)
             REFERENCES public.subscriptions(id),
     CONSTRAINT fk_transaction
         FOREIGN KEY (transaction_id)
             REFERENCES public.transactions(id)
);
INSERT INTO public.transaction_items (subscription_id, transaction_id) VALUES (1,1);

-- ---------------------------------------------------
-- -- table wallets
-- CREATE TABLE public.wallets (
--     id bigserial NOT NULL,
--     balance integer DEFAULT 0 NOT NULL,
--     created_at timestamp without time zone DEFAULT now() NOT NULL,
--     updated_at timestamp without time zone DEFAULT now() NOT NULL,
--     deleted_at timestamp without time zone,
--     PRIMARY KEY(id)
-- );
--
-- -- table transaction_types
-- CREATE TABLE public.transaction_types (
--                   id int NOT NULL,
--                   name character varying NOT NULL UNIQUE,
--                   created_at timestamp without time zone DEFAULT now() NOT NULL,
--                   updated_at timestamp without time zone DEFAULT now() NOT NULL,
--                   deleted_at timestamp without time zone,
--                   PRIMARY KEY(id)
-- );
--
-- -- table top_up_sources
-- CREATE TABLE public.top_up_sources (
--     id int NOT NULL,
--     name character varying NOT NULL UNIQUE,
--     created_at timestamp without time zone DEFAULT now() NOT NULL,
--     updated_at timestamp without time zone DEFAULT now() NOT NULL,
--     deleted_at timestamp without time zone,
--     PRIMARY KEY(id)
-- );
--
-- -- table transactions
-- CREATE TABLE public.transactions (
--     id bigserial NOT NULL,
--     transaction_type_id int NOT NULL,
--     top_up_source_id int NOT NULL,
--     source_wallet_id bigint NOT NULL,
--     destination_wallet_id bigint NOT NULL,
--     amount int NOT NULL,
--     description character varying NOT NULL,
--     created_at timestamp without time zone DEFAULT now() NOT NULL,
--     updated_at timestamp without time zone DEFAULT now() NOT NULL,
--     deleted_at timestamp without time zone,
--     PRIMARY KEY(id),
--     CONSTRAINT fk_transaction_type
--      FOREIGN KEY(transaction_type_id)
--          REFERENCES transaction_types(id),
--     CONSTRAINT fk_top_up_source
--      FOREIGN KEY(top_up_source_id)
--          REFERENCES top_up_sources(id),
--     CONSTRAINT fk_source_wallet
--      FOREIGN KEY(source_wallet_id)
--          REFERENCES wallets(id),
--     CONSTRAINT fk_destination_wallet
--      FOREIGN KEY(destination_wallet_id)
--          REFERENCES wallets(id)
-- );