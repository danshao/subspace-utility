package validator

import "testing"

func TestAtLeastOne(t *testing.T) {
	case1 := "AAAA"
	if !AtLeastOne(case1, UPPERCASE) {
		HandleError(t, case1)
	}

	case2 := "BBBB"
	if AtLeastOne(case2, LOWERCASE) {
		HandleError(t, case1)
	}

	case3 := "cccc"
	if !AtLeastOne(case3, LOWERCASE) {
		HandleError(t, case1)
	}

	case4 := "1111"
	if !AtLeastOne(case4, DIGIT) {
		HandleError(t, case1)
	}
}

func TestOnlyContains(t *testing.T) {
	case1 := "AAA"
	if !OnlyContains(case1, 3, UPPERCASE) {
		HandleError(t, case1)
	}

	case2 := "[[[]]]"
	if !OnlyContains(case2, 3, SPECIAL_CHARACTER) {
		HandleError(t, case2)
	}
}

func TestValidatePassword(t *testing.T) {
	// Positive case
	positiveCase1 := "Abcd1234"
	if !IsValidPassword(positiveCase1) {
		HandleError(t, positiveCase1)
	}
	
	positiveCase2 := "Abcd1234[]"
	if !IsValidPassword(positiveCase2) {
		HandleError(t, positiveCase2)
	}

	// Negative case
	negativeCase1 := "aaaaaaaa"
	if IsValidPassword(negativeCase1) {
		HandleError(t, negativeCase1)
	}

	negativeCase2 := "AAAAAAAA"
	if IsValidPassword(negativeCase2) {
		HandleError(t, negativeCase2)
	}

	negativeCase3 := "AAAAoooo"
	if IsValidPassword(negativeCase3) {
		HandleError(t, negativeCase3)
	}

	negativeCase4 := "22222222"
	if IsValidPassword(negativeCase4) {
		HandleError(t, negativeCase4)
	}

	negativeCase5 := "Aa4"
	if IsValidPassword(negativeCase5) {
		HandleError(t, negativeCase5)
	}

	negativeCase6 := "Abc中文123"
	if IsValidPassword(negativeCase6) {
		HandleError(t, negativeCase6)
	}
}

func HandleError(t *testing.T, str string) {
	t.Error(str, "not validate.")
}