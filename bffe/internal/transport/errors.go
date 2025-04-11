package handlererrors

import "github.com/gofiber/fiber/v3/log"

func BindRequestToDtoError(handler string) { log.Error("failed to bind request to dto: %s", handler) }
func ValidatorError(handler string)        { log.Error("failed dto validation: %s", handler) }
