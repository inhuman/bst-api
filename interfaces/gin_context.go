package interfaces

type GinContext interface {
	AbortWithStatus(code int)
	AbortWithStatusJSON(code int, jsonObj interface{})
	Query(key string) string
	BindJSON(obj interface{}) error
}
