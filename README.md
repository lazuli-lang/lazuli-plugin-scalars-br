# lazuli-plugin-scalars-br

Lazuli `@plugin/scalars-br` provides pure-Go, dependency-free validators and formatters for Brazilian semantic scalar values: CPF, CNPJ, CEP, and Brazilian DDD phone numbers. Until Lazuli's locale semantic-type plugin mechanism lands, consumers call this package directly from generated or hand-written Go handlers.

## Surface

| Semantic type | Alias | Validator | Formatter | Notes |
|---|---|---|---|---|
| BrazilianCPF | `@semantic.BrazilianCPF` | `ValidateCPF` | `FormatCPF` | 11 digits, check-digit validated, rejects all-same digits |
| BrazilianCNPJ | `@semantic.BrazilianCNPJ` | `ValidateCNPJ` | `FormatCNPJ` | 14 digits, check-digit validated, rejects all-same digits |
| BrazilianCEP | `@semantic.BrazilianCEP` | `ValidateCEP` | `FormatCEP` | 8 digits, format-only |
| BrazilianPhone | `@semantic.BrazilianPhone` | `ValidatePhone` | `FormatPhone` | DDD plus 8 or 9 digits, DDD whitelist |

## Install

```bash
go get lazuli.dev/plugin/scalars-br@v0.1.0
```

## Quick start

```go
package host

import scalarsbr "lazuli.dev/plugin/scalars-br"

func ValidateHostDocument(cpf string) error {
	return scalarsbr.ValidateCPF(cpf)
}
```

```go
cpf, err := scalarsbr.FormatCPF("52998224725")
cnpj, err := scalarsbr.FormatCNPJ("11222333000181")
cep, err := scalarsbr.FormatCEP("01001000")
phone, err := scalarsbr.FormatPhone("99999-0000", "11")
_, _, _, _ = cpf, cnpj, cep, phone
```

```toml
[plugins]
"@plugin/scalars-br" = { module = "lazuli.dev/plugin/scalars-br", version = "v0.1.0" }
```
