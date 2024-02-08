package secretManager

import (
	"encoding/json"

	"fmt"
	"github.com/PortfolioCarabelli/gambituser/awsgo"
	"github.com/PortfolioCarabelli/gambituser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson

	fmt.Println("> Pido Secreto" + nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		fmt.Println(err.Error())

		return datosSecret, err
	}
	
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura Secret OK" + nameSecret)
	return datosSecret, nil
}
