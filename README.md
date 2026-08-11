# GoSys

<p align="center">
  <img src="https://cdn.simpleicons.org/go/808080" alt="Go" width="100">
</p>

<p align="center">
  <strong>Toolkit de utilidades de sistema escrita em Go.</strong>
</p>

<p align="center">
  Ferramentas rápidas, leves e reutilizáveis para operações de sistema.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25%2B-808080?style=flat-square&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/License-MIT-808080?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/Status-Em%20desenvolvimento-orange?style=flat-square" alt="Status">
</p>

---

## Sobre

**GoSys** é uma toolkit open-source desenvolvida em Go para reunir diversas utilidades de sistema em uma única ferramenta.

O projeto oferece funcionalidades para manipulação e análise de arquivos, armazenamento, processos e informações do sistema operacional.

A proposta é simples: **resolver tarefas pequenas e recorrentes de forma rápida, através de uma única CLI e uma biblioteca reutilizável.**

O GoSys não depende de interface gráfica, banco de dados ou serviços externos.

---

## Funcionalidades

### Files

Ferramentas relacionadas a arquivos e diretórios.

- `scan` — analisa arquivos e diretórios
- `duplicates` — encontra arquivos duplicados
- `compare` — compara dois diretórios
- `integrity` — verifica a integridade de arquivos
- `rename` — renomeia arquivos em lote
- `hash` — calcula hashes de arquivos
- `search` — procura arquivos
- `empty` — encontra arquivos e diretórios vazios
- `size` — analisa o tamanho de arquivos e diretórios
- `modified` — filtra arquivos por data de modificação
- `extension` — filtra arquivos por extensão

### Disk

Ferramentas relacionadas ao armazenamento.

- `usage` — mostra espaço utilizado e disponível
- `largest` — encontra os maiores arquivos
- `dirs` — identifica os maiores diretórios
- `stats` — apresenta estatísticas de armazenamento
- `drives` — lista unidades disponíveis
- `size` — calcula o espaço ocupado por diretórios

### Process

Ferramentas para inspeção de processos.

- `list` — lista processos em execução
- `find` — procura processos por nome ou PID
- `info` — mostra informações detalhadas de um processo
- `tree` — mostra a árvore de processos
- `stats` — mostra informações de consumo
- `pid` — consulta informações através do PID

### System

Informações gerais da máquina.

- `cpu` — informações da CPU
- `memory` — informações da memória RAM
- `os` — informações do sistema operacional
- `uptime` — tempo desde a inicialização
- `hostname` — nome da máquina
- `arch` — arquitetura do sistema
- `info` — resumo geral da máquina

---

## Exemplos

```bash
gosys files scan ./Downloads
gosys files duplicates ./Downloads
gosys files compare ./backup-old ./backup-new
gosys files hash ./arquivo.zip

gosys disk usage C:\
gosys disk largest C:\

gosys process list
gosys process find chrome
gosys process info 1240
gosys process tree

gosys system cpu
gosys system memory
gosys system os
gosys system uptime
gosys system info

```
## Estrutura

```bash

                         ┌─────────────────────┐
                         │        GoSys        │
                         │   System Toolkit    │
                         └──────────┬──────────┘
                                    │
                    ┌───────────────┴───────────────┐
                    │                               │
             ┌──────▼──────┐                 ┌──────▼──────┐
             │   Library   │                 │     CLI     │
             │    /pkg     │                 │    /cmd     │
             └──────┬──────┘                 └──────┬──────┘
                    │                               │
       ┌────────────┼────────────┬────────────┐     │
       ▼            ▼            ▼            ▼     ▼
    Files         Disk        Process       System Commands
       │            │            │            │
       └────────────┴────────────┴────────────┘
                              │
                    ┌─────────▼─────────┐
                    │  OS Abstraction   │
                    └─────────┬─────────┘
                              │
                 ┌────────────┼────────────┐
                 ▼            ▼            ▼
              Windows       Linux        macOS

```
