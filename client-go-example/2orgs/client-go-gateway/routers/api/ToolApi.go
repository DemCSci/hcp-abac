package api

import (
	"client-go-gateway/contract"
	"client-go-gateway/model"
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	//BuyPrivateAttributeRequest attributeRequest = BuyPrivateAttributeRequest.builder().attributeId("")
	//.buyer("user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e")
	//.seller("user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e")
	//.attributeId(attributeId)
	//.build();
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
