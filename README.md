# GitOps Infrastructure State Reconciler

Agente que detecta drift entre estado real do servidor e repositório Git — reverte mudanças manuais via SSH.

## Stack

- Go, Kubernetes (opcional), ArgoCD patterns

## Princípios

- **Imutabilidade**: servidor = manifestação do Git
- **Event-driven**: webhook Git push → reconcile
- **Drift detection**: hash de configs a cada N segundos

## Arquitetura

```
Git repo ──webhook──► Reconciler Agent ──SSH/Ansible──► Servers
                         │
                    state hash store
```

## Run

```bash
go run ./cmd/reconciler --repo https://github.com/org/infra --interval 60s
```

Webhook: `POST /hooks/git` com secret HMAC.

Documentação: [docs/IMMUTABILITY.md](docs/IMMUTABILITY.md)
