BEGIN;

DROP TABLE IF EXISTS indicadores.versao_resposta_numerica CASCADE;
-- Remover a tabela indicadores.resposta_numerica, pois depende de indicadores.resposta

DROP TABLE IF EXISTS indicadores.resposta_numerica CASCADE;

-- Remover a tabela indicadores.resposta
DROP TABLE IF EXISTS indicadores.resposta CASCADE;

-- Remover a tabela indicadores.variavel_permissao
DROP TABLE IF EXISTS indicadores.variavel_permissao CASCADE;


COMMIT;
