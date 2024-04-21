# api.devbook

## Requisitos
- É preciso ter o GO na versão **`1.2*`**;
- MySQL 8.*;

## Pré tutorial
1. Abra a pasta **`sql`**;
2. Copie e execute os comandos SQL do arquivo **`sql.sql`** para criar as tabelas;
3. Depois, copie e execute os comandos SQL do aquivo **`data.sql`** para alimentar as tabelas.
4. No arquivo **`.env`** coloque as informações conforme a imagem:

![Screenshot 2024-04-21 125533](https://github.com/IuryHirabara/public.api.devbook/assets/107448972/2a2314e6-d3fd-4f74-a0c0-096a0c06e98d)

## Tutorial
1. Navegue até o diretório do projeto e rode o comando `go mod tidy` para baixar as dependências;
2. Em seguida, execute o comando `go run main.go` para iniciar o servidor da API;
3. Utilize uma ferramenta para fazer as requisições para API como o **`Postman`** ou inicie o [Frontend](https://github.com/IuryHirabara/public.app.devbook) da aplicação.
