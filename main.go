package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/PortfolioCarabelli/gambituser/awsgo"
	"github.com/PortfolioCarabelli/gambituser/BaseDeDatos"
	"github.com/PortfolioCarabelli/gambituser/models"
	
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()
	if !ValidoParameters() {
		fmt.Println("Error en los Parametros debe enviar 'SecretName'")
		err := errors.New("Error en los Parametros debe enviar 'SecretName'")
		return event, err
	}

	var datos models.SingUp

	for row, att := range event.Request.UserAttributes{
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email: " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub: " + datos.UserUUID)
		}
		
	}

	err := BaseDeDatos.ReadSecret()

	if(err != nil){
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}

	err = BaseDeDatos.SingUp(datos)
	return event,err
}


func ValidoParameters () bool {
	var traeParametro bool

	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}