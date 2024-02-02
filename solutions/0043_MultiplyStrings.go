package solutions

import "strings"

func Multiply(num1 string, num2 string) string {
	result := make([]int, len(num1)+len(num2)+1)

	for i := 0; i < len(num1); i++ {
		zeroNum := len(num1) - i - 1
		tmp := make([]int, zeroNum, zeroNum+len(num2)+1)
		digit := int(num1[i] - '0')
		carry := 0
		for j := len(num2) - 1; j >= 0; j-- {
			p := digit*int(num2[j]-'0') + carry
			tmp = append(tmp, p%10)
			carry = p / 10
		}
		if carry != 0 {
			tmp = append(tmp, carry)
		}

		i := 0
		carry = 0
		for ; i < len(tmp); i++ {
			s := result[i] + tmp[i] + carry
			result[i] = s % 10
			carry = s / 10
		}
		for carry != 0 {
			s := result[i] + carry
			result[i] = s % 10
			carry = s / 10
			i++
		}
	}

	sb := strings.Builder{}
	i := len(result) - 1
	for i >= 0 && result[i] == 0 {
		i--
	}
	for i >= 0 {
		sb.WriteByte(byte(result[i] + '0'))
		i--
	}
	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}
