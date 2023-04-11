package handler

import (
	"net/http"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/dentista"
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

// NewDentistaHandler crea un nuevo controller de dentista
func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

/*
	POST: agregar dentista.
	GET: traer dentista por ID.
	PUT: actualizar dentista.
	PATCH: actualizar un dentista por alguno de sus campos.
	DELETE: eliminar dentista.
*/

func (h *dentistaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista

		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		d, err := h.s.Create(d)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, d)
	}
}

func (h *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := h.s.ReadAll()
		c.JSON(http.StatusOK, res)
	}
}

func (h *dentistaHandler) GetByMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		matricula, err := strconv.Atoi(c.Param("matricula"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid matricula"})
			return
		}
		res, err := h.s.Read(matricula)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The specified matricula could not be found"})
			return
		}

		c.JSON(http.StatusOK, res)

	}
}

// func (h *dentistaHandler) Put() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var d domain.Dentista
// 		if err := c.ShouldBindJSON(&d); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		matricula := c.Param("matricula")

// 		err := h.s.Put(matricula, d)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		c.JSON(http.StatusOK, "Dentista updated succesfully")
// 	}
// }

// func (h *dentistaHandler) Patch() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var d domain.Dentista
// 		if err := c.ShouldBindJSON(&d); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		matricula := c.Param("matricula")

// 		err := h.s.Patch(matricula, d)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		c.JSON(http.StatusOK, "Dentista updated succesfully")
// 	}
// }

func (h *dentistaHandler) Update(functionType string) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		matricula := c.Param("matricula")

		err := h.s.Update(matricula, d, functionType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Dentista updated succesfully")
	}
}

func (h *dentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

