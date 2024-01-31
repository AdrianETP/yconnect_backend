package controllers

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

// funcion para crear una organizacion
func CreateOrganization(c *fiber.Ctx) error {
	// PENDIENTE: Cambiar de parsear directo a parsear , a parsar a una vaiable body que tenga el atributo organziation

	// creacion de una organizacion en base al modelo
	var Org models.Organization
	// se le agrega un id de tipo id
	Org.ID = primitive.NewObjectID()
	// se le agrega a esta variable todo lo que se pueda encontrar en la request body
	err := c.BodyParser(&Org)
	// si no se pudo efectuar
	if err != nil {
		// se regresa un error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	// insertar a la organizacion a la base de datos
	result, error := config.Database.Collection("Organizations").InsertOne(context.TODO(), Org)

	// si hubo un error
	if error != nil {
		// se regresa el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  error.Error(),
		})
	}

	// se regresa el resultado de la creacion de la base de datos
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

// obtener la organizacion por tag
func GetOrgByTag(c *fiber.Ctx) error {
	// variable tags (puramente para parsear el body)
	var Tag struct {
		Tags []string `json:tags`
	}
	// parseamos el body
	c.BodyParser(&Tag)
	// sacamos el array de tags
	tags := Tag.Tags

	// coleccion de organizaciones
	orgCol := config.Database.Collection("Organizations")

	// variable donde vamos a guardar los resultados
	var organizations []models.Organization
	// iteramos por cada tag
	for _, t := range tags {
		// hacemos un query de todas las arganizaciones  que tengan el tag iterado
		results, _ := orgCol.Find(context.TODO(), bson.D{
			{"tags", bson.D{{"$elemMatch", bson.D{{"$eq", t}}}}},
		})

		// por cada organizacion encontrada
		for results.Next(context.TODO()) {
			// agregamos la organizacion a una variable temporal
			var organization models.Organization
			results.Decode(&organization)
			// variable para revisar si ya habiamos agregado esa organizacion a nuestros resultados
			inOrg := false
			// si el largo de la variable organizations es 0
			if len(organizations) <= 0 {
				// agregamos la organizacion a fuerza (el codigo de abajo no podria correr por que estariamos iterando en una variable de largo 0)
				organizations = append(organizations, organization)
			} else {
				// iteramos por todas las organizaciones que hemos agregado (para revisar que no vayamos a meter una repetida)
				for i := 0; i < len(organizations); i++ {
					// si ya la tenemos agregada
					if organization.ID.String() == organizations[i].ID.String() {
						// ponemos esto como true para que ya no la agregue
						inOrg = true
					}
				}
				// si no la hemos agregado
				if !inOrg {
					// la agregamos
					organizations = append(organizations, organization)
				}
			}

		}
	}
	//  regresamos las organizaciones
	return c.JSON(fiber.Map{
		"status":        200,
		"organizations": organizations,
	})
}

func ChangeGrade(c *fiber.Ctx) error {
	var body struct {
		Grade string
		OrgId primitive.ObjectID
	}
	c.BodyParser(&body)

	res := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var Organization models.Organization
	res.Decode(&Organization)

	if Organization.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "organization doesn't exist",
		})
	}

	result, err := config.Database.Collection("Organizations").
		UpdateOne(context.TODO(), bson.D{{"_id", Organization.ID}}, bson.D{
			{"$set", bson.D{
				{"grade", Organization.Grade},
			}},
		})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"result": result,
	})
}

