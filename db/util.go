package db

func internoParaBancoDeDados(v string) string {
switch v {
	case "saude": return "SAÚDE"
	case "desporto-e-lazer": return "DESPORTO E LAZER"
	case "energia": return "ENERGIA"
	case "legislativa": return "LEGISLATIVA"
	case "administração": return "ADMINISTRAÇÃO"
	case "educação": return "EDUCAÇÃO"
	case "direitos-da-cidadania": return "DIREITOS DA CIDADANIA"
	case "urbanismo": return "URBANISMO"
	case "segurança-pública": return "SEGURANÇA PÚBLICA"
	case "ciência-e-tecnologia": return "CIÊNCIA E TECNOLOGIA"
	case "previdência-social": return "PREVIDÊNCIA SOCIAL"
	case "assistência-social": return "ASSISTÊNCIA SOCIAL"
	case "comércio-e-servicos": return "COMÉRCIO E SERVIÇOS"
	case "judiciária": return "JUDICIARIA"
	case "encargos-especiais": return "ENCARGOS ESPECIAIS"
	case "gestão-ambiental": return "GESTÃO AMBIENTAL"
	case "cultura": return "CULTURA"
	case "habitação": return "HABITAÇÃO"
	case "agricultura": return "AGRICULTURA"
	case "saneamento": return "SANEAMENTO"
	case "trabalho": return "TRABALHO"
	case "comunicações": return "COMUNICAÇÕES"
	case "essencial-a-justiça": return "ESSÊNCIAL A JUSTIÇA"
	case "transporte": return "TRANSPORTE"
	default: return ""
	}
}

func bancoDeDadosParaInterno(v string) string {
	switch v {
		case "SAÚDE": return "saúde"
		case "DESPORTO E LAZER": return "desporto-e-lazer"
		case "ENERGIA": return "energia"
		case "LEGISLATIVA": return "legislativa"
		case "ADMINISTRAÇÃO": return "administração"
		case "EDUCAÇÃO": return "educação"
		case "DIREITOS DA CIDADANIA": return "direitos-da-cidadania"
		case "URBANISMO": return "urbanismo"
		case "SEGURANÇA PÚBLICA": return "seguranca-pública"
		case "CIÊNCIA E TECNOLOGIA": return "ciência-e-tecnologia"
		case "PREVIDÊNCIA SOCIAL": return "previdência-social"
		case "ASSISTÊNCIA SOCIAL": return "assistência-social"
		case "COMÉRCIO E SERVIÇOS": return "comércio-e-servicos"
		case "JUDICIARIA": return "judiciária"
		case "ENCARGOS ESPECIAIS": return "encargos-especiais"
		case "GESTÃO AMBIENTAL": return "gestão-ambiental"
		case "CULTURA": return "cultura"
		case "HABITAÇÃO": return "habitação"
		case "AGRICULTURA": return "agricultura"
		case "SANEAMENTO": return "saneamento"
		case "TRABALHO": return "trabalho"
		case "COMUNICAÇÕES": return "comunicações"
		case "ESSÊNCIAL A JUSTIÇA": return "essencial-a-justiça"
		case "TRANSPORTE": return "transporte"
		default: return ""
	}
}
