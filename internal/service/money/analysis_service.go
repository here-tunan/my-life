package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"go-my-life/internal/domain/entity/family"
	"go-my-life/internal/domain/repository/moneydb"
)

type TransactionsAnalysisParam struct {
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Type      int     `json:"type"`
	UserIds   []int64 `json:"userIds"`
	FamilyId  int64   `json:"familyId"`
}

type TransactionsAnalysisResponse struct {
	Total                decimal.Decimal                   `json:"total"`
	Expenditure          decimal.Decimal                   `json:"expenditure"`
	Income               decimal.Decimal                   `json:"income"`
	CategoryAggregations []TransactionsCategoryAggregation `json:"categoryAggregations"`
	UserAggregations     []TransactionsUserAggregation     `json:"userAggregations"`
}

type TransactionsCategoryAggregation struct {
	CategoryId   int             `json:"categoryId"`
	CategoryName string          `json:"categoryName"`
	Percent      decimal.Decimal `json:"percent"`
	Amount       decimal.Decimal `json:"amount"`
}

type TransactionsUserAggregation struct {
	UserId   int64           `json:"userId"`
	UserName string          `json:"userName"`
	Percent  decimal.Decimal `json:"percent"`
	Amount   decimal.Decimal `json:"amount"`
}

// TransactionsAnalysis 用于做一些复杂的查询，聚合查询等
func TransactionsAnalysis(param *TransactionsAnalysisParam) (*TransactionsAnalysisResponse, error) {
	if param.FamilyId != 0 {
		familyEntity := family.Family{Id: param.FamilyId}
		err := familyEntity.GetFamily()
		if err != nil {
			return nil, err
		}
		// 收集UserId
		if len(familyEntity.Members) == 0 {
			return nil, errors.New("no family members")
		} else {
			for _, member := range familyEntity.Members {
				param.UserIds = append(param.UserIds, member.UserId)
			}
		}
	}

	// 查询所有满足条件的 transactions
	queryParam := moneydb.TransactionQueryParam{
		UserIds: param.UserIds,
		// Type:      param.Type,
		StartTime: param.StartTime,
		EndTime:   param.EndTime,
	}
	allTransactions, _ := moneydb.QueryTransaction(queryParam)
	var thisTypeTransactions []moneydb.Transaction
	var totalAmount decimal.Decimal
	var expenditure decimal.Decimal
	var income decimal.Decimal

	for _, transaction := range allTransactions {
		if transaction.Type == param.Type {
			totalAmount = totalAmount.Add(decimal.NewFromFloat(transaction.Amount))
			thisTypeTransactions = append(thisTypeTransactions, transaction)
		}
		if transaction.Type == 1 {
			// 收入
			income = income.Add(decimal.NewFromFloat(transaction.Amount))
		} else {
			// 支出
			expenditure = expenditure.Add(decimal.NewFromFloat(transaction.Amount))
		}
	}
	response := &TransactionsAnalysisResponse{
		Total:                totalAmount,
		Expenditure:          expenditure,
		Income:               income,
		CategoryAggregations: collectCategoryTransactionAggregation(thisTypeTransactions),
		UserAggregations:     collectUserTransactionAggregation(thisTypeTransactions),
	}
	return response, nil
}

func collectUserTransactionAggregation(allTransactions []moneydb.Transaction) []TransactionsUserAggregation {
	// 使用map进行收集
	userTransactionMap := make(map[int64][]moneydb.Transaction)

	var totalAmount decimal.Decimal
	for _, transaction := range allTransactions {
		response, existed := userTransactionMap[transaction.UserId]
		if existed {
			response = append(response, transaction)
			userTransactionMap[transaction.UserId] = response
		} else {
			var list []moneydb.Transaction
			list = append(list, transaction)
			userTransactionMap[transaction.UserId] = list
		}
		totalAmount = totalAmount.Add(decimal.NewFromFloat(transaction.Amount))
	}

	if len(userTransactionMap) == 0 {
		return nil
	}

	var results []TransactionsUserAggregation
	for userId, transactions := range userTransactionMap {
		result := TransactionsUserAggregation{
			UserId:  userId,
			Percent: decimal.Decimal{},
			Amount:  decimal.Decimal{},
		}
		for _, transaction := range transactions {
			result.Amount = result.Amount.Add(decimal.NewFromFloat(transaction.Amount))
		}
		result.Percent = result.Amount.Mul(decimal.NewFromInt(100)).DivRound(totalAmount, 2)
		results = append(results, result)
	}
	return results
}

func collectCategoryTransactionAggregation(allTransactions []moneydb.Transaction) []TransactionsCategoryAggregation {
	// 使用map进行收集
	categoryTransactionMap := make(map[int][]moneydb.Transaction)

	var totalAmount decimal.Decimal
	for _, transaction := range allTransactions {
		response, existed := categoryTransactionMap[transaction.Category]
		if existed {
			response = append(response, transaction)
			categoryTransactionMap[transaction.Category] = response
		} else {
			var list []moneydb.Transaction
			list = append(list, transaction)
			categoryTransactionMap[transaction.Category] = list
		}
		totalAmount = totalAmount.Add(decimal.NewFromFloat(transaction.Amount))
	}

	if len(categoryTransactionMap) == 0 {
		return nil
	}

	// find all category
	categories, _ := moneydb.AllCategory()

	categoriesMap := make(map[int]string)
	for _, category := range categories {
		categoriesMap[int(category.Id)] = category.Name
	}

	var results []TransactionsCategoryAggregation
	for categoryId, transactions := range categoryTransactionMap {
		result := TransactionsCategoryAggregation{
			CategoryId:   categoryId,
			CategoryName: categoriesMap[categoryId],
			Percent:      decimal.Decimal{},
			Amount:       decimal.Decimal{},
		}
		for _, transaction := range transactions {
			result.Amount = result.Amount.Add(decimal.NewFromFloat(transaction.Amount))
		}
		result.Percent = result.Amount.Mul(decimal.NewFromInt(100)).DivRound(totalAmount, 2)
		results = append(results, result)
	}
	return results
}
