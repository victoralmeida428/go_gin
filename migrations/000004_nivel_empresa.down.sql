BEGIN;

-- Remover a coluna empresa_id da tabela indicadores.usuario
ALTER TABLE indicadores.usuario DROP COLUMN empresa_id;

-- Adicionar novamente a coluna empresa na tabela indicadores.usuario
ALTER TABLE indicadores.usuario ADD COLUMN empresa VARCHAR(255);

-- Remover a coluna metodo_id da tabela indicadores.agendamento
ALTER TABLE indicadores.agendamento DROP COLUMN metodo_id;

-- Adicionar novamente a coluna unificado na tabela indicadores.agendamento
ALTER TABLE indicadores.agendamento ADD COLUMN unificado BOOLEAN;

-- Remover os dados inseridos na tabela indicadores.nivel

-- Remover tabelas criadas: indicadores.metodo, indicadores.nivel, e indicadores.empresa
DROP TABLE IF EXISTS indicadores.metodo CASCADE;
DROP TABLE IF EXISTS indicadores.empresa CASCADE;
DROP TABLE IF EXISTS indicadores.nivel CASCADE;

COMMIT;
