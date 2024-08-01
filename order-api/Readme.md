# Nome do Projeto

API para cadastro de pedidos

## Sobre

Este projeto é uma API de microserviços para criação de pedidos a partir de uma usuário cadastrado.


## Tecnologias
- Golang v1.22.4  (Liguagem de Programação)
- Docker
- Postgres:latest para persistência dos dados
- Redis para cache
- Elasticseach para buscas mais complexas envolvendo muitos pedidos e filtros
- Swagger (Documentação da API)
- Dbeaver para consultas rápidas
- Postman e curl para test
- Git e Github (Versionamento e armazenamento de código fonte)

## Aplicação e Docs para rodar em ambiente local
- Swagger: http://localhost:3001/docs/v1/order-api/swagger/index.html#/
- Serviço: http://localhost:3001/api/v1/order-api/

## Instalação e Requisitos
Para rodar em ambiente local estes são os requisitos necessários:

1 - Golang na versão 1.22
2 - Docker e Docker compose
3 - Pacote Make para utilização do comando make


Passos para instalar e configurar o ambiente de desenvolvimento.

1. Clone o repositório:
    ```bash
    git clone git@github.com:SelecaoSerasaConsumidor/BE-RodolfoAlves.git
    ```
2. Vá para o diretório do projeto:
    ```bash
    cd BE-RodolfoAlves
    ```
3. Instale as dependências:
    ```bash
    make deps
    ```
4. Gere a documentação
     ```bash
    make generate
    ```

## Uso

Instruções sobre como usar o projeto:
Após gerar as dependências você poderá executar este micro serviço de duas formas:
1 - Forma isolada: Dentro do diretório do microserviço você pode rodar a aplicação de forma isolada para tests através do comando:
Execute a aplicação:
    ```bash
    make build-docker
    ```
Parar aplicação:
    ```bash
        make stop-docker
        ```
Iniciar aplicação sem buid:
     ```bash
        make start-docker
        ```
Para mais comando vide o arquivo Dockerfile

## Estrutura do Projeto
O projeto foi desenvolvido numa arquitetura que mistura principios de clean archtecture e arquiterura hexagonal separa a aplicação em camadas que se comunicam através de interfaces.

Durante o desenvolvimento foram utilizados vários pacotes externos que podem ser consultados após o download das dependências


## Licença

Informações sobre a licença do projeto.

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Informações de contato para dúvidas ou sugestões.

- Nome: Rodolfo Alves Gonçalves
- Email: rodolfoalves.inf@gmail.com
- LinkedIn: [Seu LinkedIn](https://www.linkedin.com/in/rodolfoalvesg)
- GitHub: [Seu GitHub](https://github.com/rodolfoalvesg)
