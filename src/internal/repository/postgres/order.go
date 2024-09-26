package mypostgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"src/internal/dto"
	"src/internal/model"
	"src/internal/repository"
	"src/pkg/storage/postgres"
)

type OrderRepository struct {
	*postgres.Postgres
}

func NewOrderRepository(db *postgres.Postgres) repository.IOrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {

	query := r.Builder.
		Insert(orderTable).
		Columns(
			userIDField,
			statusField,
			totalPriceField,
			creationDateField,
			deliveryDateField,
			addressField,
			recepientNameField).
		Values(
			order.UserID,
			order.Status,
			order.TotalPrice,
			order.CreationDate,
			order.DeliveryDate,
			order.Address,
			order.RecepientName).
		Suffix("returning id")

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&order.ID,
	)

	if err != nil {
		return err
	}

	for _, line := range order.Lines {

		query := r.Builder.
			Insert(orderRacketTable).
			Columns(
				racketIDField,
				orderIDField,
				quantityField).
			Values(
				line.RacketID,
				order.ID,
				line.Quantity).
			Suffix("returning racket_id")

		sql, args, err := query.ToSql()

		if err != nil {
			return err
		}

		row := r.Pool.QueryRow(ctx, sql, args...)

		err = row.Scan(
			&line.RacketID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) Update(ctx context.Context, order *model.Order) error {

	query := r.Builder.
		Update(orderTable).
		Set(statusField, order.Status).
		Where(squirrel.And{squirrel.Eq{
			idField: order.ID,
		}})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, orderID int) error {

	query := r.Builder.
		Delete(orderTable).
		Where(squirrel.Eq{orderIDField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	query = r.Builder.
		Delete(deliveryTable).
		Where(squirrel.Eq{orderIDField: orderID})

	sql, args, err = query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Remove(ctx context.Context, orderID int) error {

	query := r.Builder.
		Delete(orderTable).
		Where(squirrel.Eq{idField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error) {

	ordersID, err := r.getAllOrders(ctx, req)

	if err != nil {
		return nil, err
	}

	var orders []*model.Order

	for _, id := range ordersID {

		order, err := r.GetOrderByID(ctx, id)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) getAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]int, error) {

	query := r.Builder.
		Select(idField).
		From(orderTable)

	query = req.Pagination.ToSQL(query)

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var ordersID []int

	for rows.Next() {

		orderID := 0

		err = rows.Scan(
			&orderID,
		)

		if err != nil {
			return nil, err
		}

		ordersID = append(ordersID, orderID)
	}

	return ordersID, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID int) (*model.Order, error) {

	query := r.Builder.
		Select(idField,
			userIDField,
			statusField,
			totalPriceField,
			creationDateField,
			deliveryDateField,
			addressField,
			recepientNameField).
		From(orderTable).
		Where(squirrel.Eq{idField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	order, err := r.rowToModel(row)

	if err != nil {
		return nil, err
	}

	query = r.Builder.
		Select(
			racketIDField,
			quantityField,
		).
		From(orderRacketTable).
		Where(squirrel.Eq{orderIDField: orderID})

	sql, args, err = query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var lines []*model.OrderLine

	for rows.Next() {

		line, err := r.rowToModelOrderRacket(rows)

		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	order.Lines = lines

	return order, nil
}

func (r *OrderRepository) GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error) {

	ordersID, err := r.getMyOrders(ctx, userID)

	if err != nil {
		return nil, err
	}

	var orders []*model.Order

	for _, id := range ordersID {

		order, err := r.GetOrderByID(ctx, id)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) getMyOrders(ctx context.Context, userID int) ([]int, error) {

	query := r.Builder.
		Select(idField).
		From(orderTable).
		Where(squirrel.Eq{userIDField: userID})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var ordersID []int

	for rows.Next() {

		orderID := 0

		err = rows.Scan(
			&orderID,
		)

		if err != nil {
			return nil, err
		}

		ordersID = append(ordersID, orderID)
	}

	return ordersID, nil
}

func (r *OrderRepository) rowToModel(row pgx.Row) (*model.Order, error) {

	var order model.Order

	err := row.Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
		&order.TotalPrice,
		&order.CreationDate,
		&order.DeliveryDate,
		&order.Address,
		&order.RecepientName,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) rowToModelOrderRacket(row pgx.Row) (*model.OrderLine, error) {

	var orderLine model.OrderLine

	err := row.Scan(
		&orderLine.RacketID,
		&orderLine.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &orderLine, nil
}
