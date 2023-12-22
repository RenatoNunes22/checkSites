# Verificador de Disponibilidade de Sites com Registro de Logs

Este projeto é uma aplicação simples e eficiente para verificar a disponibilidade de sites, proporcionando uma visão rápida e fácil do status online de diversos endereços da web. Além disso, o sistema oferece a funcionalidade de salvar logs diários, permitindo o rastreamento de eventos e incidentes ao longo do tempo.

## Recursos Principais:

### 1. Verificação de Disponibilidade

- **Monitoramento Contínuo:** O aplicativo verifica periodicamente se os sites listados estão online.
- **Feedback Visual:** Interface intuitiva que exibe claramente o status de cada site (online/offline).

### 2. Registro de Logs

- **Histórico Diário:** Armazenamento de logs detalhados, permitindo a análise retrospectiva de eventos.
- **Timestamps:** Cada entrada no log é carimbada com a data e a hora exatas da ocorrência.

### 3. Configuração Flexível

- **Lista Personalizável de Sites:** Adicione, remova ou modifique facilmente os sites a serem monitorados.
- **Intervalos de Verificação Ajustáveis:** Configure a frequência das verificações de disponibilidade.

## Configuração Inicial:

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Crie o Arquivo de Sites**

- Crie um arquivo chamado sites.txt na raiz do projeto.
- Adicione cada URL de site em uma linha separada.

```env
https://www.exemplo1.com
https://www.exemplo2.com
https://www.exemplo3.com
```

3. **Execute o Aplicativo**
   ```bash
   go run checkSites.go
   ```
