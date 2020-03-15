package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	// 得到当前时间。
	now := time.Now()
	p(now)
	// 通过提供年月日等信息，你可以构建一个 time。时间总是关联着位置信息，例如时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 你可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// 输出是星期一到日的 Weekday 也是支持的。
	p(then.Weekday())
	// 这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)
	// 我们计算出不同单位下的时间长度值。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// 你可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔。
	p(then.Add(diff))
	p(then.Add(-diff))

	p("################")
	// 格式化
	// 这里是一个基本的按照 RFC3339 进行格式化的例子，使用对应模式常量。
	t := time.Now()
	p(t.Format(time.RFC3339))
	// 时间解析使用同 Format 相同的形式值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	// ormat 和 Parse 使用基于例子的形式来决定日期格式，
	// 一般你只要使用 time 包中提供的模式常量就行了，但是你也可以实现自定义模式。
	// 模式必须使用时间 Mon Jan 2 15:04:05 MST 2006来指定给定时间/字符串的格式化/解析方式。
	// 时间一定要按照如下所示：2006为年，15 为小时，Monday 代表星期几，等等。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	// 对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出出时间值得组成。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// Parse 函数在输入的时间格式不正确是会返回一个错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
