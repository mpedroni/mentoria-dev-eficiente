# applications/product/src/main/java/org/apache/ofbiz/shipment/shipment/ShipmentServices.java

## createShipmentEstimate

### 1a leitura

- o que é DispatchContext?
  - parece ser o Context como o do Spring
- o que é Delegator?
  - parece ser um service de utilidade de BD, tipo um [padrão que abstrai complexidade]
- o que é GenericValue?
- o que é ProductStoreShipmentMeth?

  - método de envio de um produto

- tenta achar um ProductStoreShipmentMeth pelo id (productStoreShipMethId) e retorna erro se não encontra

  - id vem do context (request?)

## 2a leitura

- busca o shipment method no banco
  - definido na request
  - retorna erro se não existe
- cria um ShipmentCostEstimate (GenericValue)
- define parametros
- realiza verificações?
- persiste a estimativa
- retorna o id

## refatoração (map da response)

- começar simples
  - eu iria pro applyQuantityBreak, mas envolve muito mais coisa. Quero começar simples
- usada em praticamente todos os métodos
- melhora a legibilidade? É o problema mais crítico?

- centralizado no ServiceUtil e ModelService
  - ServiceUtil define os tipos de response, ModelService tem os valores (campos e mensagens padrão)
