package erctech

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CategoriaPesquisa struct {
	Id   int
	Nome string
}

type Categoria struct {
	Id   int    `json:"Id"`
	Nome string `json:"Nome"`
}

func PesquisarCategorias(aPesquisaCategoria CategoriaPesquisa) (*[]Categoria, *erctech.Erro) {
	var resultado []Categoria
	db := erctech.GetDB()

	var xFiltros []string

	/*	if len(aPesquisaPlano.Nome) > 0 {
			xFiltros = append(xFiltros, "Nome LIKE '%"+aPesquisaPlano.Nome+"%'")
		}

		if aPesquisaPlano.Status > -1 {
			xFiltros = append(xFiltros, "Status="+strconv.Itoa(aPesquisaPlano.Status))
		}

		if aPesquisaPlano.CodPlanoGrupo > 0 {
			xFiltros = append(xFiltros, "codgrupoplano="+strconv.Itoa(aPesquisaPlano.CodPlanoGrupo))
		}
	*/
	var vWHERE string
	/*
		if len(xFiltros) > 0 {
			vWHERE = " WHERE " + xFiltros[0]
		}
		for i := 1; i < len(xFiltros); i++ {
			vWHERE = vWHERE + " AND " + xFiltros[i]
		}
	*/
	rows, err := db.Query("select id, nom from cadcategoria " + vWHERE)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		vCategoria := Categoria{}
		err = rows.Scan(&vCategoria.Id, &vCategoria.Nome)

		if err != nil {
			panic(err.Error())
		}

		resultado = append(resultado, vCategoria)
	}

	return &resultado, nil
}

func RestCategoriaPesquisar(w http.ResponseWriter, r *http.Request) {
	vPesquisa := CategoriaPesquisa{}
	vPesquisa.Nome = r.FormValue("Nome")
	vPesquisa.Id, _ = strconv.Atoi(r.FormValue("Id"))

	vCategorias, err := erctech.PesquisarCategorias(vPesquisa)
	if err != nil {

		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(vCategorias)
	}
}
