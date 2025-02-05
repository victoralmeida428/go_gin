begin;

INSERT INTO indicadores.nivel (nome)
VALUES 
    ('Stat 1'),
    ('Stat 2'),
    ('Stat 3');

INSERT INTO indicadores.tipo_variavel (nome) 
VALUES 
    ('Numérico');

INSERT INTO indicadores.variavel (texto, tipo_variavel_id, pergunta_id, possui_item, obrigatorio)
VALUES
	('COVID', 1, null, true, true),
	('DENGUE', 1, null, true, true),
	('H1N1', 1, null, true, false),
	('INFLUENZA', 1, null, true, false),
	('HIV', 1, null, true, false),
	('HEPATITE A', 1, null, true, false),
	('HEPATITE B', 1, null, true, false),
	('HEPATITE C', 1, null, true, false),
	('HEPATITE D', 1, null, true, false),
	('HEPATITE E', 1, null, true, false),
	('Número de exames realizados', null, 1, false, true),
	('Número de exames positivos', null, 1, false, true),
	('Número de exames realizados', null, 2, false, true),
	('Número de exames positivos', null, 2, false, true),
	('Número de exames realizados', null, 3, false, false),
	('Número de exames positivos', null, 3, false, false),
    ('Número de exames realizados', null, 4, false, false),
    ('Número de exames positivos', null, 4, false, false),
    ('Número de exames realizados', null, 5, false, false),
    ('Número de exames positivos', null, 5, false, false),
    ('Número de exames realizados', null, 6, false, false),
    ('Número de exames positivos', null, 6, false, false),
    ('Número de exames realizados', null, 7, false, false),
    ('Número de exames positivos', null, 7, false, false),
    ('Número de exames realizados', null, 8, false, false),
    ('Número de exames positivos', null, 8, false, false),
    ('Número de exames realizados', null, 9, false, false),
    ('Número de exames positivos', null, 9, false, false),
    ('Número de exames realizados', null, 10, false, false),
	('Número de exames positivos', null, 10, false, false);

INSERT INTO indicadores.formulario(nome, descricao, ativo, nivel_id)
VALUES (
	'Indicadores Epidemiológicos',
	'Os dados a seguir serão utilizados para calcular semanalmente os Indicadores de Produção e Positividade de cada exame listado. Esses dados possuem um valor fundamental para publicações de caráter técnico-científico no setor da saúde, bem como para o monitoramento do cenário epidemiológico no país.',
	true,
	1
);

INSERT INTO
	indicadores.versao_formulario(formulario_id, versao)
	VALUES (1, 1);

INSERT INTO
    indicadores.variavel_formulario(versao_formulario_id, variavel_id)
VALUES
    (1, 11),
    (1, 12),
    (1, 13),
    (1, 14),
    (1, 15),
    (1, 16),
    (1, 17),
    (1, 18),
    (1, 19),
    (1, 20),
    (1, 21),
    (1, 22),
    (1, 23),
    (1, 24),
    (1, 25),
    (1, 26),
    (1, 27),
    (1, 28),
    (1, 29),
    (1, 30);

INSERT INTO indicadores.empresa (nome, nivel_id)
VALUES ('Controllab', 3);

INSERT INTO indicadores.periodicidade (nome, intervalo)
VALUES 
('Semanal', '7 days'),
('Mensal', '1 mon'),
('Trimestral', '3 mons'),
('Semestral', '6 mons'),
('Anual', '1 year');


INSERT INTO indicadores.metodo (nome)
VALUES 
('Soma'),
('Média');


commit;