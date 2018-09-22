package talk

type Talk interface {
	Hello(userName string) string
	Talk(heard string)(saying string ,end bool, err error)
}


