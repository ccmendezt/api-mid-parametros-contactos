package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/api_mid_parametros_contactos/helpers"
)

type ParametrosContactosController struct {
	beego.Controller
}

func (c *ParametrosContactosController) URLMapping() {
}

// GetAll ...
// @Title Get All
// @Description get ParametrosContactos
// @Success 200 {object} models.ParametrosContactos
// @Failure 403
// @router / [get]
func (c *ParametrosContactosController) GetAll() {

	defer helpers.ErrorController(c.Controller, "ParametrosController")
	if v, err := helpers.ListarParametrosContactos(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Listado consultado con éxito", "Data": v}
	} else {
		panic(err)
	}
	c.ServeJSON()
}
