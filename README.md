# Customer Eligibility - Vehicle Payment

Sistema de regras de elegibilidade do cliente para financiamento de veículos.

## 👤 Regras de Elegibilidade Básica

### Dados do Cliente

- **Idade mínima:** 18 anos (alguns bancos exigem 21)
- **CPF:** válido e regularizado (sem pendências na Receita)
- **Renda mínima comprovada:** geralmente 2-3x o valor da parcela
- **Tempo mínimo de trabalho/renda:** 3-6 meses
- **Residência fixa:** comprovada

### Análise de Crédito

- **Score de crédito mínimo:** acima de 400-500 pontos
- **Consulta em bureaus:** Serasa, Boa Vista, SPC
- **Ausência de protestos** ou execuções judiciais ativas
- **Limite de restrições:** sem negativações nos últimos 6-12 meses
- **Cheques devolvidos:** máximo permitido (ex: nenhum nos últimos 12 meses)

## 🔍 Validações Realizadas

- ✅ Validação de CPF
- ✅ Consulta de score de crédito
- ✅ Verificação de negativações
- ✅ Análise de renda x capacidade de pagamento
- ✅ Consulta de protestos e execuções
- ✅ Histórico de cheques devolvidos

## 🚀 Tecnologias

- go

