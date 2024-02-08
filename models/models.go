package models

type SecretRDSJson struct {
	UserName            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Prot                int    `json:"prot"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SingUp struct {
	UserEmail string `json: "UserMail"`
	UserUUID string `json: "UserUUID"`
}
