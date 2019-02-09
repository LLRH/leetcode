package main

import "log"

func main() {
	s := "PAYPALISHIRING"

	convert(s, 4)

	convertV2(s, 4)
}

func convert(s string, numRows int) string {
	status := numRows
	zigzags := make([]string, 0, 0)
	for len(s) > 0 {
		log.Printf("zigzags: %v", zigzags)
		if status == numRows {
			temp := ""
			if len(s) > numRows {
				temp = s[:numRows]
				s = s[numRows:]
			} else {
				temp = s
				for len(temp) < numRows {
					temp = temp + "_"
				}
				s = ""
			}
			zigzags = append(zigzags, temp)
			status--
		} else if status > 1 {
			temp := ""
			for i := 0; i < numRows; i++ {
				if i == status-1 {
					temp = temp + string(s[0])
				} else {
					temp = temp + "_"
				}
			}
			zigzags = append(zigzags, temp)
			s = s[1:]
			status --
		} else {
			//reset the status to origin
			status = numRows
		}
	}

	res := ""
	for row := 0; row < numRows; row++ {
		temp := ""
		for _, v := range zigzags {
			if v[row] != '_' {
				temp = temp + string(v[row])
			}
		}
		res = res + temp
	}

	log.Printf("zigzags:%v res:%v", zigzags, res)
	return res
}

func convertV2(s string, numRows int) string {
	totalLen := len(s)
	unitLen := 2 * (numRows - 1)
	res := ""
	for row := 0; row < numRows; row++ {
		if row == 0 || row == numRows-1 {
			idx := row
			for idx < totalLen {
				res = res + string(s[idx])
				idx += unitLen
			}
		} else {
			idx := row
			for idx < totalLen {
				res = res + string(s[idx])

				idx2 := idx + 2*(numRows-1-row)
				if idx2 < totalLen {
					res = res + string(s[idx2])
				}

				idx += unitLen
			}
		}
	}
	log.Printf("[convertV2] res:%v", res)
	return ""
}