// funcion para obtener todas las organizaciones
func GetAllOrgs(c *fiber.Ctx) error {
	var body struct {
		Token string `json:token`
	}
	c.BodyParser(&body)
	token, tokenError := validateToken(body.Token)
	if token == "" {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  tokenError,
		})
	}

	// vamos a guardar los usuarios decodificados aqui
	var organizations []models.Organization
	// aqui vamos a llamar a mongo y decirle que encuentre a usuarios pero sin filtro ( osea que saque a todos los usuarios)
	results, err := config.Database.Collection("Organizations").Find(context.TODO(), bson.M{})
	// si da un error
	if err != nil {
		// regresamos el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}
	// aca vamos a iterar por todos los resultados y decodificarlos
	for results.Next(context.TODO()) {
		// creamos una variable temporal
		var organnization models.Organization
		results.Decode(&organnization)
		// agremaos esa organizacion al resultado final
		organizations = append(organizations, organnization)
	}
	// regresamos a las organizaciones como json
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

// obtener organizaciones en base a favoritos
func GetFavorites(c *fiber.Ctx) error {
	// hacemos una variable que simule el body y solo pida el user id
	var body struct {
		UserID primitive.ObjectID `json:userId`
	}
	// parseamos el body a esta variable
	c.BodyParser(&body)

	// sacamos al usuario de la db
	results := config.Database.Collection("Users").
		FindOne(context.TODO(), bson.D{{"_id", body.UserID}})
		//  lo agregamos a una variable de tipo usuario
	var user models.User
	results.Decode(&user)
	// hacemos una variable donde guardemos a todas las organizaciones favoritas
	var organizations []models.Organization
	// por cada favorita (favoritos siendo el id de las organizaciones)
	for _, v := range user.Favorites {
		//  buscamos la organizacion
		r, err := config.Database.Collection("Organizations").
			Find(context.TODO(), bson.D{{"_id", v}})
			// si da un error
		if err != nil {
			// regresamos el error
			return c.JSON(fiber.Map{
				"status": 400,
				"error":  err.Error(),
			})
		}

		// por cada organizacion encontrada
		for r.Next(context.TODO()) {
			// creamos una variable temporal
			var org models.Organization

			// agregamos el resultado a la variable temporal
			r.Decode(&org)
			// hacemos apend de esta organizacion al resultado
			organizations = append(organizations, org)

		}
	}

	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

func GetOrgByName(c *fiber.Ctx) error {
	var body struct {
		Name string `json:name`
	}
	c.BodyParser(&body)

	// lo que checa este query es que el nombre de la organizacion este en cualquier parte de la string (sin importar las mayusculas)
	result, err := config.Database.Collection("Organizations").
		Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", primitive.Regex{Pattern: ".*" + body.Name + ".*", Options: "i"}}}}}) // el Options: i significa case insensitive
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	// donde va a estar el resultado
	var organizations []models.Organization

	// iteramos por cada organizacion
	for result.Next(context.TODO()) {
		// la agregamos a una variable temporal
		var organization models.Organization
		result.Decode(&organization)
		// lo agregamos a los resultados
		organizations = append(organizations, organization)
	}
	// regresamos los resultados
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

// borrar una organizacion
func DeleteOrg(c *fiber.Ctx) error {
	// creamos una variable body para parsear el body de la request
	var body struct {
		OrgId primitive.ObjectID `json:orgId`
	}
	// parseamos el body en esa variable
	c.BodyParser(&body)

	// borramos la organizacion en la base de datos
	result, err := config.Database.Collection("Organizations").DeleteOne(context.TODO(), bson.D{
		{"_id", body.OrgId},
	})
	// si hay un error
	if err != nil {
		// regresa el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	// regresa el resultado
	return c.JSON(fiber.Map{
		"data":   result,
		"status": 200,
	})
}

// obtiene la organizacion por id
func GetOrgById(c *fiber.Ctx) error {
	// creamos una variable body para parsear el body de la request
	var body struct {
		OrgId primitive.ObjectID `json:orgid`
	}

	// parseamos el body de la request
	c.BodyParser(&body)

	// buscamos una organizacion con ese id
	result := config.Database.Collection("Organizations").
		FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

		// creamos una variable para el resultado
	var organization models.Organization

	result.Decode(&organization)
	// enviamos el resultado
	return c.JSON(fiber.Map{
		"status":       200,
		"organization": organization,
	})
}

// modificar una organizacion
func ModifyOrg(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		Organization models.Organization `json:organization`
	}
	// parseamos a body
	c.BodyParser(&body)

	// query para modificar la organizacion (la encuentra en base al id)
	result, err := config.Database.Collection("Organizations").
		UpdateOne(context.TODO(), bson.D{{"_id", body.Organization.ID}}, bson.D{
			{"$set", bson.D{
				{"name", body.Organization.Name},
				{"location", body.Organization.Location},
				{"telephone", body.Organization.Telephone},
				{"tags", body.Organization.Tags},
				{"igurrl", body.Organization.IgUrl},
				{"description", body.Organization.Description},
				{"email", body.Organization.Email},
				{"grade", body.Organization.Grade},
			}},
		})
		// si da un error
	if err != nil {
		// regresa el error
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": 400,
		})
	}

	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

// funcion para mandar un email de una organizacion que se quiera registrar
func SendMail(c *fiber.Ctx) error {
	// variable para parsear el body
	var body models.Organization

	// parseamos el body
	c.BodyParser(&body)

	// creamos un nuevo mensaje
	m := gomail.NewMessage()

	// agarramos todos los datos del .env
	adminMail := os.Getenv("ADMINMAIL")
	recieverMail := os.Getenv("RECIEVERMAIL")
	adminPass := os.Getenv("ADMINPASS")

	// Set E-Mail sender
	m.SetHeader("From", adminMail)

	// Set E-Mail receivers
	m.SetHeader("To", recieverMail)

	// Set E-Mail subject
	m.SetHeader("Subject", "nueva organizacion")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", `<h1>new organization</h1>
        <p>una nueva organizacion ah solicitado  registrarse, estos son sus datos</p>
        <ul>
            <li><b>name: </b>`+body.Name+`</li>
            <li><b>description: </b>`+body.Description+`</li>
            <li><b>telephone: </b>`+body.Telephone+`</li>
            <li><b>location: </b>`+body.Location+`</li>
            <li><b>email: </b>`+body.Email+`</li>
            <li><b>igtag: </b>`+body.Igtag+`</li>
        </ul>`)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, adminMail, adminPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   "email sent to admin",
	})
}
