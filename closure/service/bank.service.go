package service 

func Bank(totalBalance int) func(int,bool)int{

	balance := totalBalance

	return func (amount int , isDeposite bool) int{

		if(isDeposite){
			balance = balance + amount
			return balance
		} else {
			balance = balance - amount
			return balance
		}

	}


}