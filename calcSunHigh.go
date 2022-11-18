// 基于北京时间和当地经纬度计算当地太阳高度角
//  参数说明:
//     latitude: 纬度，北纬为正 南纬为负
//      longitude: 经度，东经为正 西经为负
//     pekingTime: 计算的时间(北京时间)，未传入该参数则基于当前时间计算
//  返回值: 太阳高度角(度)
//  计算公式1(计算太阳高度角): sinh=sinφsinδ+cosφcosδcost (φ纬度 δ:太阳赤纬 t:地方时)
//  计算公式2(计算赤纬角):    sinCWJ= 0.39795*cos(0.98563*(N-173)/180*pi) (N:积日)
func calcSunHighAngle(latitude float64, longitude float64, pekingTime ...time.Time) float64 {
	var calcTime time.Time
	if len(pekingTime) == 0 {
		calcTime = time.Now()
	} else {
		calcTime = pekingTime[0]
	}
	localTime := time.Unix(calcTime.Unix()+int64((longitude-120.0)*60*60/15.0), 0) // 基于经度和北京时间计算当地时间
	localHour := localTime.Hour()
	localMinute := localTime.Minute()
	overHour := float64(localHour-12) + float64(localMinute)/60 // 当前时间相对于12点偏差小时数
	loaclTimeAngle := (overHour * 15) / 180 * math.Pi           // 计算当前地方时

	// 计算太阳赤纬角(当天阳光直射的维度)
	// 当地年1月1日unix时间戳
	yearFirstDay := time.Date(localTime.Year(), time.January, 1, 0, 0, 0, 0, time.Local).Unix()
	N := (localTime.Unix()-yearFirstDay)/86400 + 1                   // 计算积日
	sinCWJ := 0.39795 * math.Cos(0.98563*float64(N-173)/180*math.Pi) // 赤纬角的正弦值
	CWJ := math.Asin(sinCWJ)                                         // 赤纬角(圆周角)

	// 计算太阳高度角
	sinh := math.Sin(latitude/180*math.Pi)*math.Sin(CWJ) + math.Cos(latitude/180*math.Pi)*math.Cos(CWJ)*math.Cos(loaclTimeAngle)
	return math.Asin(sinh * 180 / math.Pi)
}
