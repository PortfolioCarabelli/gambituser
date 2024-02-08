package BaseDeDatos

import (
	"fmt"

	"github.com/PortfolioCarabelli/gambituser/models"
	"github.com/PortfolioCarabelli/gambituser/tools"
	_ "github.com/denisenkom/go-mssqldb"
)

func SingUp(sign models.SingUp) error {
	fmt.Println("Comienza Registro")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	query := fmt.Sprintf(`
    INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%v' , '%v', '%v')`,
		sign.UserEmail,
		sign.UserUUID,
		tools.FechaSQLServer(),
	)

	fmt.Println(query)

	_, err = Db.Exec(query)

	if err != nil {
		fmt.Println("Hubo un error al ejecutar la Query " + err.Error())
	}

	fmt.Println("SingUp > Ejecucion Exitosa")
	return nil
}
