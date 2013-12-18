package main

//import "os"
import "fmt"
import "crypto/tls"
//import "encoding/pem"
func main() {
	println("i am tls")


	fmt.Println("cert:")
//	if cert, err := tls.LoadX509KeyPair("./key/server.crt", "./key/server.key"); err == nil {
	if cert, err := tls.X509KeyPair(localhostCert, localhostKey); err == nil {
		println("ok")
		fmt.Println("cert:",cert)
	} else {
		println("error")
		fmt.Println("error running server:", err)
	}

	println("--------------------------------------------------")

	println("--------------------------------------------------")

}


var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDbzCCAlegAwIBAgIJAOT6rqwHnpE0MA0GCSqGSIb3DQEBBQUAME0xCzAJBgNV
BAYTAkNOMQ8wDQYDVQQIDAZIZWliZWkxEDAOBgNVBAcMB0JlaWppbmcxDjAMBgNV
BAoMBXptIGhsMQswCQYDVQQLDAJ6bTAgFw0xMzEyMDUxMjM1NDhaGA8yMTEzMTEx
MTEyMzU0OFowTTELMAkGA1UEBhMCQ04xDzANBgNVBAgMBkhlaWJlaTEQMA4GA1UE
BwwHQmVpamluZzEOMAwGA1UECgwFem0gaGwxCzAJBgNVBAsMAnptMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr43eag8O2EiGwSCjhngvurYrN/KMiLMq
pYthrDadLCebDNhkJPkuJqjZa4eGVmjR/xMaVaQt6q+ODT2jvD4fuRKdbxu2rk9O
DHbNSFYtgxoFLhESokLUybyQ2zd+ea9wye+PA1yVWHCoPbuKzIa3F+SRcB3pnkem
KghUZxEuCVDsbV/V/8PtQd7JM60xRhcINCWzgSv3/kyCnXrGG6mpcyFEF/NVg6rg
JymIfIM++eqZpDUOMprENPiRYq8pQ4iuVTIYMeQyCt6CPuL2cysOtN4WAL1O22S7
CTLEbvcQqgK2w5bfrpDmQvXANMABkPJwi0EZlcn8GpMRS/+ULw1SkwIDAQABo1Aw
TjAdBgNVHQ4EFgQULaRb49aATzx1Y9sWOzyROMhNkeYwHwYDVR0jBBgwFoAULaRb
49aATzx1Y9sWOzyROMhNkeYwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOC
AQEAYxy8212a2znH6zo6cEqQtNi0yo2U7VkY0MVHqnreH1/1msxRmoBpqlIrWR4+
eLLfr5gSL2M9N3AiST25yQMV/eGalOzeiPHKYPyayIkJDV4D7XiaxN0sEZWLxs0w
njxEQOewU9EB5MdaGqOhBUVsWqfHDDe5P96/GhQBaLz6Q5YGapShrhVv/gLRFGsD
XzLSntMVKW3UNFjnc+e+AsB1dkknujSPUgm2mFVUCzlVMhGS6vz8WzMvbSL128xe
L0CkFTI2kPxBwgscoeHhQptK708/hNmuao9ZiYpumGHej7xo99JtgexTaJ5D1YVx
hS9QtGmC9lYXrutH6jL/ZebNSg==
-----END CERTIFICATE-----`)

var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAr43eag8O2EiGwSCjhngvurYrN/KMiLMqpYthrDadLCebDNhk
JPkuJqjZa4eGVmjR/xMaVaQt6q+ODT2jvD4fuRKdbxu2rk9ODHbNSFYtgxoFLhES
okLUybyQ2zd+ea9wye+PA1yVWHCoPbuKzIa3F+SRcB3pnkemKghUZxEuCVDsbV/V
/8PtQd7JM60xRhcINCWzgSv3/kyCnXrGG6mpcyFEF/NVg6rgJymIfIM++eqZpDUO
MprENPiRYq8pQ4iuVTIYMeQyCt6CPuL2cysOtN4WAL1O22S7CTLEbvcQqgK2w5bf
rpDmQvXANMABkPJwi0EZlcn8GpMRS/+ULw1SkwIDAQABAoIBAFjgbxistOxEk4xU
1NTwJeHV9j1l20Ydxtp2nSVNpPbEulefedvKF1ZVJ3Wr8BPxD1eeuTdgXiqxZC5J
5YoYvYC57uBgXHUVC0N+JeGYV4RG+RZgaFmjkgzsHT4Oc0Zxzp09Xd7q0WUr69EN
EHEJRsqe2g6z5iFOzs3615gyd2cZLQYPU5jCbk8J+l9EO6tZWjoLMHOideA+SWPW
T5FwenT9D57iqhO/OXpu+FAh1EKdiP7HchOcZmyEFX/kcg8IknmkL23lt0dpCXjO
xzSdVdyoIyTsdQOnq1lAdK6n7gAgbZY6q5zqh/c4y4VwrMIgWmMqNM/7mQqc71+0
qPtlVqECgYEA4JhFIWP8Bzt3C2naNiLBDjTazklRr146kbjesdJbAhrFInDFug9L
JMSqyr91XBgWvCPUMsAIRiORTD5/oYxY07gsFxabtcgm0WQpKcDOnu95jYCs/gWv
P7ZJibcvsRJiytXCwmdtIgZJrcm10bwN04K9Ea6e27yKbK5dTrFA/PECgYEAyBoc
bcMXner/LLkhX42+Q/ryBZg+8mGWPs4wql/iEjRO2w/yoOgLtrYvH+OxiKK27TNX
Vk35DqtENa/bKrLmW5Z3u8myKpLCDdCL+8hioukzjinCO61reu5UlCheoJy8pkI4
nt/XKm8E475wAnzi+UVPojFgCHBxk7SmTS0DF8MCgYEAu7lKfwBNVdY0SIvvYjRm
XPN8t9O8vPHrzoVRThxfCUFK6OCNxFioHFFt2A1zkMhpqFPiN5Ee2VE9TxcmIudm
D4DvCVVSTPzAxWAMr+OCUIIlfUrCTYh+1KgOi1dcO4r8nEFTip5cL4ZF60af/HBM
ti+ezLrB+TR8Fx8eynmJsmECgYBkWUPcf7P+S/VAFm94HftG2Dg7OgiHQQ053ell
58PV4UmSbkL9EqGdNWwYj/VeUktuVQr9iwjVpGoGinRcy27ei2zsTc/9ra+HgghR
ckKU9GIYbzSq/OMjuRXNXxllbnTk49zFP/gFnbtUAxLlDIA/BdVTv3MriEhNDJFB
rssYRQKBgQDATFDZLQumc+XjzE11L9VTpa4N/iOXh37KeEzFV+u8RHmg1YfueO3x
J+z7SalPFwLRL09ubJLkFqHWk6OAP5p0fvb1AZ0vZCyawLe9e5BzWsb26bPDpt52
xZ9maolINn1eRpPXjwY98rVXCCoC4fC4mq4TURjqvx4h93hxQW0D3A==
-----END RSA PRIVATE KEY-----`)



