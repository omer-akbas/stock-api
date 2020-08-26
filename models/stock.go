package models

type Stock struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Price struct {
	// Id            int     `json:"id"`
	LastPrice     float64 `json:"lastPrice"`
	PreviousPrice float64 `json:"previousPrice"`
	Bid           float64 `json:"bid"`
	Ask           float64 `json:"ask"`
	Status        string  `json:"status"`
	StokId        int     `json:"stokId"`
}

type StockPrice struct {
	Stock
	Price
}

func StockByCode(code string) (StockPrice, error) {
	// var stock Stock
	// err := db.QueryRow("SELECT id, name, code FROM stock WHERE code = ? LIMIT 1", code).Scan(&stock.Id, &stock.Name, &stock.Code)
	// if err != nil {
	// 	return Stock{}, err
	// }

	// return stock, nil
	var stock Stock
	err := db.QueryRow("SELECT id, name, code FROM stock WHERE code = ? LIMIT 1", code).Scan(&stock.Id, &stock.Name, &stock.Code)
	if err != nil {
		return StockPrice{}, err
	}

	var price Price
	err = db.QueryRow("SELECT lastPrice, previousPrice, bid, ask, status, stockId  FROM price WHERE stockId = ? ORDER BY id DESC LIMIT 1", stock.Id).Scan(&price.LastPrice, &price.PreviousPrice, &price.Bid, &price.Ask, &price.Status, &price.StokId)
	if err != nil {
		return StockPrice{}, err
	}

	return StockPrice{Stock: stock, Price: price}, nil
}

func StockRatesByCode(code string) (string, error) {
	var id int
	err := db.QueryRow("SELECT id FROM stock WHERE code = ? LIMIT 1", code).Scan(&id)
	if err != nil {
		return "", err
	}

	var status string
	err = db.QueryRow("SELECT status FROM price WHERE stockId = ? ORDER BY id DESC LIMIT 1", id).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

func StockListRates() (Data, error) {
	rows, err := db.Query("SELECT s.name, s.code, p.lastPrice, p.previousPrice, p.bid, p.ask, p.`status`, p.stockId FROM stock AS s INNER JOIN price AS p ON s.id=p.stockId ORDER BY p.id DESC")
	if err != nil {
		return Data{}, err
	}
	sps := []StockPrice{}
	for rows.Next() {
		var sp StockPrice
		_ = rows.Scan(&sp.Name, &sp.Code, &sp.LastPrice, &sp.PreviousPrice, &sp.Bid, &sp.Ask, &sp.Status, &sp.StokId)
		sps = append(sps, sp)
	}
	return Data{Count: len(sps), Body: sps}, nil
}
