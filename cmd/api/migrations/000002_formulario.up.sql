BEGIN;
CREATE TABLE IF NOT EXISTS indicadores.formulario
(
    id        SERIAL PRIMARY KEY,
    nome      CHARACTER VARYING NOT NULL unique,
    descricao CHARACTER VARYING,
    ativo     BOOL
);

CREATE TABLE IF NOT EXISTS indicadores.versao_formulario
(
    id            SERIAL PRIMARY KEY,
    formulario_id INTEGER REFERENCES indicadores.formulario (id) not null,
    versao        SMALLINT                                       not null,
    criado_em     timestamp DEFAULT now()                        not null
);

create table if not exists indicadores.periodicidade
(
    id        serial primary key,
    nome      character varying not null,
    intervalo interval          not null
);

create table if not exists indicadores.agendamento
(
    id                   serial primary key,
    versao_formulario_id integer references indicadores.versao_formulario (id) not null,
    usuario_id              integer references indicadores.usuario (id)           not null,
    periodicidade_id     integer references indicadores.periodicidade (id)     not null,
    inicio               timestamp                                             not null,
    unificado            bool                                                  not null,
    ativo                bool                                                  not null
);

CREATE INDEX agendamento_user_index
    ON indicadores.agendamento USING btree
    (usuario_id);

CREATE INDEX agendamento_form_index
    ON indicadores.agendamento USING btree
    (versao_formulario_id);

create table if not exists indicadores.tipo_variavel(
    id serial primary key,
    nome character varying  not null
);

create table if not exists indicadores.grupo(
    id serial primary key,
    nome character varying not null UNIQUE
);

create table if not exists indicadores.variavel(
    id serial primary key,
    tipo_variavel_id integer references indicadores.tipo_variavel(id),
    pergunta_id integer references indicadores.variavel(id),
    grupo_id integer REFERENCES indicadores.grupo(id),
    possui_item bool not null,
    obrigatorio bool not null,
    texto character varying not null
);

create table if not exists indicadores.variavel_formulario(
    id serial primary key,
    versao_formulario_id integer references indicadores.versao_formulario(id) not null,
    variavel_id integer references indicadores.variavel(id) not null
);

COMMIT;