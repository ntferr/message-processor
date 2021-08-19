# message-processor 
API que processa uma mensagem SQS.

<img title="sqs" height="50" alt="aws sqs" src="https://seeklogo.com/images/A/aws-sqs-simple-queue-service-logo-8884A71ECB-seeklogo.com.png"><img title="Golang" height="50" src="https://cdn.iconscout.com/icon/free/png-256/go-77-1175166.png" alt="golang"/>

## Entendo a estrutura do projeto:
- workdir
  - cmd
    - main.go {Por onde inicializa o projeto}
  - internal
    - controllers
    - http
    - models
    - service
    - settings
    - utils
  - .env {Variáveis da aplicação}
  

## Arquitetura:

## Frameworks:
  [Netflix envoriment](github.com/Netflix/go-env)
  
  [GIN CORS](github.com/gin-contrib/cors)
  
  [GIN GZIP](github.com/gin-contrib/gzip)
  
  [GIN](github.com/gin-gonic/gin)
  
  [Google UUID](github.com/google/uuid)
  
## Como subir o serviço no VS Code:
