BEGIN;

ALTER TABLE IF EXISTS indicadores.log_respostas_numerica DROP CONSTRAINT IF EXISTS log_respostas_numerica_agendamento_id_fkey;
ALTER TABLE IF EXISTS indicadores.log_respostas_numerica DROP CONSTRAINT IF EXISTS log_respostas_numerica_usuario_id_fkey;
ALTER TABLE IF EXISTS indicadores.log_respostas_numerica DROP CONSTRAINT IF EXISTS log_respostas_numerica_variavel_id_fkey;
ALTER TABLE IF EXISTS indicadores.resposta DROP CONSTRAINT IF EXISTS resposta_variavel_agendamento_unique;



DROP TABLE IF EXISTS indicadores.log_resposta_numerica CASCADE;
-- Remover a tabela indicadores.resposta_numerica, pois depende de indicadores.resposta

DROP TABLE IF EXISTS indicadores.resposta_numerica CASCADE;

-- Remover a tabela indicadores.resposta
DROP TABLE IF EXISTS indicadores.resposta CASCADE;

-- Remover a tabela indicadores.variavel_permissao
DROP TABLE IF EXISTS indicadores.variavel_permissao CASCADE;


COMMIT;
