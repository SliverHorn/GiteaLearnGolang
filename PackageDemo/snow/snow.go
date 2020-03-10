/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: snow.go
 * @Version: 1.13.8
 * @Time:  2:49 下午
 * @Project: LearnGolang
 * @Package:
 * @Software: GoLand
 */
package snow

import "fmt"

// 被calc包导入的一个包

func Snow() {
	fmt.Println("snow")
}

func init() {
	fmt.Println("snow.init()")
}
