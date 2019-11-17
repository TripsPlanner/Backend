package service

import (
	"gomod/dao"
	"log"
)

type Ret struct {
	X           int
	Y           int
	V           int
	N           string
	T           string
	CoordinateX float64
	CoordinateY float64
}

func RecommandLine(target string, span int) ([]Ret, error) {
	var curSpan int
	data, err := GetJourneyList(target)
	if err != nil {
		return nil, err
	}

	if len(*data) == 0 {
		return nil, nil
	}

	jourData := *data
	spanLen := span //4个小时为一个最小单位
	log.Printf("24 span is :%v", spanLen)
	//var outspaned int

	//curSpan = spanLen - outspaned
	log.Printf("curspan is %v", curSpan)
	//var dp [len(jourData)][spanLen]int

	//dp := make([len(jourData)][spanLen]int)
	log.Printf("%v", spanLen)
	log.Printf("%v", len(jourData))
	//dp := make([][]int, len(jourData), spanLen)
	//dp := make([][]int, 0, 0)

	//var dp [6][spanLen]int
	//return nil, nil
	var res []Ret

	log.Printf("jou is :%v", jourData)
	// for j := 1; j < spanLen+1; j++ {
	// 	re := Ret{
	// 		X: 0,
	// 		Y: j,
	// 		V: jourData[0].Favorite,
	// 		N: jourData[0].Destination,
	// 	}

	// 	res = append(res, re)
	// }

	for i := 0; i < len(jourData); i++ {
		for j := 1; j < spanLen+1; j++ {
			//var re Ret

			if jourData[i].Stime/4 > j {
				var re Ret

				log.Printf("gl ,so abondon%+v", jourData[i])
				var v2 int
				var n2 string
				var t2 string

				if i == 0 {
					re = Ret{
						X: 0,
						Y: j,
						V: 0,
						N: jourData[0].Destination,
						T: jourData[0].Destination,
					}
					res = append(res, re)
					continue
				}

				for _, v := range res {
					if v.X == i-1 && v.Y == j {
						v2 = v.V
						n2 = v.N
						t2 = v.T
					}
					re = Ret{
						X: i,
						Y: j,
						V: v2,
						N: n2,
						T: t2,
					}
				}

				// re = Ret{
				// 	X: i,
				// 	Y: j,
				// 	// V: jourData[i-1].Favorite,

				// 	N: jourData[i-1].Destination,
				// }
				//}
				res = append(res, re)
				continue
			}

			var v1 int
			var y1 int
			var t1 string
			for _, v := range res {
				if v.X == i-1 && v.Y == j {
					v1 = v.V

				}

				if v.X == i-1 && v.Y == j-jourData[i].Stime/4 {
					y1 = v.V
					t1 = v.T
				}
			}

			//加入当前商品，总权重更小
			if v1 > jourData[i].Favorite+y1 {
				log.Printf("111111%v", jourData[i].ID)
				var re Ret
				var v2 int
				var n2 string
				var t2 string
				for _, v := range res {
					if v.X == i-1 && v.Y == j {
						v2 = v.V
						n2 = v.N
						t2 = v.T
					}
					re = Ret{
						X: i,
						Y: j,
						V: v2,
						N: n2,
						T: t2,
					}
				}

				res = append(res, re)
			} else {
				log.Printf("y1 is :%v", y1)
				re := Ret{
					X: i,
					Y: j,
					V: jourData[i].Favorite + y1,
					N: jourData[i].Destination,
					T: jourData[i].Destination + ";" + t1,
				}

				res = append(res, re)
			}

		}
	}

	log.Printf("%v", res)

	log.Printf("res is new%v", res)
	targetLine := FindTarget(res, span, *data)

	targetLineNew := make([]Ret, 0)

	datalist, _ := GetJourneyList(target)
	for _, vvv := range targetLine {
		for _, vvvvv1 := range *datalist {
			if vvv.N == vvvvv1.Destination {
				vvv.CoordinateX = vvvvv1.CoordinateX
				vvv.CoordinateY = vvvvv1.CoordinateY

				targetLineNew = append(targetLineNew, vvv)
			}
		}
	}

	return targetLineNew, nil
	// return targetLine, nil

}

func GetMax() {

}

func FindTarget(res []Ret, span int, data []dao.Journey) []Ret {
	var max int
	var re Ret
	var ress []Ret

	//var target []string

	for _, v := range res {
		if v.Y != span {
			continue
		}

		if v.V >= max {
			//获得最大值
			max = v.V
			re = v
		}
	}

	log.Printf("max v is :%v", re)

	re = Ret{
		X: re.X,
		Y: data[re.X].Stime / 4,
		N: data[re.X].Destination,
	}

	ress = append(ress, re)
	log.Printf("ress ban:%v", ress)

	//var

	var tSpan int
	for tSpan <= span {
		for _, vl := range res {
			var takeSpan int
			for _, v := range ress {
				takeSpan = takeSpan + v.Y

			}

			if vl.Y == span-takeSpan && vl.V != 0 {
				ok := false
				for _, v3 := range ress {
					if vl.X == v3.X {
						ok = true
					} else {
						//ok = false
					}
				}

				if !ok {
					re = Ret{
						X: vl.X,
						Y: data[vl.X].Stime / 4,
						N: data[vl.X].Destination,
					}

					log.Printf("take span is %v", takeSpan)
					log.Printf("take span vl.y is %v", vl)
					tSpan = takeSpan
					ress = append(ress, re)
					//continue
				}

			}

		}

		var badhost []int
		for _, v2 := range data {
			ok := true
			for _, v := range ress {
				if v2.Destination == data[v.X].Destination {
					ok = true
				} else {
					ok = false
				}
			}
			if !ok {
				badhost = append(badhost, v2.Stime/4)
			}

		}
		for _, vt := range badhost {
			if span-tSpan < vt {
				return ress
			}
		}

		if len(ress) == len(data) {
			return ress
		}
		//log.Printf("ress222:%v", ress)
	}

	log.Printf("ress:%v", ress)
	return ress
}
