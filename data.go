package main

func Data() {
	typeAnimal1 := TypeAnimal{}
	typeAnimal1.Type = "Touro"
	typeAnimal1.Description = "Macho bovino não castrado"
	db.Save(&typeAnimal1)

	typeAnimal2 := TypeAnimal{}
	typeAnimal2.Type = "Vaca"
	typeAnimal2.Description = "Fêmea bovina"
	db.Save(&typeAnimal2)

	typeAnimal3 := TypeAnimal{}
	typeAnimal3.Type = "Boi"
	typeAnimal3.Description = "Macho bovino castrado"
	db.Save(&typeAnimal3)

	breed1 := Breed{}
	breed1.Breed = "Angus"
	breed1.Description = "Essa é a mais famosa raça de taurinos no Brasil. Seu nome ficou conhecido e a raça se popularizou especialmente a partir do investimento de grandes empresas, como o MC Donald’s, que criou um hambúrguer com a carne Angus. De acordo com a Associação Brasileira de Angus, as principais vantagens da raça para a criação são a alta fertilidade e precocidade, pois atingem a puberdade e o estado de abate mais cedo. Seu diferencial é a ótima qualidade da carne, que é marmorizada e macia."
	db.Save(&breed1)

	breed2 := Breed{}
	breed2.Breed = "Nelore"
	breed2.Description = "É a raça predominante no Brasil, muito procurada por produtores de carne. Sua característica marcante é a pelagem branca, que pode ter tons de cinza claro. Tem orelhas pontiagudas e chifres curtos, mas algumas variações são mochos, ou seja, sem chifres."
	db.Save(&breed2)

	breed3 := Breed{}
	breed3.Breed = "Guzerá"
	breed3.Description = "Foi a primeira raça de zebuíno a ser trazida para o País e é uma das mais antigas do mundo. É reconhecida por possuir um par de chifres grandes e curvados para cima e pode ser direcionada tanto para a pecuária de corte como de leite. A pelagem varia em tons de cinza, do mais claro ao escuro. Os animais apresentam grande porte, a raça é muito fértil e resistente à seca."
	db.Save(&breed3)

	breed4 := Breed{}
	breed4.Breed = "Gir"
	breed4.Description = "Trazido ao Brasil em 1911, das montanhas Gir na Índia, a raça é indicada para a pecuária de leite. Inclusive, a raça Girolando, a mais famosa na produção leiteira no Brasil é resultado de cruzamento de Gir com a vaca Holandesa. Os indivíduos dessa raça apresentam chifres compridos e torcidos para baixo, com orelhas enroladas na parte superior. A pelagem varia do vermelho ao amarelado e pode apresentar pintas. A raça é dócil e as fêmeas têm grande habilidade materna."
	db.Save(&breed4)

	breed5 := Breed{}
	breed5.Breed = "Cangaian"
	breed5.Description = "Essa raça chegou ao Brasil entre 1962 e 1963, vindo da região Sul da Índia. A raça representa um rebanho pequeno e pouco representativo, em números, no Brasil. Os bois têm pequena estatura e possuem chifres longos e grossos, mas são indicados somente à produção de carne, porque não produzem muito leite. São muito resistentes ao calor e a doenças."
	db.Save(&breed5)

	breed6 := Breed{}
	breed6.Breed = "Brahman"
	breed6.Description = "Veio em 1994 dos Estados Unidos e é o resultado do cruzamento de Nelore, Guzerá, Sindi, Cangaian e Indubrasil. A coloração pode ser cinza-claro, cinza-escuro ou vermelho. Não tem chifres e as orelhas são de tamanho médio. É indicado como gado de corte."
	db.Save(&breed6)

	breed7 := Breed{}
	breed7.Breed = "Tabapuã"
	breed7.Description = "Surgiu ao cruzar zebuínos Nelore, Gir e Guzerá com os mochos brasileiros e apesar de ser uma raça  nacional, é criada também em outros países, na maioria da América do Sul. A pelagem varia do branco ao cinza e não possui chifres. É usado na produção de carne, porque tem boa musculatura."
	db.Save(&breed7)

	breed8 := Breed{}
	breed8.Breed = "Sindi"
	breed8.Description = "Originária da província de Sindi, no Paquistão, a raça veio para o Brasil em 1952 e é formada por animais resistentes, que sobrevivem em locais secos e com pouco pasto sem perder peso. Por causa disso, são criados principalmente em regiões nordestinas. São bois pequenos, com chifres curtos e pelo vermelho. Podem ser usados para a produção de carne ou leite."
	db.Save(&breed8)

	breed9 := Breed{}
	breed9.Breed = "Indubrasil"
	breed9.Description = "A raça é fruto do cruzamento de Nelore, Gir e Guzerá. Surgiu no Brasil em 1930, sendo criação de bovinicultores do Triângulo Mineiro. A pelagem pode ser branca, cinza ou vermelha e tem chifres médios. É usado como gado de corte e já foi exportado para os Estados Unidos."
	db.Save(&breed9)

	breed10 := Breed{}
	breed10.Breed = "Caracu"
	breed10.Description = "É um gado taurino português, trazido para o Brasil na época colonial, que tem pelagem amarela ou alaranjada. Segundo informações do Conselho Nacional de Pecuária de Corte, a raça é extremamente rústica, atingindo níveis de engorda mesmo em pastagens ruins. Outra vantagem da raça é ser resistente a doenças endêmicas brasileiras e a ectoparasitas. É usada como gado de corte ou de leite e também como animal de tração."
	db.Save(&breed10)

	breed11 := Breed{}
	breed11.Breed = "Charolês"
	breed11.Description = "De origem francesa, essa raça taurina é excelente para produção de carne. Informações do Conselho Nacional de Pecuária de Corte indicam que, no Brasil, é também muito usada na criação de mestiços, como o gado Canchim. A raça possui pelagem branca ou creme, com narinas rosas e é uma das melhores para engorda em confinamento, porque chega a atingir, em machos adultos, mais de uma tonelada."
	db.Save(&breed11)

	purpose1 := Purpose{}
	purpose1.Purpose = "Genética"
	purpose1.Description = "Animal destinado a gerar descendentes.'"
	db.Save(&purpose1)

	purpose2 := Purpose{}
	purpose2.Purpose = "Leite"
	purpose2.Description = "Animal destinado a produção de leite."
	db.Save(&purpose2)

	purpose3 := Purpose{}
	purpose3.Purpose = "Engorda"
	purpose3.Description = "Animal destinado a engorda para produção de carne."
	db.Save(&purpose3)
}
