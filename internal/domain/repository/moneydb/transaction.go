package moneydb

import (
	"errors"
	"fmt"
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
)

type Transaction struct {
	// Id
	Id int64 `json:"id"`
	// 金额
	Amount float64 `json:"amount"`
	// 描述
	Description string `json:"description"`
	// 使用人
	UserId int64 `json:"userId"`
	// 金额类型（收入/支出）
	Type int `json:"type"`
	// 用途分类
	Category int `json:"category"`
	// 账户
	Account int `json:"account"`
	// 操作时间
	Time model.LocalTime `json:"time"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

type TransactionIndex struct {
	// Id
	Id int64 `json:"id"`
	// 描述
	Description string `json:"description"`
	// 金额类型（收入/支出）
	Type int `json:"type"`
	// 用途分类
	Category int `json:"category"`
}

var EsIndex = "transaction_index"

func (record *Transaction) Insert() error {
	// 插入mysql
	_, err := infrastructure.Mysql.InsertOne(record)
	if err != nil {
		return err
	}

	// 插入es
	/*transactionIndex := TransactionIndex{
		Id:          record.Id,
		Description: record.Description,
		Type:        record.Type,
		Category:    record.Category,
	}
	_, err = infrastructure.EsClient.Index(EsIndex).Id(strconv.FormatInt(record.Id, 10)).Request(transactionIndex).Do(context.Background())*/
	return err
}

func (record *Transaction) Update() error {
	_, err := infrastructure.Mysql.ID(record.Id).Omit("GmtCreate").Update(record)
	if err != nil {
		return err
	}
	/*transactionIndex := TransactionIndex{
		Id:          record.Id,
		Description: record.Description,
		Type:        record.Type,
		Category:    record.Category,
	}
	// 更新es
	_, err = infrastructure.EsClient.Index(EsIndex).Id(strconv.FormatInt(record.Id, 10)).Request(transactionIndex).Do(context.Background())*/
	return nil
}

func (record *Transaction) Delete() error {
	if record.Id == 0 {
		return errors.New("transaction id can't be null")
	}
	_, err := infrastructure.Mysql.Table(&Transaction{}).ID(record.Id).Update(map[string]interface{}{"is_deleted": true})
	if err != nil {
		return err
	}

	// es 根据id删除
	//matchQuery := `{"query": {"match": {"id": "` + strconv.FormatInt(record.Id, 10) + `"}}}`
	//request := esapi.DeleteByQueryRequest{
	//	Index: []string{EsIndex},
	//	Body:  strings.NewReader(matchQuery),
	//}
	//_, err = request.Do(context.Background(), infrastructure.EsClient)
	/*request := esapi.DeleteRequest{
		Index:      EsIndex,
		DocumentID: strconv.FormatInt(record.Id, 10),
	}
	_, err = request.Do(context.Background(), infrastructure.EsClient)*/
	return err
}

func BatchPutTransaction(transactions []*Transaction) error {
	// 事务管理
	session := infrastructure.Mysql.NewSession()

	defer session.Close()
	_ = session.Begin()

	var err error
	var num int64
	for _, transaction := range transactions {
		if transaction.Id == 0 {
			num, err = session.Insert(transaction)
			fmt.Println(num)
		} else {
			num, err = session.Omit("GmtCreate").Update(transaction)
			fmt.Println(num)
		}
		if err != nil {
			_ = session.Rollback()
			break
		}
	}
	// 事务必须提交才行
	err = session.Commit()
	if err != nil {
		return err
	}

	// 插入数据库成功后进行es的插入
	//for _, transaction := range transactions {
	//	// 插入es
	//	//transactionIndex := TransactionIndex{
	//	//	Id:          transaction.Id,
	//	//	Description: transaction.Description,
	//	//	Type:        transaction.Type,
	//	//	Category:    transaction.Category,
	//	//}
	//	// _, err := infrastructure.EsClient.Index(EsIndex).Id(strconv.FormatInt(transaction.Id, 10)).Request(transactionIndex).Do(context.Background())
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
