CREATE TABLE users (
                       id uuid PRIMARY KEY,
                       username text UNIQUE NOT NULL,
                       email text UNIQUE NOT NULL,
                       password text NOT NULL,
                       verify_email boolean DEFAULT false,
                       deleted boolean DEFAULT false,
                       created_at timestamp DEFAULT now(),
                       updated_at timestamp DEFAULT now()

);

CREATE TABLE short_link (
                            id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                            user_id uuid REFERENCES users(id),
                            url text NOT NULL,
                            token text NOT NULL,
                            s_key bigint,
                            deleted boolean DEFAULT false,
                            created_at timestamp DEFAULT now(),
                            updated_at timestamp DEFAULT now()
);

CREATE TABLE custom_link (
                             id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                             user_id uuid REFERENCES users(id),
                             url text NOT NULL,
                             token text UNIQUE NOT NULL,
                             is_suggestion boolean DEFAULT false,
                             suggestion_id bigint,
                             deleted boolean DEFAULT false,
                             created_at timestamp DEFAULT now(),
                             updated_at timestamp DEFAULT now()
);

CREATE TABLE recycle_link (
                              id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                              s_key bigint NOT NULL
);

CREATE TABLE aux_s_key (
                           N bigint,
                           "End" bigint,
                           Step bigint,
                           A0 bigint,
                           N0 bigint
);

CREATE INDEX ON users (id);
CREATE INDEX ON users (email);
CREATE INDEX ON short_link (s_key);
CREATE INDEX ON custom_link (token);
CREATE INDEX ON recycle_link (id);