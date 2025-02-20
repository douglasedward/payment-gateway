package factory

import "github.com/douglasedward/payment-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransaction() repository.TransactionRepository
}
