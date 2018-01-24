package main

import (
                "encoding/base64"
                "github.com/fullsailor/pkcs7"
                "fmt"
                
)

//Take the PEM encoded pkcs7 from the Rest API call and strip off the headers, footers, line feeds and carriage returns.

var pkcs7PEMEncoded = "MIAGCSqGSIb3DQEHAqCAMIACAQExADCABgkqhkiG9w0BBwEAAKCAMIIDGDCCAgCgAwIBAgIUJc1OQKJ3EVUslZijiUR5qnKcXs8wDQYJKoZIhvcNAQEFBQAwgZIxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJOSjEWMBQGA1UEBwwNQmFza2luZyBSaWRnZTEZMBcGA1UECgwQVmVyaXpvbiBXaXJlbGVzczEbMBkGA1UECwwSVGhpbmdTcGFjZSBEZXZlbG9wMSYwJAYDVQQDDB1UaGluZ1NwYWNlIFByaW1hcnkgSXNzdWluZyBDQTAeFw0xNzA0MjEyMzEyNTRaFw0xODA0MjEyMzEyNDZaMGwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MREwDwYDVQQHDAhCZXJuYXJkczEQMA4GA1UECgwHVmVyaXpvbjEMMAoGA1UECwwDSW9UMRUwEwYDVQQDDAx0c19iZXRhXzE4MDEwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAOIRNv9gz5IZ9ZXZRhO5vaaYxckhOBM0PFzxpJ+bZkC0TwdUqWV4fUNycgQcbFJMCSWpyvPPuXFAZa09GMB2nUyGFoxx01CqhUw/EceEUJi0nL3IWAFS9xKKkjDzrPbgNqHm7l5gF1uTGoFVJ+m6xd7c9rHvSh9n9fGGAqPZhcz7AgMBAAGjDzANMAsGA1UdDwQEAwIFoDANBgkqhkiG9w0BAQUFAAOCAQEArwFRHVMLOQgsVjbS5dwfJWwTEmDNIYyJo6UMzDfOm3zZoLHqSCVmZ53HDFH2hv4ZVneW4XhkdlRacmAykAyF3tgjlp4auscMIsOJAd53kY2HsBXPKdwb0l0wmVunL0l+OCzphDhcnUvbbcxZ34ftwSeFkANSpmGjUdiqUOIiLmn8ac+lZq9vZiNHs6YANlG0WCpwUxXOY5xVugIDJwKGD9sd7YoGnEDG7TpqtRPtJBFw9Ox4YQnINOnTwnWrpNRfiRKYHQE8RsWSGFm6ZsMupoaP04JyvIk5tdclci00u7552+Yqa22yd1FhxOFT/F6Q+HGmdXlTdsMSK44+uTzNSDCCBOswggPToAMCAQICFAQddoMIlwAfUC/NBbSMF/7rsq0yMA0GCSqGSIb3DQEBCwUAMIGaMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFjAUBgNVBAcMDUJhc2tpbmcgUmlkZ2UxGTAXBgNVBAoMEFZlcml6b24gV2lyZWxlc3MxGzAZBgNVBAsMElRoaW5nU3BhY2UgRGV2ZWxvcDEuMCwGA1UEAwwlVGhpbmdTcGFjZSBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eTAeFw0xNjA3MTgxNTAwMjBaFw0yNjA3MTgxNTAwMjBaMIGSMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFjAUBgNVBAcMDUJhc2tpbmcgUmlkZ2UxGTAXBgNVBAoMEFZlcml6b24gV2lyZWxlc3MxGzAZBgNVBAsMElRoaW5nU3BhY2UgRGV2ZWxvcDEmMCQGA1UEAwwdVGhpbmdTcGFjZSBQcmltYXJ5IElzc3VpbmcgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC7WOcpSRrSvMiN005ojukcexYOxuHN+znE5EYH9xCjepULt/Wx9WmpvOyz/Ro65uIQXJYgGoU1nGprz7WDlLGXrV3Nyr67ffYBrgKZtKWkQD/8nGPOxJ+q20QDTnqN/cANjuJkImlIdBM+IEDI6D/dsiznad0scK8aSC3E1EV7UBVldW0g6pea9YgRrqM1u24wWHacVCn1Twm7GZrAirYNUCLxECJqGIvnyHvEsmrvlQZbQGnFmWBsXHKj4hSvTaQIASDfPjJN6GxGqCeuzLWxNdULsPgpiZv9Bf7gx5Soj5WNQKTEe8F4JG3RgtP/zJMeCvQSaSEwKpAsHgM6h4EZAgMBAAGjggEtMIIBKTASBgNVHRMBAf8ECDAGAQH/AgECMH4GCCsGAQUFBwEBBHIwcDBBBggrBgEFBQcwAoY1aHR0cDovL2lvdC1jZHAtcG9jLnZlcml6b24uY29tL0lPVFNDL0FJQS90c3Jvb3RjYS5jcnQwKwYIKwYBBQUHMAGGH2h0dHA6Ly9pb3Qtb2NzcC1wb2MudmVyaXpvbi5jb20wCwYDVR0PBAQDAgHGMB8GA1UdIwQYMBaAFPKB+f/NtiyWniqcMssuSm0j0a/UMEYGA1UdHwQ/MD0wO6A5oDeGNWh0dHA6Ly9pb3QtY2RwLXBvYy52ZXJpem9uLmNvbS9JT1RTQy9DRFAvdHNyb290Y2EuY3JsMB0GA1UdDgQWBBTKTUBntYaRMEUNzWMTUvwLmsjhYTANBgkqhkiG9w0BAQsFAAOCAQEAGHMSRee2wAWWwMIb12HUgcQbUbny3ow3U1UhCNWcWtMeoLPEHIIaaDN7pTp0LjY0wGaD1uwQtVgQ0NxcmNTpOFuyWDReZNg1XfI0iJznonYCFIg19uomiDHlSub4FkBGa2Q5voBiZ67oMqlc+BwqI98A4W0pmB2gY8EBaEi6c2K20E5+ueCN7mCBQ1QqWMtGBtQwHA2FAcBIj49nKUL0kbGCyNsYKoXkiVqTVgOPYeTmtFZNqAevV6vet0zbDX/edOGuy04p1gyzm97yh0Xk5oFQg11BuLz8fSbz53kz+6DVRcwDMZZHXRFu9uwGxzeEn8PezzhMM5Y92pVVcpTR7jCCBCcwggMPoAMCAQICFDocmVnOiQsOm1ShbB4jNncJ9g2pMA0GCSqGSIb3DQEBCwUAMIGaMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFjAUBgNVBAcMDUJhc2tpbmcgUmlkZ2UxGTAXBgNVBAoMEFZlcml6b24gV2lyZWxlc3MxGzAZBgNVBAsMElRoaW5nU3BhY2UgRGV2ZWxvcDEuMCwGA1UEAwwlVGhpbmdTcGFjZSBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eTAeFw0xNjA3MTgxMjUxMTlaFw00NjA3MTgxMjQzMzJaMIGaMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFjAUBgNVBAcMDUJhc2tpbmcgUmlkZ2UxGTAXBgNVBAoMEFZlcml6b24gV2lyZWxlc3MxGzAZBgNVBAsMElRoaW5nU3BhY2UgRGV2ZWxvcDEuMCwGA1UEAwwlVGhpbmdTcGFjZSBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANDfIUBahzBHKZqfK7Qkcs0aVjuCQZz1Dv+KanyxdzzYZED3TDAxbbVkqWnEvFPndXBH0Y1fXwueB6x8GxI6OjVNBXrWviRKdTOcbk1vhNuT9JHlfpL43GevSve9yvU4OwH4M6GesYy7nj9YATJl+Iac8a4cjtfFuxZc3qLwt7ywj+EPeovqhtkwD44WQ0tDizxXNfpJobwbXBf78TqxCRGXBCvw+Ulq6EBZtre6cOWY8F4yNzQbpHCwHluxLSkdTm3cM9NvPo0ch7BNo5LO6j9Zvqbq9VfkU9aFTOJb6fUGxb/eaUJmrxaeoAOMsD4mP+GGE6s4TlRkiWP2jAISpB8CAwEAAaNjMGEwEgYDVR0TAQH/BAgwBgEB/wIBAzALBgNVHQ8EBAMCAcYwHwYDVR0jBBgwFoAU8oH5/822LJaeKpwyyy5KbSPRr9QwHQYDVR0OBBYEFPKB+f/NtiyWniqcMssuSm0j0a/UMA0GCSqGSIb3DQEBCwUAA4IBAQC0ZKWcmyG+mMVIzvg/VugetouUkuyqswYRtLP9Icu32am6Op+lHEaRKHnmjg1VAB6/fMnx8UZJPmWPcBDE7hFO710DSib1ef3+Gbh9XdDl7C58vBWtWKi6vyTykJ74ZLU7QsNqdyPC72mPgq5uMX0XdlLBDkFr/ul/GgYrAPhHRrCti21c5dlCMkSY8QDroTTycdeeA8YnRZnlZdXPx9wErBR/WPyygAxJeTKuObHClSpE7O38FlWhIr2Xa9/w+ywJOeaXWGZy+FxvAasvFAxAFvP9+0m5ri5idbXEUMR/lWbh671tVYJfa0oo6FN8NLjQbd4StQGy/3vLA9mioYuhAAAxAAAAAAAAAA=="


func main() {
                //fullsailor requires DER encoded pkcs7 - You can take the PEM encoded PKCS7 and strip off the headers and footers and carriage returns line feeds.
                str := pkcs7PEMEncoded
                //Base64 decode the string
                data, err := base64.StdEncoding.DecodeString(str)
                if err != nil {
                                fmt.Println("error:", err)
                                return
                }
                fmt.Printf("%q\n", data)
                //Pass the DER encoded p7 to the pkcs7 package
                p7, err := pkcs7.Parse(data)
                //Parse out the certificates.
                fmt.Println("x509.Certificate - Certificate: ", p7.Certificates[0])
                Encode the raw asn1 certificate to PEM.
                cert := base64.StdEncoding.EncodeToString(p7.Certificates[0].Raw)
                fmt.Println("PEM Encoded Cert: " ,cert)
}
