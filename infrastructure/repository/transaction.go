package repository

import (
    "database/sql"
    "github.com/brandaoplaster/payment-gateway/domain"
)

type TransactionRepository struct {
    db               *sql.DB
    creditCardRepo   *CreditCardRepository
}

func NewTransactionRepository(db *sql.DB, ccRepo *CreditCardRepository) *TransactionRepository {
    return &TransactionRepository{
        db:             db,
        creditCardRepo: ccRepo,
    }
}

func (r *TransactionRepository) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
    stmt, err := r.db.Prepare(`insert into transactions(id, credit_card_id, amount, status, description, store, created_at)
                                values($1, $2, $3, $4, $5, $6, $7)`)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        transaction.ID,
        transaction.CreditCardId,
        transaction.Amount,
        transaction.Status,
        transaction.Description,
        transaction.Store,
        transaction.CreatedAt,
    )
    if err != nil {
        return err
    }

    if transaction.Status == "approved" {
        err = r.creditCardRepo.UpdateBalance(creditCard)
        if err != nil {
            return err
        }
    }

    return nil
}
