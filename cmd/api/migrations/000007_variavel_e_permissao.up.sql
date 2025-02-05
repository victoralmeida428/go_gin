BEGIN;

CREATE TABLE IF NOT EXISTS indicadores.resposta(
    id SERIAL PRIMARY KEY NOT NULL,
    variavel_id INTEGER NOT NULL REFERENCES indicadores.variavel(id),
    usuario_id INTEGER NOT NULL REFERENCES indicadores.usuario(id),
    agendamento_id INTEGER NOT NULL REFERENCES indicadores.agendamento(id),
    criado_em TIMESTAMP DEFAULT now() NOT NULL,
    resposta_numerica FLOAT
);

ALTER TABLE IF EXISTS indicadores.resposta
    ADD CONSTRAINT resposta_variavel_agendamento_unique UNIQUE (variavel_id, agendamento_id);


CREATE TABLE IF NOT EXISTS indicadores.log_resposta_numerica(
   id SERIAL PRIMARY KEY NOT NULL,
    variavel_id INTEGER,
    usuario_id INTEGER,
    agendamento_id INTEGER,
    resposta_numerica FLOAT NOT NULL,
    criado_em TIMESTAMP DEFAULT now() NOT NULL
);



COMMIT;
