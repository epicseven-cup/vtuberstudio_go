package internal

//https://stackoverflow.com/questions/47911187/golang-elegantly-json-decode-different-structures

/*
Two different idea here, that I can have a metadata wrapper structure and have json raw value to parse it correctlu
Or
I can just have a bunch of if statements and parse it
*/

func ParseByMessageType(data []byte) ([]interface{}, error) {

	return nil, nil
}