//var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
//MIIBdzCCASOgAwIBAgIBADALBgkqhkiG9w0BAQUwEjEQMA4GA1UEChMHQWNtZSBD
//bzAeFw03MDAxMDEwMDAwMDBaFw00OTEyMzEyMzU5NTlaMBIxEDAOBgNVBAoTB0Fj
//bWUgQ28wWjALBgkqhkiG9w0BAQEDSwAwSAJBAN55NcYKZeInyTuhcCwFMhDHCmwa
//IUSdtXdcbItRB/yfXGBhiex00IaLXQnSU+QZPRZWYqeTEbFSgihqi1PUDy8CAwEA
//AaNoMGYwDgYDVR0PAQH/BAQDAgCkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1Ud
//EwEB/wQFMAMBAf8wLgYDVR0RBCcwJYILZXhhbXBsZS5jb22HBH8AAAGHEAAAAAAA
//AAAAAAAAAAAAAAEwCwYJKoZIhvcNAQEFA0EAAoQn/ytgqpiLcZu9XKbCJsJcvkgk
//Se6AbGXgSlq+ZCEVo0qIwSgeBqmsJxUu7NCSOwVJLYNEBO2DtIxoYVk+MA==
//-----END CERTIFICATE-----`)
//
//// localhostKey is the private key for localhostCert.
//var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
//MIIBPAIBAAJBAN55NcYKZeInyTuhcCwFMhDHCmwaIUSdtXdcbItRB/yfXGBhiex0
//0IaLXQnSU+QZPRZWYqeTEbFSgihqi1PUDy8CAwEAAQJBAQdUx66rfh8sYsgfdcvV
//NoafYpnEcB5s4m/vSVe6SU7dCK6eYec9f9wpT353ljhDUHq3EbmE4foNzJngh35d
//AekCIQDhRQG5Li0Wj8TM4obOnnXUXf1jRv0UkzE9AHWLG5q3AwIhAPzSjpYUDjVW
//MCUXgckTpKCuGwbJk7424Nb8bLzf3kllAiA5mUBgjfr/WtFSJdWcPQ4Zt9KTMNKD
//EUO0ukpTwEIl6wIhAMbGqZK3zAAFdq8DD2jPx+UJXnh0rnOkZBzDtJ6/iN69AiEA
//1Aq8MJgTaYsDQWyU/hDq5YkDJc9e9DSCvUIzqxQWMQE=
//-----END RSA PRIVATE KEY-----`)



