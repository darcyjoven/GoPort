package utils

import (
	"context"

	"gorm.io/gorm"
)

type ColumnMeta struct {
	Name      string `json:"name"`      // 字段名
	DBType    string `json:"dbType"`    // Oracle类型: VARCHAR2, NUMBER, DATE等
	GoType    string `json:"goType"`    // Go扫描类型
	Nullable  bool   `json:"nullable"`  // 是否可空
	Length    int64  `json:"length"`    // 长度
	Precision int64  `json:"precision"` // 精度
	Scale     int64  `json:"scale"`     // 小数位
}

type Result struct {
	Data    [][]any      `json:"data"`
	Columns []ColumnMeta `json:"columns"`
}

// QueryArray 根据指定SQL查询结果和字段信息
func QueryArray(ctx context.Context, db *gorm.DB, sql string, args ...any) (*Result, error) {
	rows, err := db.Raw(sql, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	numCols := len(colTypes)
	columns := make([]ColumnMeta, 0, numCols)
	for _, ct := range colTypes {
		precision, scale, _ := ct.DecimalSize()
		length, _ := ct.Length()
		nullable, _ := ct.Nullable()

		columns = append(columns, ColumnMeta{
			Name:      ct.Name(),
			DBType:    ct.DatabaseTypeName(),
			GoType:    ct.ScanType().String(),
			Nullable:  nullable,
			Length:    length,
			Precision: precision,
			Scale:     scale,
		})
	}

	// 返回空切片而不是 nil，避免调用方 panic
	results := make([][]any, 0)

	// 预分配指针切片（列数固定，可复用）
	ptrs := make([]any, numCols)

	for rows.Next() {
		// 检查 context 是否已取消 / 超时
		if ctx != nil {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
		}

		values := make([]any, numCols)
		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		// 转换 []byte 为 string（若追求零拷贝可考虑 unsafe，但需谨慎）
		for i, v := range values {
			if b, ok := v.([]byte); ok {
				values[i] = string(b)
			}
		}

		results = append(results, values)
	}

	// 检查遍历过程中是否发生错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &Result{
		Data:    results,
		Columns: columns,
	}, nil
}
