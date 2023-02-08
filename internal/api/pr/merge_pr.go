package pr

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: LixvYang 2690688423@qq.com
 * @Date: 2023-02-08 14:13:06
 * @LastEditors: LixvYang 2690688423@qq.com
 * @LastEditTime: Do not edit
 * @FilePath: Do not edit
 * @Description:
 *
 * Copyright (c) 2023 by Lixvyang, All Rights Reserved.
 * @param {*gin.Context} c
 */
func MergePR(c *gin.Context) {
	// pr是否被合并
	fmt.Println("MergePRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRr")
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(x))
	// 检测对应的issue  #merged closed

	//查看一下参与人是否绑定钱包，绑定的话，则想pr发起者500token  给issure发起者50token

	// 没有的话数据库暂时存储token等待github绑定钱包时一起绑定

}
