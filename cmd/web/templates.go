package main

import "github.com/abassGarane/snippet/pkg/models"

type templateData struct{
  Snippet *models.Snippet
  Snippets []*models.Snippet
}
