BEGIN;

CREATE TABLE IF NOT EXISTS indicadores.variavel_permissao(
    id SERIAL PRIMARY KEY NOT NULL,
    usuario_id INTEGER NOT NULL REFERENCES indicadores.usuario(id),
    variavel_id INTEGER NOT NULL REFERENCES indicadores.variavel(id),
    visualizar BOOLEAN NOT NULL,
    enviar BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS indicadores.resposta(
    id SERIAL PRIMARY KEY NOT NULL,
    variavel_id INTEGER NOT NULL REFERENCES indicadores.variavel(id),
    usuario_id INTEGER NOT NULL REFERENCES indicadores.usuario(id),
    agendamento_id INTEGER NOT NULL REFERENCES indicadores.agendamento(id)    
);
CREATE TABLE IF NOT EXISTS indicadores.resposta_numerica(
    id SERIAL PRIMARY KEY NOT NULL,
    valor FLOAT NOT NULL,
    ativo BOOLEAN DEFAULT true
);

CREATE TABLE IF NOT EXISTS indicadores.versao_resposta_numerica(
    id SERIAL PRIMARY KEY NOT NULL,
    resposta_id INTEGER NOT NULL REFERENCES indicadores.resposta(id),
    resposta_numerica_id INTEGER NOT NULL REFERENCES indicadores.resposta_numerica(id),
    versao SMALLINT NOT NULL,
    criado_em TIMESTAMP DEFAULT now() NOT NULL
);



COMMIT;
