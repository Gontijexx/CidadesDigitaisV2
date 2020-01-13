package middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"CidadesDigitaisV2/api/auth"
	"CidadesDigitaisV2/api/responses"

	"github.com/gorilla/mux"
)

func SetMiddleJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddleAuthMod(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
			return
		}
		vars := mux.Vars(r)
		PagMod, err := strconv.ParseFloat(vars["modulo"], 64)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		mod, err := auth.ExtractTokenMod(r)
		umod := InterfaceSlice(mod)
		fmt.Printf("eu sou umod %v", umod)
		umodInt := make([]float64, len(umod))
		for i := range umod {
			umodInt[i] = umod[i].(float64)
		}
		fmt.Printf("eu sou umodit %v", umodInt)
		for _, v := range umodInt {
			fmt.Printf("eu sou v %v", v)
			if v == PagMod {
				next(w, r)
				return
			}else {
				responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL Unauthorized"))
				return
			}

		}

	}
}

func SetMiddleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
			return
		}
		next(w, r)
	}
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
