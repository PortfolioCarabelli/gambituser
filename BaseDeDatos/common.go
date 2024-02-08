package BaseDeDatos

import (
	"database/sql"
	"os"
	"fmt"
	"github.com/PortfolioCarabelli/gambituser/models"
	"github.com/PortfolioCarabelli/gambituser/secretManager"
	_ "github.com/denisenkom/go-mssqldb"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretManager.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("sqlserver", ConnStr(SecretModel))
	if err != nil {
        fmt.Println("Error abriendo la conexión:", err.Error())
		return err
    }

    // Intentar conectar
    err = Db.Ping()
    if err != nil {
        fmt.Println("Error conectando a la base de datos:", err.Error())
		return err
    }

    fmt.Println("Conexión exitosa a la base de datos SQL Server")
	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
    dbUser := clave.UserName
    authToken := clave.Password
    dbEndpoint := clave.Host
    dbName := "gambit"

    // Construye la cadena de conexión
    connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
        dbEndpoint, dbUser, authToken, dbName)

    return connStr
}
