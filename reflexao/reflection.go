package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		switch campo.Kind() {
		case reflect.String:
			fn(campo.String())
		case reflect.Struct:
			percorre(campo.Interface(), fn)
		}
	}
}
