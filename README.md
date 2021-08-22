# message-processor
### [Processador feito, falta consumidor]
API que processa uma mensagem padrão SQS.

<img title="sqs" height="50" alt="aws sqs" src="https://seeklogo.com/images/A/aws-sqs-simple-queue-service-logo-8884A71ECB-seeklogo.com.png"><img title="Golang" height="50" src="https://cdn.iconscout.com/icon/free/png-256/go-77-1175166.png" alt="golang"/>

## Entendo a estrutura do projeto:
- workdir
  - cmd
    - main.go {Por onde inicializa o projeto}
  - internal
    - controllers
      - login.go {Onde fica a lógica do processo/Regra de negócio} 
    - http
      - handlers.go {Handlers do server}
      - server.go {Inicialização do Gin}
    - models
      - user.go {Estrutura de um usário}
    - service
      - aws
        - sqs.go {Instância de sqs e envio de mensagem} 
    - settings
      - settings.go {Configurações do projeto/Envs}
    - utils
      - uuid.go {Gera e retorna um UUID v5, com base no payload}
  - .env {Variáveis da aplicação}
  

## Arquitetura:
![arquitetura](https://github.com/ntferr/images/blob/main/message-processor/message.png)

## Payload + endpoint:
### Endpoint
`127.0.0.1:8080/v1/login/create`
### Payload
`{
	"name": "String",
	"email": "String",
	"birthdate": "String",
	"password": "String"
}`
#### Example:
![payload-insomnia](https://github.com/ntferr/images/blob/main/message-processor/payload-insomnia.png)

## Frameworks:
  [Netflix envoriment](github.com/Netflix/go-env)
  
  [GIN CORS](github.com/gin-contrib/cors)
  
  [GIN GZIP](github.com/gin-contrib/gzip)
  
  [GIN](github.com/gin-gonic/gin)
  
  [Google UUID](github.com/google/uuid)
  
## Como subir o serviço no VS Code:

## Output:
![payload-aws](https://github.com/ntferr/images/blob/main/message-processor/payload-aws.png)
