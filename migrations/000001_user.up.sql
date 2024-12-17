begin;
create schema if not exists indicadores;

create table if not exists indicadores.usuario(
    id serial primary key,
    usuario character varying not null unique,
    email character varying,
    empresa character varying,
    manager_id integer references indicadores.usuario(id),
    senha character varying not null
);
commit;