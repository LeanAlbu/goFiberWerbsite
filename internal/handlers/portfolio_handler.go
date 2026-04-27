package handlers

import "github.com/gofiber/fiber/v2"

type Projeto struct{
	Nome string
	Linguagem string
	Descricao string
}

func RenderHome(c *fiber.Ctx) error{
	myProjects := []Projeto{
		{
			Nome:      "PortScanner",
			Linguagem: "Rust",
			Descricao: "Scanner de portas de alta performance com algoritmos otimizados para varredura de rede.",
		},
		{
			Nome:      "LogAnalyser",
			Linguagem: "Java",
			Descricao: "Ferramenta de análise de logs estruturados focada em segurança e rastreabilidade.",
		},
		{
			Nome:      "BICentral API",
			Linguagem: "Go / Spring Boot",
			Descricao: "Arquitetura backend para sistema de gestão universitária com autenticação JWT.",
		},
	}
	return c.Render("index", fiber.Map{
		"Projetos": myProjects,
	})

}
