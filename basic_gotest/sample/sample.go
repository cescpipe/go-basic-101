package sample

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateThailandCitizenID(idNo string) error {

	if len(idNo) != 13 {
		return errors.New("identification no is not 13 digits")
	}

	_, err := strconv.Atoi(idNo)
	if err != nil {
		return errors.New("identification no must be numeric only")
	}

	idNoSp := idNo[:12]
	//fmt.Println("idNoSp :", idNoSp)

	sum := 0

	for i := 0; i < 12; i++ {
		nd, err := strconv.Atoi(idNoSp[i : i+1])
		if err != nil {
			return err
		}
		//fmt.Println("num digit :", nd , " calculate by : ", nd * (13 - i))
		sum += nd * (13 - i)
	}
	//fmt.Println("total summary of num :", sum)
	//fmt.Println("total summary of num with mod :", sum % 11)

	ld := sum % 11

	if ld > 9 {
		ld = ld - 10
	}

	calIdno := fmt.Sprintf("%s%d", idNo[:12], ld)

	//fmt.Println("calculate identification id :", calIdno)
	//fmt.Println("input  identification id :", idNo)

	if strings.Compare(calIdno, idNo) != 0 {
		return errors.New("identification id mismatch")
	}

	return nil
}
