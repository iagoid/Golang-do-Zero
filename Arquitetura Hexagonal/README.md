https://youtu.be/wYyHZL1r7dQ

Tem objetivo de isolar camadas, cada qual com sua responsabilidade, tendo como objetivo isolar totalmente a lógica da aplicação do mundo externo.

**Adapters**
Camada separada em duas partes, sendo uma o input (todos os tipos e formas de entrada de dados da aplicação) e output (todos os tipos de formas de saída)

*Input*
-controllers/consumer/listener: são formas de entrada de dados que existem num adapter input
-model: modelo de entrada sendo separado por request e response
-mappers/converteres: mapeiam o modelo de request/response para o modelo de domínio e vice-versa
-validations (opcional): pode ser camada que valida os objetos antes de serem convertidos para domain

*OutPut*
-restClient/grpcClient/producer/repository: formas de saida de dadados
-model para modelo de saída: para API, filas, entity de banco de dados (request e response)
-mappers/converteres: mapeiam o modelo de request/response para o modelo de domínio e vice-versa
-clientServices (opcional): para fazer o enriquecimento de informações

**Aplication**
Camadas mais importantes da aplicação, a application, temos todas as regras de negócio em services e o mais importante: a camada domain, onde temos o objeto mais forte da aplicação inteira
-Service:: camada onde tudo será validado e criado na sua aplicação
-useCases/Ports: onde se tem o acesso as interfaces de entrada e saída
-Domain: identifica o objeto forte da aplicação, onde você terá todas as informações daquele fluxo no qual está criado

**Configurations**
Configurações referentes a acesso externo, ambiente, feature-toggles, etc
Não existe um padrão, mas sempre bom seguir uma ideia
- Repository: adapter/output/database/mongodb
  Configuration: database/mongodb
- HttpTestMS: adapter/output/http/testMs
  Configuration: http/testMs