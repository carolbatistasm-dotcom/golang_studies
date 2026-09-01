package main

func main() {

	notas := map[string][]float64{} // uma fatia de floats para cada aluno

	notas["carol"] = []float64{9.0, 8.5, 7.0}
	notas["joao"] = []float64{6.0, 7.5, 8.0}
	notas["maria"] = []float64{10.0, 9.5, 8.5}

	cursos := map[string][]string{
		"carol": {"Matemática", "Física", "Química"},
		"joao":  {"História", "Geografia", "Sociologia"},
		"maria": {"Biologia", "Química", "Física"},
	} // uma fatia de strings para cada aluno

	cursos["teo"] = []string{"Inglês", "Espanhol", "Francês"} // adicionando um novo aluno ao mapa de cursos

	// acessando as notas de um aluno
	notasCarol := notas["carol"]
	println("Notas da Carol:", notasCarol[0], notasCarol[1], notasCarol[2])

	// acessando os cursos de um aluno
	cursosJoao := cursos["joao"]
	println("Cursos do João:", cursosJoao[0], cursosJoao[1], cursosJoao[2])
}
