# Informações sobre o desenvolvimento e execução da aplicão
A aplicação foi tentando aplicar os príncipios de arquitetura limpa com separação de contextos ou melhor, de camadas.
- Na aplicação de USER-API pode-se cadastrar um usuário, atualizar, lista por id e por filtro e excluir. Foram implementadas várias lógicas para que o sistema mantenha a consistencia, como por exemplo cadastro dois usuários iguais. 

- Na aplicação de ORDER-API pode-se cadastrar um pedido desde que o usuário exista. Também pode-se deletar, atualizar e excluir. Também foi implementada algumas lógicas de consistências. Incluindo comunicação com o microserviço da USER-API para validação do usuário.

- Em ambos os serviços utilizei o redis, o postgres e o elasticseach de forma compatilhada para economizar tempo.

- Existem alguns pontos de melhorias que podemos discutir na code review.

- Para que ambos os serviços funcionem será necessário entrar em cada um deles e executar o comando do make para instação de depenências conforme documentação interna. A partir disso volte nesse diretório e execute o comando para que a imagem dos microserviços e suas aplicações dependentes sejam criadas e executadas em containeres.

####Instruções sobre como usar o projeto:
Após gerar as dependências você poderá executar as duas aplicações de uma vez:
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