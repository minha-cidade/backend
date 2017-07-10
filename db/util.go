package db

func internoParaBancoDeDados(v string) string {
	switch v {
	case "saude":
		return "SAÚDE"
	case "desporto-e-lazer":
		return "DESPORTO E LAZER"
	case "energia":
		return "ENERGIA"
	case "legislativa":
		return "LEGISLATIVA"
	case "administracao":
		return "ADMINISTRAÇÃO"
	case "educacao":
		return "EDUCAÇÃO"
	case "direitos-da-cidadania":
		return "DIREITOS DA CIDADANIA"
	case "urbanismo":
		return "URBANISMO"
	case "seguranca-publica":
		return "SEGURANÇA PÚBLICA"
	case "ciencia-e-tecnologia":
		return "CIÊNCIA E TECNOLOGIA"
	case "previdencia-social":
		return "PREVIDÊNCIA SOCIAL"
	case "assistencia-social":
		return "ASSISTÊNCIA SOCIAL"
	case "comercio-e-servicos":
		return "COMÉRCIO E SERVIÇOS"
	case "judiciaria":
		return "JUDICIARIA"
	case "encargos-especiais":
		return "ENCARGOS ESPECIAIS"
	case "gestao-ambiental":
		return "GESTÃO AMBIENTAL"
	case "cultura":
		return "CULTURA"
	case "habitacao":
		return "HABITAÇÃO"
	case "agricultura":
		return "AGRICULTURA"
	case "saneamento":
		return "SANEAMENTO"
	case "trabalho":
		return "TRABALHO"
	case "comunicacoes":
		return "COMUNICAÇÕES"
	case "essencial-a-justica":
		return "ESSÊNCIAL A JUSTIÇA"
	case "transporte":
		return "TRANSPORTE"
	default:
		return ""
	}
}

func bancoDeDadosParaInterno(v string) string {
	switch v {
	case "SAÚDE":
		return "saude"
	case "DESPORTO E LAZER":
		return "desporto-e-lazer"
	case "ENERGIA":
		return "energia"
	case "LEGISLATIVA":
		return "legislativa"
	case "ADMINISTRAÇÃO":
		return "administracao"
	case "EDUCAÇÃO":
		return "educacao"
	case "DIREITOS DA CIDADANIA":
		return "direitos-da-cidadania"
	case "URBANISMO":
		return "urbanismo"
	case "SEGURANÇA PÚBLICA":
		return "seguranca-publica"
	case "CIÊNCIA E TECNOLOGIA":
		return "ciencia-e-tecnologia"
	case "PREVIDÊNCIA SOCIAL":
		return "previdencia-social"
	case "ASSISTÊNCIA SOCIAL":
		return "assistencia-social"
	case "COMÉRCIO E SERVIÇOS":
		return "comercio-e-servicos"
	case "JUDICIARIA":
		return "judiciaria"
	case "ENCARGOS ESPECIAIS":
		return "encargos-especiais"
	case "GESTÃO AMBIENTAL":
		return "gestao-ambiental"
	case "CULTURA":
		return "cultura"
	case "HABITAÇÃO":
		return "habitacao"
	case "AGRICULTURA":
		return "agricultura"
	case "SANEAMENTO":
		return "saneamento"
	case "TRABALHO":
		return "trabalho"
	case "COMUNICAÇÕES":
		return "comunicacoes"
	case "ESSÊNCIAL A JUSTIÇA":
		return "essencial-a-justica"
	case "TRANSPORTE":
		return "transporte"
	default:
		return ""
	}
}
