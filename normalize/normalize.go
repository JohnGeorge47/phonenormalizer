package normalize

import (
	"github.com/JohnGeorge47/phonenormalizer/models"
	"fmt"
	"regexp"
	"sort"
)

func NormalizeNumber(numberStruct *[]models.NonNormalizedNo)[]string{
	formatedArray:=make([]string,len(*numberStruct))
	re:=regexp.MustCompile("^[0-9]{10}$")
	reg,_ := regexp.Compile("[^0-9]+")
    for _,num:=range *numberStruct{
    	check:=re.MatchString(num.PhoneNumber)
    	if !check{
    		formatted:=reg.ReplaceAllString(num.PhoneNumber,"")
    		if re.MatchString(formatted){
				formatedArray=append(formatedArray,formatted)
			}
		}else{
			formatedArray=append(formatedArray,num.PhoneNumber)
		}
	}
	fmt.Println(formatedArray)
	arrayToinsert:=SortandRemoveDuplicates(formatedArray)
	fmt.Println(arrayToinsert)
	return arrayToinsert
}

func SortandRemoveDuplicates(preDupString[]string)[]string{
	sort.Strings(preDupString)
	cleanedString:=Remove(preDupString)
	return cleanedString
}

func Remove(s []string)[]string{
	length:=len(s)
	tmp:=make([]string,length)
	count:=0
	for i:=0;i<length-1 ;i++  {
		if s[i]!=s[i+1]{
			tmp[count]=s[i]
			count++
		}
	}
	tmp[count+1]=s[length-1]
	return tmp
}

