package tests

import (
	_ "net/http"
	_ "net/http/httptest"
	"testing"
	_ "testing"

	"github.com/labstack/echo"
	_ "github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	e := echo.New()

}
