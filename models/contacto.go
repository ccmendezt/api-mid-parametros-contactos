package models

type Contacto struct {
	Id                int
	Nombre            string
	Documento         int64
	Direccion         string
	FechaCreacion     string
	FechaModificacion string
	Activo            bool
}
