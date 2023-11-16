package main

import(
	"fmt"
	// "strings"
	"strconv"
)

func getScoreDifference(numSeq []int32) int32 {
    // Write your code here
	var firstScore, secondScore int32 = 0, 0
	isPlayerOne := true

	for len(numSeq) > 0{
		poppedNumber := numSeq[0]
		numSeq = numSeq[1:]

		
		if isPlayerOne{
			firstScore += poppedNumber
		}else{
			secondScore += poppedNumber
		}

		if poppedNumber % 2 == 0 {
			reverseSlice(numSeq)
		}

		isPlayerOne = !isPlayerOne
	}

	// for i := 0; i < len(numSeq) - 1; i++ {
	// 	firstScore += numSeq[i]

	// 	if numSeq[i] % 2 == 0 {
	// 		reverseSlice(numSeq)
	// 	}

	// 	secondScore += numSeq[i + 1]

	// 	if numSeq[i + 1] % 2 == 0 {
	// 		reverseSlice(numSeq)
	// 	}
	// }

	fmt.Println("first score:", firstScore, "second score:", secondScore)

	return firstScore - secondScore
}

func reverseSlice(slice []int32){
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]  //reverse the slice
	 }
}

func delete_at_index(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func superReducedString(s string) string {
	srs := "Empty String"
	stringBuilder := make([]bool, len(s))

	switch len(s) {
		case 1:
			return s
		case 2:
			if s[0] == s[1] {
				return "Empty String"
			}else{
				return s
			}
		default:
			for i := 1; i < len(s); i++ {
				if s[i] == s[i - 1]  {
					s = s[:i] + s[:i + 1]
				}
			}		
	}

	for i, isCharDeleted := range stringBuilder {
		if !isCharDeleted {
			srs += string(s[i])
		}
	}
	fmt.Println("string builder:", stringBuilder)

	
	return srs
}
	
	
func gemstones(rocks []string) int{
	commonMinerals := make(map[rune]int)

	for _, c := range rocks[0] {
		commonMinerals[c] = 0
	}

	for _, rock := range rocks {
		intersection := make(map[rune]int)

		for _, mineral := range rock {
			if _, found := commonMinerals[mineral]; found {
				intersection[mineral] = 0
			}
		}

		commonMinerals = intersection
	}
 
	for key := range commonMinerals {
		fmt.Println("char:", string(key))
	}

	return len(commonMinerals)
}


func minimumNumber(n int32, password string) int32 {
	numCorrections := 0
	specialCharacters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+"}
	criteria := map[string]bool{
		"special": false,
		"lower": false,
		"upper": false,
		"number": false,
	}

	for _, c := range password{

		//Check if c is uppercase
		if c >= 65 && c <= 90{
			criteria["upper"] = true
		}

		//Check if c is lowercase
		if c >= 97 && c <= 122 {
			criteria["lower"] = true
		}

		//Check for special character
		for _, specialChar := range specialCharacters {
			if string(c) == specialChar{
				criteria["special"] = true
				break
			}
		}

		//Check for number
		if _, err := strconv.Atoi(string(c)); err == nil{
			criteria["number"] = true
		}
	}

	for _, value := range criteria{
		if !value{
			numCorrections++
		}
	}

	//If after noting which characters the password needs to add it still comes out to less than 6, add the difference
	if numCorrections + len(password) < 6{
		numCorrections += (6 - (numCorrections + len(password)))
	}

	fmt.Println("criteria", criteria)

	return int32(numCorrections)
}

func marsExploration(s string) int32 {
    //numSOS := len(s) / 3
    numChangedSOS := 0
    
	fmt.Println("len:", len(s))
    for i := 0; i < len(s); i += 3{
		slice := s[i : i + 3]
        if slice != "SOS"{
            numChangedSOS++   
        }   
		
		fmt.Println(s[i : i + 3])
    }
    
    return int32(numChangedSOS)
}


