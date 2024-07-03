# Desafio Multithreading
<p>
<a href="https://github.com/waanvieira/desafio-multithreading/blob/main/LICENSE"><img src="https://img.shields.io/packagist/l/laravel/framework" alt="License"></a>
</p>

# Sobre o projeto
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

# Tecnologias utilizadas
- go1.22.1 linux/amd64 

# Como executar o projeto

## Pré-requisitos
GO - https://go.dev/

```bash
# clonar repositório
git clone https://github.com/waanvieira/simplified_payment_platform.git

# entrar na pasta do projeto back end
cd desafio-multithreading 

# Executar o server
go run main.go

```
# Autor

Wanderson Alves Vieira

https://www.linkedin.com/in/wanderson-alves-vieira-59b832148
