# Projeto: Comunicação entre Goroutines com Canais (Ping Pong)
## 🚀 Descrição
Este projeto é uma aplicação simples em Go (Golang) que demonstra o uso de Goroutines e Canais para comunicação concorrente entre funções. Inspirado no desafio da plataforma DIO (Digital Innovation One), o objetivo é visualizar o fluxo assíncrono de mensagens (prints) no terminal, criando um comportamento semelhante ao clássico jogo de Ping Pong.

## 🧑‍💻 O que aprendi com este projeto
✅ Como criar e usar Goroutines
✅ Como implementar Canais (Channels) para comunicação entre processos concorrentes
✅ Como usar a estrutura select para manipular múltiplos canais simultaneamente
✅ Como controlar fluxo de execução com time.Sleep
✅ Visualização de execução concorrente com mensagens alternadas no terminal

## ✅ Como o código funciona
  Goroutine ping: Envia a mensagem "ping" a cada 1 segundo para o canal c1.
  Goroutine pong: Envia a mensagem "pong" a cada 2 segundos para o canal c2.
  Função main:
  Cria dois canais (c1 e c2).
  Inicia as duas Goroutines.
  Usa um laço com um select para ouvir os dois canais, imprimindo as mensagens assim que elas chegam.
  No exemplo atual, o programa faz apenas 2 iterações (imprime duas mensagens) antes de finalizar.

## 🛠️ Como executar o projeto localmente
Passos:
```
# Clone o repositório
git clone https://github.com/seu-usuario/nome-do-repositorio.git

# Acesse o diretório do projeto
cd nome-do-repositorio

# Execute o código
go run main.go
```

## 🎯 Resultado Esperado
Ao executar, o terminal deve exibir algo como:

```
ping
pong
(A ordem pode variar dependendo do agendamento das goroutines pelo scheduler do Go)
```
## 📚 Referências

- [Documentação Oficial do Go - Goroutines](https://go.dev/doc/effective_go#goroutines)
- [Documentação Oficial do Go - Channels](https://go.dev/doc/effective_go#channels)
- [Pacote time - Golang](https://pkg.go.dev/time)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go by Example - Select](https://gobyexample.com/select)
- [Curso e Desafio - Digital Innovation One (DIO)](https://web.dio.me/track/formacao-go-developer)


## 📌 Sobre o Desafio
Este projeto faz parte do módulo "Entendendo o Fluxo de Execução Concorrente em Go", dentro da formação de desenvolvedores da DIO.

## ✨ Contribuições e Feedback
Fique à vontade para abrir issues, sugerir melhorias ou até fazer um fork e criar suas próprias versões.

👨‍💻 Desenvolvido por **Luis Arthur**
