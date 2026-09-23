package config

import (
	"fmt"
	"os"
	"strconv"
)

func LoadSingleEnvVar[T string | int | bool ](key string,defaultval T) T {
	val , exists := os.LookupEnv(key)
	
	if !exists{
		return defaultval
	}

	switch any(defaultval).(type){
	case string:
		return any(val).(T)
	
	case int:
		i,err := strconv.Atoi(val)

		if err!= nil{
			fmt.Println("something went wrong while parsing this new var" + err.Error())
			return defaultval
		}
		return  any(i).(T)
	case float32:
		f,err := strconv.ParseFloat(val,32)
		if err!= nil{
			fmt.Println("something went wrong while parsing this new var" + err.Error())
			return defaultval
		}
		return  any(float32(f)).(T)
	
	case bool:
		b,err := strconv.ParseBool(val)
		if err!= nil{
			fmt.Println("something went wrong while parsing this new var" + err.Error())
			return defaultval
		}
		return  any(b).(T)
	
	default:
		fmt.Println("Unsupported datatype has been passed for the env var:")
		return defaultval
	}



}