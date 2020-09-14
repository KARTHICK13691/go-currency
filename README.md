# go-currency

A  current currency course and its historical fluctuations API. Data taken from European Central Bank API.

Database = postgresql
Authentication = JWT
Framework = gorilla mux

Response format.

```json
{
    "base": "EUR",
    "date": "2020-09-14",
    "rates": {
        "USD": 1.1876,
        "JPY": 125.82,
        "BGN": 1.9558,
        "CZK": 26.660,
        "...": "and so on...",
    }
}
```



### **Run**

```bash
go run go-currency.go
```



