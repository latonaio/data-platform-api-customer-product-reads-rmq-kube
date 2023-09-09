package dpfm_api_output_formatter

import (
	"data-platform-api-customer-product-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToCustomerProduct(rows *sql.Rows) (*[]CustomerProduct, error) {
	defer rows.Close()
	customerProduct := make([]CustomerProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CustomerProduct{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Customer,
			&pm.SalesOrganization,
			&pm.DistributionChannel,
			&pm.Product,
			&pm.ProductByCustomer,
			&pm.ProductDescriptionByCustomer,
			&pm.BaseUnit,
			&pm.SalesUnit,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &customerProduct, err
		}

		data := pm
		customerProduct = append(customerProduct, CustomerProduct{
			BusinessPartner:              data.BusinessPartner,
			Customer:                     data.Customer,
			SalesOrganization:            data.SalesOrganization,
			DistributionChannel:          data.DistributionChannel,
			Product:                      data.Product,
			ProductByCustomer:            data.ProductByCustomer,
			ProductDescriptionByCustomer: data.ProductDescriptionByCustomer,
			BaseUnit:                     data.BaseUnit,
			SalesUnit:                    data.SalesUnit,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &customerProduct, nil
	}

	return &customerProduct, nil
}
