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
REQUESTS_PER_SECOND=10           # Limite padrão de requisições/segundo
WindowSize=10                    # Tempo de expiração em segundos
API_KEY_LIMITS={"abc123": 100}   # Chaves especiais com limites personalizados (API_KEY)
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

| Variável            | Descrição                                  |
|---------------------|--------------------------------------------|
| REDIS_ADDR          | Endereço do servidor Redis                 |
| WEB_SERVER_PORT     | Porta HTTP do serviço                      |
| REQUESTS_PER_SECOND | Limite global de requisições/segundo       |
| WINDOW_SIZE         | Janela temporal para contagem (segundos)   |
| API_KEY_LIMITS      | Mapa de tokens com limites personalizados  |

## Teste de Carga

Para realizar testes de carga no serviço, você pode utilizar a ferramenta [vegeta](https://github.com/tsenart/vegeta):

```bash
# Teste simples com uma API_KEY específica (6 segundos de duração)
echo "GET http://localhost:8080/hello" | ./vegeta attack -duration=6s -header "API_KEY: def456" | tee results.bin | ./vegeta report
