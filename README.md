# Conversor de Relatórios CSV para JSON e Parquet

Este projeto é uma aplicação CLI desenvolvida em Go, com o objetivo de converter relatórios em formato CSV para os formatos JSON e Parquet. A aplicação foi estruturada utilizando o **Padrão de Projeto Adapter**, facilitando a extensão e a manutenção do código para suportar múltiplos formatos de saída. Este projeto foi desenvolvido com foco em modularidade, reusabilidade e eficiência, sendo uma adição relevante para o portfólio de soluções CLI.

## Funcionalidades

- **Leitura de CSV**: Aceita um arquivo CSV como entrada, especificado pelo usuário via linha de comando.
- **Conversão para JSON e Parquet**: Gera relatórios em dois formatos de saída (JSON e Parquet), conforme a escolha do usuário.
- **Estrutura Modular**: Implementado com o Padrão de Projeto Adapter, permitindo a fácil integração de novos tipos de saída sem alterações na lógica principal.
- **CLI Customizável**: Parâmetros CLI para configurar o arquivo de entrada, o nome do arquivo de saída e o tipo de formato desejado.

## Estrutura do Projeto

```plaintext
├── app
│   ├── adapter
│   │   ├── json_adapter.go              # Adapter para conversão para JSON
│   │   └── parquet_adapter.go           # Adapter para conversão para Parquet
│   └── legacy
│       ├── legacy_reporter_generator.go # Código legado para geração de relatórios
│       └── report_generator_interface.go# Interface para geração de relatórios
├── output_report                        # Diretório para os arquivos de saída gerados
│   ├── export_json.json                 # Saída no formato JSON
│   └── export_parquet.parquet           # Saída no formato Parquet
├── .env                                 # Arquivo de configuração de ambiente
├── converterreportscsvtojson            # Executável do conversor
├── export.csv                           # Arquivo CSV de entrada
├── go.mod                               # Gerenciador de dependências do Go
├── go.sum                               # Checksum das dependências do Go
└── main.go                              # Arquivo principal da aplicação
```

## Parâmetros CLI

Os parâmetros CLI permitem ao usuário especificar o caminho do arquivo de entrada, o caminho do arquivo de saída, e o formato de saída desejado.

**Exemplo de uso:**

```bash
go run main.go -inputfile=export.csv -outputfile=output_report -outputtype=json
```

**Parâmetros:**

- `inputfile`: Especifica o caminho do arquivo CSV de entrada. O padrão é `csv_report`.
- `outputfile`: Especifica o caminho para o diretório ou arquivo de saída. O padrão é `output_report`.
- `outputtype`: Define o formato do relatório de saída. Opções: `json` ou `parquet`.

## Configuração do Ambiente

O arquivo `.env` contém uma variável de configuração importante:

```plaintext
PATH_OUTPUT=./output_report/
```

Essa variável define o caminho padrão onde os arquivos de saída serão armazenados. O projeto utilizará o valor especificado em `PATH_OUTPUT` para salvar os relatórios convertidos, a menos que o usuário especifique um caminho de saída diferente pelo parâmetro CLI `outputfile`.

## Padrão de Projeto Utilizado: Adapter

Este projeto foi desenvolvido utilizando o **Padrão de Projeto Adapter**, que permite a integração de múltiplos formatos de saída de maneira independente. Cada adaptador implementa uma interface comum de geração de relatórios, tornando o código modular e extensível. Novos formatos podem ser adicionados no futuro sem alterar o núcleo da aplicação.

### Como Funciona o Adapter

- **Adapters JSON e Parquet**: Os arquivos `json_adapter.go` e `parquet_adapter.go` contêm a lógica específica para converter o CSV de entrada em JSON e Parquet, respectivamente.
- **Interface de Geração de Relatórios**: A interface `report_generator_interface.go` define os métodos que cada adaptador deve implementar, garantindo a padronização.
- **Compatibilidade com Código Legado**: O código legado `legacy_reporter_generator.go` foi integrado usando essa interface, garantindo compatibilidade e permitindo a reutilização de lógica antiga sem modificações significativas.

## Requisitos

- **Go**: Certifique-se de ter o Go instalado (versão recomendada: >= 1.16).
- **Dependências**: Execute `go mod tidy` para instalar as dependências listadas no `go.mod`.

## Execução

Para executar a aplicação, utilize o comando abaixo, substituindo os valores dos parâmetros conforme necessário:

```bash
go run main.go -inputfile=export.csv -outputfile=output_report -outputtype=json
```

## Diretório de Saída

Os arquivos gerados serão armazenados no diretório configurado em `PATH_OUTPUT` no arquivo `.env` (padrão: `./output_report/`), ou no caminho especificado pelo parâmetro CLI `outputfile`.

## Licença

Este projeto é licenciado sob a Licença MIT.

---

Este README agora inclui a explicação sobre a configuração `PATH_OUTPUT` no arquivo `.env`, destacando seu papel no caminho de saída dos relatórios gerados.