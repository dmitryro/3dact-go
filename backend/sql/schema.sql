DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'postgres') THEN
      RAISE NOTICE 'Database already exists';  -- optional
   ELSE
      PERFORM dblink_exec('dbname=' || current_database()  -- current db
                        , 'CREATE DATABASE postgres');
   END IF;
END
$do$;

ALTER ROLE postgres SET client_encoding TO 'utf8';
ALTER ROLE postgres SET default_transaction_isolation TO 'read committed';
ALTER ROLE postgres SET timezone TO 'America/New_York';
GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;

CREATE SEQUENCE config_id_seq;
 
CREATE TABLE config(
    id integer NOT NULL DEFAULT nextval('config_id_seq'),
    email varchar(550),
    session_id varchar(550),
    is_automatic BOOLEAN DEFAULT FALSE, 
    date_updated timestamptz DEFAULT NOW()    
);

CREATE SEQUENCE attitude_id_seq;

CREATE TABLE attitudes(
        id integer NOT NULL DEFAULT nextval('attitude_id_seq'),
        name varchar(75) NOT NULL,
        code varchar(10) NOT NULL,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz,
        UNIQUE(id)
    );

CREATE SEQUENCE relationship_id_seq;

CREATE TABLE relationships(
        id integer NOT NULL DEFAULT nextval('relationship_id_seq'),
        name varchar(75) NOT NULL,
        code varchar(10) NOT NULL,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz,
        UNIQUE(id)
    );

CREATE SEQUENCE post_id_seq;

CREATE TABLE posts (
        id integer  NOT NULL DEFAULT nextval('post_id_seq'),
        title varchar(300) NOT NULL,
        body varchar(5000) NOT NULL,
        time_published timestamptz DEFAULT NOW(),
        time_last_edited timestamptz DEFAULT NOW(),
        time_last_commented timestamptz,
        attitude_id integer,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz, 
        FOREIGN KEY(attitude_id) REFERENCES attitudes(id) ON DELETE CASCADE,
        UNIQUE(id)
   );

CREATE SEQUENCE comment_id_seq;

CREATE TABLE comments (
        id integer  NOT NULL DEFAULT nextval('comment_id_seq'),
        text varchar(5000) NOT NULL,
        score integer,
        parent_id int REFERENCES comments(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
        post_id int NOT NULL  REFERENCES posts(id) ON UPDATE CASCADE ON DELETE CASCADE,
        attitude_id integer,
        FOREIGN KEY(attitude_id) REFERENCES attitudes(id) ON DELETE CASCADE,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz,
        hide boolean not null default false,
        UNIQUE(id)
   );

CREATE SEQUENCE item_id_seq;

CREATE TABLE items (
        id integer NOT NULL DEFAULT nextval('item_id_seq'),
        title varchar(75) NOT NULL,
        price float,
        category varchar(100)        
    );

CREATE SEQUENCE address_id_seq;

CREATE TABLE addresses (
        id integer NOT NULL DEFAULT nextval('address_id_seq'),
        address1 varchar(200) NOT NULL,
        address2 varchar(200),
        city varchar(200) NOT NULL,
        state varchar(100),
        postalcode varchar(100) NOT NULL,
        country varchar(200) NOT NULL
    );

CREATE SEQUENCE user_id_seq;

CREATE TABLE users (
        id integer NOT NULL DEFAULT nextval('user_id_seq'),
        username varchar(30) NOT NULL UNIQUE,
        category varchar(30),
        first_name varchar(30) NOT NULL,
        last_name varchar(30) NOT NULL,
        email varchar(75) NOT NULL,
        bio varchar(2048),
        password varchar(128) NOT NULL,
        in_chat BOOLEAN DEFAULT FALSE, 
        is_staff BOOLEAN NOT NULL,
        is_admin BOOLEAN DEFAULT FALSE,
        is_cleared BOOLEAN DEFAULT FALSE,
        is_suspended BOOLEAN DEFAULT FALSE,
        is_active BOOLEAN NOT NULL,
        is_superuser BOOLEAN NOT NULL,
        last_login timestamptz,
        date_joined date NOT NULL,
        password_recovery_key varchar(500),
        activation_key varchar(500),
        profile_image varchar(1000),
        avatar_image varchar(1000),
        avatar_thumbnail varchar(1000),
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz,
        UNIQUE(id)
    );

CREATE SEQUENCE peer_id_seq;

CREATE TABLE peers (
        id integer NOT NULL DEFAULT nextval('peer_id_seq'),
        relation_id integer NOT NULL,
        FOREIGN KEY(relation_id) REFERENCES relationships(id) ON DELETE CASCADE,
        initiator_id integer NOT NULL,
        FOREIGN KEY(initiator_id) REFERENCES users(id) ON DELETE CASCADE,
        acceptor_id integer NOT NULL,
        FOREIGN KEY(acceptor_id) REFERENCES users(id) ON DELETE CASCADE,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW(),
        deleted_at timestamptz,
        UNIQUE(id) 
);

CREATE SEQUENCE contact_id_seq;

CREATE TABLE contacts (
        id integer NOT NULL DEFAULT nextval('contact_id_seq'),
        username varchar(30) NOT NULL UNIQUE,
        first_name varchar(30) NOT NULL,
        last_name varchar(30) NOT NULL,
        email varchar(75) NOT NULL,
        phone varchar(30) NOT NULL
    );

CREATE SEQUENCE state_id_seq;

CREATE TABLE states (
        id integer NOT NULL DEFAULT nextval('state_id_seq'),
        name varchar(256) NOT NULL UNIQUE,
        code varchar(10) NOT NULL UNIQUE
    );


ALTER SEQUENCE attitude_id_seq
OWNED BY attitudes.id;

ALTER SEQUENCE post_id_seq
OWNED BY posts.id;

ALTER SEQUENCE user_id_seq
OWNED BY users.id;

ALTER SEQUENCE state_id_seq
OWNED BY states.id;

ALTER SEQUENCE contact_id_seq
OWNED BY contacts.id;

ALTER SEQUENCE config_id_seq
OWNED BY config.id;

ALTER SEQUENCE item_id_seq
OWNED BY items.id;

ALTER SEQUENCE address_id_seq
OWNED BY addresses.id;

ALTER SEQUENCE comment_id_seq
OWNED BY comments.id;

ALTER SEQUENCE relationship_id_seq
OWNED BY relationships.id;

ALTER SEQUENCE peer_id_seq
OWNED BY peers.id;
