# Learning Rate Limiter

Este projeto demonstra a implementação de um middleware para definir rate limits em um servidor, protegendo-o contra sobrecarga de requisições. Utilizando Redis como armazenamento para contabilizar e controlar o número de requisições por cliente em determinados intervalos de tempo.

## Características

- Middleware para integração com servidores web
- Controle de requisições baseado em endereço IP ou API_KEY
- Configuração flexível de limites e intervalos
- Armazenamento das informações de contagem utilizando Redis

## Configuração

1. Renomeie o arquivo de configuração:

```bash
cp template.env .env
```

1. Edite o `.env` com suas configurações:

```env
REDIS_ADDR='localhost:6379'      # Endereço do Redis
WEB_SERVER_PORT=:8080            # Porta do serviço
RATE_LIMIT=10                    # Limite padrão de requisições/segundo
EXPIRE=10                        # Tempo de expiração em segundos
TOKEN_TO_LIMIT={"abc123": 100}   # Chaves especiais com limites personalizados (API_KEY)
```

## Execução com Docker

```bash
# Iniciar Redis e o serviço
docker compose up -d --build

# Parar os containers
docker compose down
```

## Execução local (sem Docker)

```bash
# Iniciar Redis
docker compose up -d redis

# Compilar e executar o serviço
go run cmd/main.go
```

## Utilização

Verifique o status do rate limiting:

```bash
curl -X GET http://localhost:8080/hello \
  -H "API_KEY: Bearer abc123"
```

### Exemplo de resposta

```json
{
    "message":"Hello, World!"
}
```

## Variáveis de Ambiente

| Variável        | Descrição                                  | Padrão         |
|-----------------|-------------------------------------------|----------------|
| REDIS_ADDR      | Endereço do servidor Redis                | localhost:6379 |
| WEB_SERVER_PORT | Porta HTTP do serviço                     | :8080          |
| RATE_LIMIT      | Limite global de requisições/segundo      | 2              |
| EXPIRE          | Janela temporal para contagem (segundos)  | 10             |
| TOKEN_TO_LIMIT  | Mapa de tokens com limites personalizados | {}             |
