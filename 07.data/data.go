package main

import "fmt"

func main() {
	// hello world
	fmt.Println("hello world")

	fmt.Println("##################### arrays")
	arrays()

	fmt.Println("##################### slice")
	slice()

	fmt.Println("##################### mapFunc")
	mapFunc()

	fmt.Println("##################### rangeFunc")
	rangeFunc()
}

// 数组
func arrays() {
	// 这里我们创建了一个数组 test1 来存放刚好 5 个 int。
	// 元素的类型和长度都是数组类型的一部分。
	// 数组默认是零值的，对于 int 数组来说也就是 0。
	var test1 [6]int
	fmt.Println("内容:", test1)
	// 我们可以使用 array[index] = value 语法来设置数组指定位置的值，或者用 array[index] 得到值。
	test1[4] = 100
	fmt.Println("设置:", test1)
	fmt.Println("获取:", test1[4])
	// 使用内置函数 len 返回数组的长度
	fmt.Println("长度:", len(test1))

	// 使用这个语法在一行内初始化一个数组
	test2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数据:", test2)

	// 数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构。
	var twoTest [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoTest[i][j] = i + j
		}
	}
	// 注意，在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt.Println("二维: ", twoTest)
}

// 切片
func slice() {
	// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口

	// 不像数组，slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。
	// 要创建一个长度非零的空slice，需要使用内建的方法 make。
	// 这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）。
	test1 := make([]string, 3)
	fmt.Println("数据:", test1)
	// 我们可以和数组一样设置和得到值
	test1[0] = "A"
	test1[1] = "C"
	test1[2] = "B"
	fmt.Println("数据:", test1)
	fmt.Println("获取:", test1[2])
	// 如你所料，len 返回 slice 的长度
	fmt.Println("长度:", len(test1))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 append，它返回一个包含了一个或者多个新值的 slice。
	// 注意我们接受返回由 append返回的新的 slice 值。
	test1 = append(test1, "D")
	test1 = append(test1, "E", "F")
	fmt.Println("追加:", test1)

	// Slice 也可以被 copy。这里我们创建一个空的和 test1 有相同长度的 slice test2，并且将 test1 复制给 test2。
	test2 := make([]string, len(test1))
	copy(test2, test1)
	fmt.Println("拷贝:", test2)
	// Slice 支持通过 slice[low:high] 语法进行“切片”操作。例如，这里得到一个包含元素 test1[2], test1[3],test1[4] 的 slice。
	l := test1[2:5]
	fmt.Println("切片1:", l)
	// 这个 slice 从 test1[0] 到（但是不包含）test1[5]。
	l = test1[:5]
	fmt.Println("切片2:", l)
	// 这个 slice 从（包含）test1[2] 到 slice 的后一个值。
	l = test1[2:]
	fmt.Println("切片3:", l)
	// 我们可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("数据:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
	twoTest := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoTest[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoTest[i][j] = i + j
		}
	}
	// 注意，slice 和数组不同，虽然它们通过 fmt.Println 输出差不多。
	fmt.Println("二维: ", twoTest)
}

// 键值对 key/value
func mapFunc() {
	// 要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type).
	map1 := make(map[string]int)
	// 使用典型的 make[key] = val 语法来设置键值对。
	map1["k1"] = 7
	map1["k2"] = 13
	// 使用例如 Println 来打印一个 map 将会输出所有的键值对。
	fmt.Println("数据:", map1)
	// 使用 name[key] 来获取一个键的值
	v1 := map1["k1"]
	fmt.Println("值: ", v1)
	// 当对一个 map 调用内建的 len 时，返回的是键值对数目
	fmt.Println("长度:", len(map1))
	// 内建的 delete 可以从一个 map 中移除键值对
	delete(map1, "k2")
	fmt.Println("数据:", map1)
	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。
	// 这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	_, prs := map1["k2"]
	fmt.Println("是否存在:", prs)
	// 你也可以通过这个语法在同一行申明和初始化一个新的map。
	map2 := map[string]int{"F": 1, "B": 2}
	// 注意一个 map 在使用 fmt.Println 打印的时候，是以 map[k:v k:v]的格式输出的。
	fmt.Println("数据:", map2)
}

// Range 遍历
func rangeFunc() {
	// 这里我们使用 range 来统计一个 slice 的元素个数。数组也可以采用这种方法。
	array1 := []int{2, 3, 4}
	sum := 0
	for _, num := range array1 {
		sum += num
	}
	fmt.Println("求和:", sum)

	// range 在数组和 slice 中都同样提供每个项的索引和值。
	// 上面我们不需要索引，所以我们使用 空值定义符_ 来忽略它。
	// 有时候我们实际上是需要这个索引的。
	for i, num := range array1 {
		if num == 3 {
			fmt.Println("索引:", i)
		}
	}

	// range 在 map 中迭代键值对。
	map1 := map[string]string{"A": "苹果", "B": "香蕉"}
	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range map1 {
		fmt.Println("键:", k)
	}

	// range 在字符串中迭代 unicode 编码。
	// 第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
	for i, c := range "abA" {
		fmt.Println(i, c)
	}
}
