CREATE VIEW indicadores.view_variavel AS
SELECT
    COALESCE(subitem.id, item.id) AS id,
    item.texto AS pergunta,
    subitem.texto AS item,
    item.tipo_variavel_id,
    item.possui_item,
    item.obrigatorio
FROM
    indicadores.variavel item
LEFT JOIN
    indicadores.variavel subitem ON item.id = subitem.pergunta_id
WHERE
    subitem.texto IS NOT NULL
    OR item.pergunta_id IS NULL;