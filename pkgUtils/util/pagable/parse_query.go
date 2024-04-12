package pagable

import "github.com/gofiber/fiber/v2"

// GetQueryFromFiberCtx Get pagination query struct from
func GetQueryFromFiberCtx(c *fiber.Ctx) (*Query, error) {
	q := &Query{}

	if err := q.SetPage(c.Query("page")); err != nil {
		return nil, err
	}

	if err := q.SetSize(c.Query("size")); err != nil {
		return nil, err
	}

	queryString := string(c.Request().URI().QueryString())

	filters, err := FilterBinding(queryString)
	if err != nil {
		return q, err
	}

	q.ExpressionFilters = filters

	return q, nil
}
