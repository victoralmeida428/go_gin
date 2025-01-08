BEGIN;

CREATE TABLE IF NOT EXISTS indicadores.metodo (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS indicadores.nivel(
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS indicadores.empresa(
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL UNIQUE,
    nivel_id INTEGER REFERENCES indicadores.nivel(id)
);

ALTER TABLE indicadores.agendamento DROP COLUMN unificado;

ALTER TABLE indicadores.agendamento ADD COLUMN metodo_id INTEGER REFERENCES indicadores.metodo(id);

ALTER TABLE indicadores.usuario  DROP COLUMN empresa;

ALTER TABLE indicadores.usuario ADD COLUMN empresa_id INTEGER REFERENCES indicadores.empresa(id);

ALTER TABLE indicadores.agendamento ADD COLUMN empresa_id INTEGER REFERENCES indicadores.empresa(id);
ALTER TABLE indicadores.formulario ADD COLUMN nivel_id INTEGER REFERENCES indicadores.nivel(id);


COMMIT;