begin;

CREATE TABLE IF NOT EXISTS indicadores.grupamento_usuario (
    id SERIAL PRIMARY KEY,
    grupamento_id INTEGER REFERENCES indicadores.grupamento(id) NOT NULL,
    usuario_id INTEGER REFERENCES indicadores.usuario(id) NOT NULL
);

ALTER TABLE IF EXISTS indicadores.grupamento_usuario
    ADD UNIQUE (grupamento_id, usuario_id);


COMMIT;