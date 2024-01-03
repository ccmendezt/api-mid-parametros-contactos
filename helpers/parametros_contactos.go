package helpers

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/udistrital/api_mid_parametros_contactos/models"
)

func ListarParametrosContactos() (parametros_Contactos models.ParametrosContactos, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListarParametros", "err": err, "status": "500"}
		}
	}()
	var parametros_query []models.Parametro
	var contactos_query []models.Contacto
	//Limit 0 para traer todos los registros, si no se pone limit trae solo 10
	url := "parametro?query=TipoParametroId__Nombre:Cargos&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros_query); err != nil {
		log.Error(err)
		panic(err.Error())
	}
	url = "contacto"
	if err := GetRequestNew("UrlCrudAgenda", url, &contactos_query); err != nil {
		log.Error(err)
		panic(err.Error())
	}
	fmt.Println("Contactos: ", contactos_query)
	parametros_Contactos.Parametros = parametros_query
	parametros_Contactos.Contactos = contactos_query
	return parametros_Contactos, outputError
}
