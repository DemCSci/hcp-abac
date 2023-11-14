package api

import (
	"client-go-gateway/contract"
	"client-go-gateway/model"
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ToolApi struct {
	respUtil utils.RespUtil
}

func (toolApi *ToolApi) Init(ctx *gin.Context) {

	con := setting.ClientInfoMap["org1MSP"].Contract

	//初始化 所有用户身份
	//for _, val := range setting.ClientInfoMap {
	//	contract.AddUser(val.Contract)
	//}
	contract.AddUser(con)
	fmt.Println("用户身份初始化完成")
	// 添加公有属性
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return
	}
	id := "attribute:" + uuid
	attribute := &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "age",
		Value:      "40",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      50,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}
	fmt.Println("公有属性添加完成")
	resourceId, err := utils.GenerateUUID()
	resourceId = "resource:" + resourceId
	resource := &model.Resource{
		Id:          resourceId,
		Owner:       "",
		Url:         "http://www.baidu.com",
		Description: "访问百度",
		Controllers: nil,
	}
	addResourceRes := contract.AddResource(con, resource)
	if err != nil {
		return
	}
	log.Println(addResourceRes)
	log.Println("资源添加完成,资源id：" + resourceId)
	attributeId, _ := utils.GenerateUUID()
	//发布私有属性
	privateAttributeId := "attribute:" + resourceId + ":" + attributeId
	//获取当前的身份
	identity := contract.GetSubmittingClientIdentity(con)
	privateAttribute := &contract.Attribute{
		Id:         privateAttributeId,
		Type:       "PRIVATE",
		ResourceId: resourceId,
		Owner:      identity,
		Key:        "occupation",
		Value:      "doctor",
		NotBefore:  "1669791474807",
		NotAfter:   "1672383443000",
		Money:      50,
	}
	privateAttributeId, err = contract.PublishPrivateAttribute(con, privateAttribute)
	if err != nil {
		log.Printf("发布私有属性失败: %v \n", err)
	}
	//添加私有属性
	buyPrivateAttributeRequest := &model.BuyAttributeRequest{
		Buyer:       identity,
		Seller:      identity,
		AttributeId: privateAttributeId,
	}
	_, err = contract.BuyPrivateAttribute(con, buyPrivateAttributeRequest)
	if err != nil {
		log.Printf("添加私有属性失败: %v \n", err)
	}
	//添加资源控制器
	request := &model.AddResourceControllerRequest{
		ResourceId:   resourceId,
		ControllerId: identity,
	}
	res := contract.AddResourceController(con, request)
	log.Printf("添加资源控制器成功 %v\n", res)
	log.Println("全部初始化完成")
	toolApi.respUtil.SuccessResp("全部初始化完毕", ctx)
}

//为了4个属性判断增加额外的个属性
func (toolApi *ToolApi) Init4(ctx *gin.Context) {

	con := setting.ClientInfoMap["org1MSP"].Contract

	// 添加公有属性
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return
	}
	id := "attribute:" + uuid
	attribute := &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "ip",
		Value:      "192.168.2.1",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      50,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}
	fmt.Println("公有属性添加完成")
	log.Println("全部初始化完成")
	toolApi.respUtil.SuccessResp("全部初始化完毕", ctx)
}

//为了8个属性在4个属性的基础上判断增加额外的个属性
func (toolApi *ToolApi) Init8(ctx *gin.Context) {

	con := setting.ClientInfoMap["org1MSP"].Contract

	// 添加公有属性
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return
	}
	id := "attribute:" + uuid
	attribute := &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "nationality",
		Value:      "American",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "company",
		Value:      "google",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "group",
		Value:      "softDevelop",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "rank",
		Value:      "8",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	fmt.Println("公有属性添加完成")
	log.Println("全部初始化完成")
	toolApi.respUtil.SuccessResp("全部初始化完毕", ctx)
}

//为了16个属性在8个属性的基础上判断增加额外的个属性
func (toolApi *ToolApi) Init16(ctx *gin.Context) {

	con := setting.ClientInfoMap["org1MSP"].Contract

	// 添加公有属性
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return
	}
	id := "attribute:" + uuid
	attribute := &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A1",
		Value:      "V1",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A2",
		Value:      "V2",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A3",
		Value:      "V3",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A4",
		Value:      "V4",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}
	// 添加公有属性
	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A5",
		Value:      "V5",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A6",
		Value:      "V6",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A7",
		Value:      "V7",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}

	uuid, err = utils.GenerateUUID()
	if err != nil {
		return
	}
	id = "attribute:" + uuid
	attribute = &contract.Attribute{
		Id:         id,
		Type:       "PUBLIC",
		ResourceId: "",
		Owner:      "",
		Key:        "A8",
		Value:      "V8",
		NotBefore:  "1669791474807",
		NotAfter:   "1772383443000",
		Money:      10,
	}
	_, err = contract.AddAttribute(con, attribute)
	if err != nil {
		fmt.Println("公有属性添加失败")
		return
	}
	fmt.Println("公有属性添加完成")
	log.Println("全部初始化完成")
	toolApi.respUtil.SuccessResp("全部初始化完毕", ctx)
}

//为了32个属性在16个属性的基础上判断增加额外的个属性
func (toolApi *ToolApi) Init32(ctx *gin.Context) {

	con := setting.ClientInfoMap["org1MSP"].Contract

	// 从 A9-V9  到 A24-V24
	// 添加公有属性
	keyNum := 9
	valueNum := 9
	for keyNum <= 24 {
		k := "A" + strconv.Itoa(keyNum)
		v := "V" + strconv.Itoa(valueNum)

		uuid, err := utils.GenerateUUID()
		if err != nil {
			return
		}
		id := "attribute:" + uuid
		attribute := &contract.Attribute{
			Id:         id,
			Type:       "PUBLIC",
			ResourceId: "",
			Owner:      "",
			Key:        k,
			Value:      v,
			NotBefore:  "1669791474807",
			NotAfter:   "1772383443000",
			Money:      10,
		}
		_, err = contract.AddAttribute(con, attribute)
		if err != nil {
			fmt.Println("公有属性添加失败")
			return
		}
		fmt.Printf("K=%s, V=%s\n 属性添加完成", k, v)
		keyNum++
		valueNum++
	}

	fmt.Println("公有属性添加完成")
	log.Println("全部初始化完成")
	toolApi.respUtil.SuccessResp("全部初始化完毕", ctx)
}
