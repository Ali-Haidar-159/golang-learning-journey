package service

func CheckAge (age int) error{

	if(age > 18){
		return nil
	}else{
		return AgeError{Age:age}
	}

}