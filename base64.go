package main

//import "os"
import "fmt"
import "encoding/base64"
func main() {

str_ca := `-----BEGIN CERTIFICATE-----
MIIDiTCCAnGgAwIBAgIJAIxAcRQ8Chj1MA0GCSqGSIb3DQEBBQUAMFoxCzAJBgNV
BAYTAkNOMQ8wDQYDVQQIDAZIZWliZWkxEDAOBgNVBAcMB0JlaWppbmcxDjAMBgNV
BAoMBXptIGhsMQswCQYDVQQLDAJ6bTELMAkGA1UEAwwCdm0wIBcNMTMxMjA3MDc0
MjI5WhgPMjExMzExMTMwNzQyMjlaMFoxCzAJBgNVBAYTAkNOMQ8wDQYDVQQIDAZI
ZWliZWkxEDAOBgNVBAcMB0JlaWppbmcxDjAMBgNVBAoMBXptIGhsMQswCQYDVQQL
DAJ6bTELMAkGA1UEAwwCdm0wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQD+PrmQCsMpDBkO1MtleslJnNeI2YkhfNqk0DYQUqiGJzxecXy5ZSsIhWNhKv+b
VZajqBRW7DZD7mWMAqJdc8UnFfc70+GZes/14W16uPojnZrE2YSaa1hDNywPXyO6
n2hKt5FNubVAmuiGAZwN8x1sm9WAZE1BQ+BHEsSVpVutzFmv/lxpWEcDKgk5svWm
kdOIJX1ouHHxS5HQpnnIK+B5lADGBU3H17ig8YTKeX9Dp4hzSLrP/S5HTRDgBJlJ
VsXqdoZerc8F/0g205kukLM+bj4qWvXHaajRvjRrwHhauUXNc2FneUi7igh2ECF6
0uzSx5YBU2qaBXXgmkg48mlHAgMBAAGjUDBOMB0GA1UdDgQWBBTSjYtwZukurM5Z
sLl0jrW1UkPzADAfBgNVHSMEGDAWgBTSjYtwZukurM5ZsLl0jrW1UkPzADAMBgNV
HRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQALBcSm4JVAGdpjrCy/VY0e+X/o
OaeohlA7bGYSRkfj5ZVulbCKAfoPo7raBFb87f2h/ST6ltoU4iTB1bYOalUTxvTZ
tAXWMoVERF1X+8I7w2+lGz8K1NFU2JkEK2mMuGZJ+9gqnFmxWfog/lat54rHvQAM
R/iIfKGTC5d+HHo677c6ZYibut1Nl9mr4rX1wwH/n95nsOI6LGyfRYTKRtqZQF3b
dSXv8F4QkG8tagn7rQ+sUPl0BrDcnmIaeAACtLgZEoXPGPqriT/1eRmyeGUQYYQj
j+Ux1HRrMH9ngdqgeEMoAfcYIvrRGeUY1GK6eJnu7ubi3QTuz4rNXfIRM4LE
-----END CERTIFICATE-----`


	data := []byte(str_ca)
	str := encryption(data,5)
//	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64: ",str)

	data ,_ = decryption(str,5)
//	data, _ = base64.StdEncoding.DecodeString(str)
	fmt.Println("old: ",string(data))
}


func encryption(data []byte,n int) (str string) {
	for i := 0; i < n ; i++{
		str =  base64.StdEncoding.EncodeToString(data)
		data = []byte(str)
	}
	return string(data)
}

func decryption(str string,n int) (date []byte,err error) {
	for i := 0; i < n ; i++{
		date,err = base64.StdEncoding.DecodeString(str)
		if err != nil{
			return date,err
		}
		str = string(date)
//		println(string(date))
	}
	return
}



