package users

import (
	"fmt"
	"time"

	"github.com/DanielGRuiz/godesdeo/modelos"
)

func AgregarUsuario() {
	usuario := new(modelos.User)

	usuario.AddUser(1, "Daniel", time.Now(), true)

	fmt.Println(usuario)
}
